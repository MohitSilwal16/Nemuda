package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/Nemuda/client/model"
	"github.com/Nemuda/client/utils"
	"github.com/gin-gonic/gin"
)

const BASE_URL = "http://localhost:8080/"

const INTERNAL_SERVER_ERROR_MESSAGE = "<script>alert('Internal Server Error');</script>"
const REQUEST_TIMED_OUT_MESSAGE = "<script>alert('Request Timed Out');</script>"
const BLOG_NOT_FOUND_MESSAGE = "<script>alert('Blog Not Found');</script>"

// Tags' slice
var tagsList = []string{"Political", "Technical", "Educational", "Geographical", "Programming", "Other"}

// Services
func Register(ctx *gin.Context) {
	// AVOID USING 204 BECAUSE IT DOESN'T SEND ANY CONTENT OR BODY

	// 200 => Registered Successfully
	// 201 => User data is not in format
	// 202 => Username or Password is Empty
	// 203 => Username is not in required format
	// 205 => Password is not in required format
	// 206 => Username is already used
	// 500 => Internal Server Error

	// Data isn't in format => Client isn't sending JSON data
	// Cannot Parse data to JSON => Cannot parse Client's data to JSON

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	var user model.User
	err := json.NewDecoder(ctx.Request.Body).Decode(&user)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Client's data is not in JSON format")
		response := "Data isn't in JSON format"

		RenderRegsiterPage(ctx, response)
		return
	}

	if user.Username == "" || user.Password == "" {
		response := "Username or Password is Empty"
		RenderRegsiterPage(ctx, response)
	} else if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(user.Username) {
		response := "Username should be alphanumeric b'twin 5-20 chars"
		RenderRegsiterPage(ctx, response)
	} else if !utils.IsPasswordInFormat(user.Password) {
		response := "Password: 8+ chars, lower & upper case, digit, symbol"
		RenderRegsiterPage(ctx, response)
	} else {
		userJSON, err := json.Marshal(user)

		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Cannot parse Client's data to JSON")
			response := "Cannot Parse data to JSON"
			RenderRegsiterPage(ctx, response)
			return
		}

		ctxTimeout, cancelFunction := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancelFunction()

		serviceURL := BASE_URL + "register"

		// Create Request with Timeout
		requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "POST", serviceURL, bytes.NewBuffer(userJSON))

		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Cannot Create POST Request with Context")

			RenderRegsiterPage(ctx, "Internal Server Error")
			return
		}
		requestToBackend_Server.Header.Set("Content-Type", "application/json")

		// Send request(with timeout) to back-end server
		res, err := http.DefaultClient.Do(requestToBackend_Server)

		if err != nil {
			if ctxTimeout.Err() == context.DeadlineExceeded {
				log.Println("Error: ", err)
				log.Println("Description: Back-end server didn't responsed in given time")
				RenderRegsiterPage(ctx, "Request Timed Out")
				return
			}
			log.Println("Error: ", err)
			log.Println("Description: Cannot send POST request(with timeout(context)) to back-end server")
			RenderRegsiterPage(ctx, "Internal Server Error")
			return
		}

		defer res.Body.Close()

		if res.StatusCode == 200 {
			responseJSONData, err := io.ReadAll(res.Body)

			if err != nil {
				log.Println("Error: ", err)
				log.Println("Decription: Cannot read body of response as JSON data")

				RenderRegsiterPage(ctx, "Internal Server Error")
				return
			}

			var responseDataStructure map[string]string
			err = json.Unmarshal(responseJSONData, &responseDataStructure)
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Decription: Cannot read body of response")

				RenderRegsiterPage(ctx, "Internal Server Error")
				return
			}

			sessionToken := responseDataStructure["sessionToken"]

			// No Session Token is sent from back-end server
			if sessionToken == "" {
				log.Println("Error: Session Token not provided by back-end server")
				RenderRegsiterPage(ctx, "Internal Server Error")
				return
			}
			// Save session token in cookie of user
			setSessionTokenInCookie(ctx.Writer, sessionToken)

			fetchBlogsByTag(ctx, "All", sessionToken)
		} else if res.StatusCode == 206 {
			RenderRegsiterPage(ctx, "Username is already used")
		} else if res.StatusCode == 500 {
			log.Println("Error: Back-end server has Internal Server Error")

			RenderRegsiterPage(ctx, "Internal Server Error")
		} else {
			log.Println("Bug: Unexpected Status Code ", res.StatusCode)

			RenderRegsiterPage(ctx, "Internal Server Error")
		}
	}
}

func Login(ctx *gin.Context) {
	// AVOID USING 204 BECAUSE IT DOESN'T SEND ANY CONTENT OR BODY

	// 200 => Login Successful
	// 201 => User data is not in format
	// 202 => Username or Password is Empty
	// 203 => Username is not in required format
	// 205 => User doesn't exists
	// 206 => Invalid Credentials
	// 500 => Internal Server Error

	// Data isn't in format => Client isn't sending JSON data
	// Cannot Parse data to JSON => Cannot parse Client's data to JSON

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	var user model.User
	err := json.NewDecoder(ctx.Request.Body).Decode(&user)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Client's data is not in JSON format")
		RenderLoginPage(ctx, "User data isn't in format")
		return
	}

	if user.Username == "" || user.Password == "" {
		response := "Username or Password is Empty"

		RenderLoginPage(ctx, response)
		return
	} else if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(user.Username) {
		response := "Username should be alphanumeric b'twin 5-20 chars"

		RenderLoginPage(ctx, response)
		return
	} else if !utils.IsPasswordInFormat(user.Password) {
		// Don't want to give idea to anonymous user idea about password pattern or format
		response := "Invalid Credentials"

		RenderLoginPage(ctx, response)
		return
	}
	userJSON, err := json.Marshal(user)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot parse Client's data to JSON")
		response := "Cannot Parse data to JSON"

		RenderLoginPage(ctx, response)
		return
	}

	ctxTimeout, cancelFunction := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunction()

	serviceURL := BASE_URL + "login"

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "POST", serviceURL, bytes.NewBuffer(userJSON))

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create POST Request with Context")
		RenderLoginPage(ctx, "Internal Server Error")
	}
	requestToBackend_Server.Header.Set("Content-Type", "application/json")

	// Send request(with timeout) to back-end server
	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")

			RenderLoginPage(ctx, "Request Timed Out")
			return
		}
		log.Println("Error: ", err)
		log.Println("Description: Cannot send POST request(with timeout(context)) to back-end server")

		RenderLoginPage(ctx, "Internal Server Error")
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		responseJSONData, err := io.ReadAll(res.Body)

		if err != nil {
			log.Println("Error: ", err)
			log.Println("Decription: Cannot read body of response")
			RenderLoginPage(ctx, "Internal Server Error")
			return
		}

		var responseDataStructure map[string]string
		err = json.Unmarshal(responseJSONData, &responseDataStructure)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Decription: Cannot read body of response as JSON data")
			RenderLoginPage(ctx, "Internal Server Error")
			return
		}
		sessionToken := responseDataStructure["sessionToken"]

		// No Session Token is sent from back-end server
		if sessionToken == "" {
			log.Println("Error: Session Token not provided by back-end server")

			RenderLoginPage(ctx, "Session Timed Out")
			return
		}
		// Save session token in cookie of user
		setSessionTokenInCookie(ctx.Writer, sessionToken)

		fetchBlogsByTag(ctx, "All", sessionToken)
	} else if res.StatusCode == 205 {
		RenderLoginPage(ctx, "User doesn't exists")
	} else if res.StatusCode == 206 {
		RenderLoginPage(ctx, "Invalid Credentials")
	} else if res.StatusCode == 500 {
		log.Println("Error: Back-end server has Internal Server Error")
		RenderLoginPage(ctx, "Internal Server Error")
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)

		RenderLoginPage(ctx, "Internal Server Error")
	}
}

func Logout(ctx *gin.Context) {
	// 200 => Log out Successful
	// 201 => Invalid Session Token
	// 500 => Internal Server Error

	// Data isn't in format => Client isn't sending JSON data
	// Cannot Parse data to JSON => Cannot parse Client's data to JSON

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	serviceURL := BASE_URL + "login/" + sessionToken

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "DELETE", serviceURL, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create DELETE Request with Context")

		RenderLoginPage(ctx, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	requestToBackend_Server.Header.Set("Content-Type", "application/json")

	// Send request(with timeout) to back-end server
	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")

			fmt.Fprint(ctx.Writer, REQUEST_TIMED_OUT_MESSAGE)
			return
		}
		log.Println("Error: ", err)
		log.Println("Description: Cannot send POST request(with timeout(context)) to back-end server")

		RenderLoginPage(ctx, "Internal Server Error")
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 || res.StatusCode == 201 {
		if res.StatusCode == 201 {
			log.Println("Invalid Session Token")
		}
		deleteSessionTokenFromCookie(ctx.Writer)

		// Even if session token is invalid log out user
		RenderLoginPage(ctx, "")
	} else if res.StatusCode == 500 {
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}
}

func SearchUserForRegistration(ctx *gin.Context) {
	// 200 => User found (Username is already used)
	// 201 => User not found (Username is not used yet)
	// 500 => Internal Server Error

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	username := ctx.Query("username")

	if len(username) < 5 {
		fmt.Fprint(ctx.Writer, "")
		return
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), time.Second)

	defer cancelFunc()

	serviceURL := BASE_URL + "users/" + username

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "GET", serviceURL, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create GET Request with Context")

		fmt.Fprint(ctx.Writer, "")
		return
	}

	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")
		} else {
			log.Println("Error: ", err)
			log.Println("Description: Cannot send GET request(with timeout(context)) to back-end server")
		}

		fmt.Fprint(ctx.Writer, "")
		return
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		fmt.Fprint(ctx.Writer, "Username is already used")
	} else if res.StatusCode == 201 {
		fmt.Fprint(ctx.Writer, "")
	} else if res.StatusCode == 500 {
		log.Println("Error: Back-end server has Internal Server Error")

		fmt.Fprint(ctx.Writer, "")
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)

		fmt.Fprint(ctx.Writer, "")
	}
}

func GetBlogsByTag(ctx *gin.Context) {
	// 200 => Blogs found
	// 201 => No blog found for the specific tag
	// 202 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	tag := ctx.Param("tag")
	tag = strings.Title(tag)

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	fetchBlogsByTag(ctx, tag, sessionToken)
}

func SearcBlogTitle_BeforePosting(ctx *gin.Context) {
	// 200 => Blog found
	// 201 => Blog not found
	// 202 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	title := ctx.Query("title")

	if len(title) < 5 {
		fmt.Fprint(ctx.Writer, "")
		return
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), time.Second)

	defer cancelFunc()

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	serviceURL := BASE_URL + "blogs/title/" + title + "?sessionToken=" + sessionToken

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "GET", serviceURL, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create GET Request with Context")

		fmt.Fprint(ctx.Writer, "")
		return
	}

	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")
		} else {
			log.Println("Error: ", err)
			log.Println("Description: Cannot send GET request(with timeout(context)) to back-end server")
		}

		fmt.Fprint(ctx.Writer, "")
		return
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		method := ctx.Param("method")
		if method == "post" {
			RenderPostBlogPage(ctx, "Title is already used")
		} else if method == "update" {
			RenderUpdateBlogPage(ctx, model.Blog{}, "Title is already used")
		} else {
			fmt.Fprint(ctx.Writer, "")
		}
	} else if res.StatusCode == 201 {
		fmt.Fprint(ctx.Writer, "")
	} else if res.StatusCode == 202 {
		RenderLoginPage(ctx, "Session Timed Out")
	} else if res.StatusCode == 500 {
		log.Println("Error: Back-end server has Internal Server Error")

		fmt.Fprint(ctx.Writer, "")
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)

		fmt.Fprint(ctx.Writer, "")
	}
}

func GetBlogByTitle(ctx *gin.Context) {
	// 200 => Blog found
	// 201 => Blog not found
	// 202 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	title := ctx.Param("title")

	if len(title) < 5 {
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	getBlogByTitle(ctx, title, "")
}

func PostBlog(ctx *gin.Context) {
	// AVOID USING 204 BECAUSE IT DOESN'T SEND ANY CONTENT OR BODY

	// 200 => Blog added
	// 201 => Title, Description, Tag is not in requested format
	// 202 => Title is already used
	// 203 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	var blog model.Blog

	blog.Title = ctx.Request.PostFormValue("title")
	blog.Tag = ctx.Request.PostFormValue("tag")
	blog.Description = ctx.Request.PostFormValue("description")

	if !regexp.MustCompile(`^[a-zA-Z0-9 ,'"&]*$`).MatchString(blog.Title) {
		response := "Title should be alphanumeric b'twin 5-20 chars"
		RenderPostBlogPage(ctx, response)
	} else if !utils.Contains(tagsList, blog.Tag) {
		response := "Unknown tag"
		RenderPostBlogPage(ctx, response)
	} else if len(blog.Description) < 4 || len(blog.Description) > 50 {
		response := "Min 4 letters & Max 50 letters"
		RenderPostBlogPage(ctx, response)
	} else {
		image, err := ctx.FormFile("image")
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Cannot fetch image from user")

			RenderPostBlogPage(ctx, "Failed to fetch image")
			return
		}

		if image.Header.Get("Content-Type") != "image/jpeg" &&
			image.Header.Get("Content-Type") != "image/png" &&
			image.Header.Get("Content-Type") != "image/jpg" {
			RenderPostBlogPage(ctx, "Invalid file type, upload an image")
			return
		}

		maxSize := 2 * 1024 * 1024 // 2MB
		if image.Size > int64(maxSize) {
			RenderPostBlogPage(ctx, "Image size exceeds 2 MB")
			return
		}

		blog.ImagePath = "./static/images/blogs/" + blog.Title + ".png"

		blog.Comments = []model.Comment{}

		blogJSON, err := json.Marshal(&blog)

		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Cannot parse Client's data to JSON")

			RenderPostBlogPage(ctx, "Internal Server Error")
			return
		}

		ctxTimeout, cancelFunction := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancelFunction()

		sessionToken := getSessionTokenFromCookie(ctx.Request)
		serviceURL := BASE_URL + "blogs?sessionToken=" + sessionToken

		// Create Request with Timeout
		requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "POST", serviceURL, bytes.NewBuffer(blogJSON))

		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Cannot Create POST Request with Context")

			RenderPostBlogPage(ctx, "Internal Server Error")
			return
		}
		requestToBackend_Server.Header.Set("Content-Type", "application/json")

		// Send request(with timeout) to back-end server
		res, err := http.DefaultClient.Do(requestToBackend_Server)

		if err != nil {
			if ctxTimeout.Err() == context.DeadlineExceeded {
				log.Println("Error: ", err)
				log.Println("Description: Back-end server didn't responsed in given time")

				RenderPostBlogPage(ctx, "Request Timed Out")
				return
			}
			log.Println("Error: ", err)
			log.Println("Description: Cannot send POST request(with timeout(context)) to back-end server")

			RenderPostBlogPage(ctx, "Internal Server Error")
			return
		}

		defer res.Body.Close()

		if res.StatusCode == 200 {
			sessionToken := getSessionTokenFromCookie(ctx.Request)

			err = ctx.SaveUploadedFile(image, blog.ImagePath)

			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Cannot save image of blog")

				fmt.Fprint(ctx.Writer, "Image of blog cannot be saved")
			}

			fetchBlogsByTag(ctx, "All", sessionToken)
		} else if res.StatusCode == 201 {
			RenderPostBlogPage(ctx, "Title, Description, Tag is not in format")
		} else if res.StatusCode == 202 {
			RenderPostBlogPage(ctx, "Title is already used")
		} else if res.StatusCode == 203 {
			RenderLoginPage(ctx, "Session Timed Out")
		} else if res.StatusCode == 500 {
			log.Println("Error: Back-end server has Internal Server Error")
			RenderPostBlogPage(ctx, "Internal Server Error")
		} else {
			log.Println("Bug: Unexpected Status Code ", res.StatusCode)
			RenderPostBlogPage(ctx, "Internal Server Error")
		}
	}
}

func LikeBlog(ctx *gin.Context) {
	// 200 => Blog liked
	// 201 => Blog already liked
	// 202 => Blog not found
	// 203 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	title := ctx.Param("title")

	if len(title) < 5 {
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	serviceURL := BASE_URL + "blogs/like/" + title + "?sessionToken=" + sessionToken

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "POST", serviceURL, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create POST Request with Context")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")
			fmt.Fprint(ctx.Writer, REQUEST_TIMED_OUT_MESSAGE)
		} else {
			log.Println("Error: ", err)
			log.Println("Description: Cannot send GET request(with timeout(context)) to back-end server")
			fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		}
		return
	}
	defer res.Body.Close()

	if res.StatusCode == 200 || res.StatusCode == 201 {
		getBlogByTitle(ctx, title, "")
	} else if res.StatusCode == 202 {
		fmt.Fprint(ctx.Writer, BLOG_NOT_FOUND_MESSAGE)
	} else if res.StatusCode == 203 {
		RenderLoginPage(ctx, "Session Timed Out")
	} else if res.StatusCode == 500 {
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}
}

func DislikeBlog(ctx *gin.Context) {
	// 200 => Blog disliked
	// 201 => Blog already disliked
	// 202 => Blog not found
	// 203 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	title := ctx.Param("title")

	if len(title) < 5 {
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	serviceURL := BASE_URL + "blogs/like/" + title + "?sessionToken=" + sessionToken

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "DELETE", serviceURL, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create DELETE Request with Context")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")
			fmt.Fprint(ctx.Writer, REQUEST_TIMED_OUT_MESSAGE)
		} else {
			log.Println("Error: ", err)
			log.Println("Description: Cannot send GET request(with timeout(context)) to back-end server")
			fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		}
		return
	}
	defer res.Body.Close()

	if res.StatusCode == 200 || res.StatusCode == 201 {
		getBlogByTitle(ctx, title, "")
	} else if res.StatusCode == 202 {
		fmt.Fprint(ctx.Writer, BLOG_NOT_FOUND_MESSAGE)
	} else if res.StatusCode == 203 {
		RenderLoginPage(ctx, "Session Timed Out")
	} else if res.StatusCode == 500 {
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}
}

func AddComment(ctx *gin.Context) {
	// 200 => Comment Added
	// 201 => Comment Description or Title is Empty
	// 202 => Blog not found
	// 203 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	title := ctx.Param("title")

	if len(title) < 5 {
		fmt.Fprint(ctx.Writer, BLOG_NOT_FOUND_MESSAGE)
		return
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	commentDescription := ctx.Query("comment")

	if commentDescription == "" {
		getBlogByTitle(ctx, title, "Comment or Title is Empty")
		return
	}

	if len(commentDescription) < 5 || len(commentDescription) > 50 {
		getBlogByTitle(ctx, title, "Comment: Min 5 & Max 50 letters")
		return
	}

	serviceURL := BASE_URL + "blogs/comment/" + commentDescription + "?sessionToken=" + sessionToken + "&title=" + title

	serviceURL = url.QueryEscape(serviceURL)

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "GET", serviceURL, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create GET Request with Context")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")
			fmt.Fprint(ctx.Writer, REQUEST_TIMED_OUT_MESSAGE)
		} else {
			log.Println("Error: ", err)
			log.Println("Description: Cannot send GET request(with timeout(context)) to back-end server")
			fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		}
		return
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		getBlogByTitle(ctx, title, "")
	} else if res.StatusCode == 202 {
		fmt.Fprint(ctx.Writer, BLOG_NOT_FOUND_MESSAGE)
	} else if res.StatusCode == 203 {
		RenderLoginPage(ctx, "Session Timed Out")
	} else if res.StatusCode == 500 {
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}
}

func UpdateBlog(ctx *gin.Context) {
	// AVOID USING 204 BECAUSE IT DOESN'T SEND ANY CONTENT OR BODY

	// 200 => Blog Updated
	// 201 => Data is not in correct format
	// 202 => User cannot update this blog
	// 203 => Blog not found
	// 205 => Invalid Session Token
	// 206 => Blog Title is already used
	// 500 => Internal Server Error

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	oldTitle := ctx.Param("title")

	var blog model.Blog

	blog.Title = ctx.Request.PostFormValue("title")
	blog.Tag = ctx.Request.PostFormValue("tag")
	blog.Description = ctx.Request.PostFormValue("description")

	if !regexp.MustCompile(`^[a-zA-Z0-9 ,'"&]*$`).MatchString(blog.Title) {
		response := "Title should be alphanumeric b'twin 5-20 chars"
		RenderUpdateBlogPage(ctx, blog, response)
	} else if !utils.Contains(tagsList, blog.Tag) {
		response := "Unknown tag"
		RenderUpdateBlogPage(ctx, blog, response)
	} else if len(blog.Description) < 4 || len(blog.Description) > 50 {
		response := "Min 4 letters & Max 50 letters"
		RenderUpdateBlogPage(ctx, blog, response)
	} else {
		image, err := ctx.FormFile("image")
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Cannot fetch image from user")

			RenderUpdateBlogPage(ctx, blog, "Failed to fetch image")
			return
		}

		if image.Header.Get("Content-Type") != "image/jpeg" &&
			image.Header.Get("Content-Type") != "image/png" &&
			image.Header.Get("Content-Type") != "image/jpg" {
			RenderUpdateBlogPage(ctx, blog, "Invalid file type, upload an image")
			return
		}

		maxSize := 2 * 1024 * 1024 // 2MB
		if image.Size > int64(maxSize) {
			RenderUpdateBlogPage(ctx, blog, "Image size exceeds 2 MB")
			return
		}

		blog.ImagePath = "./static/images/blogs/" + blog.Title + ".png"

		blog.Comments = []model.Comment{}

		blogJSON, err := json.Marshal(&blog)

		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Cannot parse Client's data to JSON")

			RenderUpdateBlogPage(ctx, blog, "Internal Server Error")
			return
		}

		ctxTimeout, cancelFunction := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancelFunction()

		sessionToken := getSessionTokenFromCookie(ctx.Request)
		serviceURL := BASE_URL + "blogs?sessionToken=" + sessionToken + "&title=" + oldTitle
		serviceURL = strings.Replace(serviceURL, " ", "%20", -1)

		// Create Request with Timeout
		requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "PUT", serviceURL, bytes.NewBuffer(blogJSON))

		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Cannot Create PUT Request with Context")

			RenderUpdateBlogPage(ctx, blog, "Internal Server Error")
			return
		}
		requestToBackend_Server.Header.Set("Content-Type", "application/json")

		// Send request(with timeout) to back-end server
		res, err := http.DefaultClient.Do(requestToBackend_Server)

		if err != nil {
			if ctxTimeout.Err() == context.DeadlineExceeded {
				log.Println("Error: ", err)
				log.Println("Description: Back-end server didn't responsed in given time")

				RenderUpdateBlogPage(ctx, blog, "Request Timed Out")
				return
			}
			log.Println("Error: ", err)
			log.Println("Description: Cannot send POST request(with timeout(context)) to back-end server")

			RenderUpdateBlogPage(ctx, blog, "Internal Server Error")
			return
		}

		defer res.Body.Close()

		if res.StatusCode == 200 {
			sessionToken := getSessionTokenFromCookie(ctx.Request)

			err = ctx.SaveUploadedFile(image, blog.ImagePath)

			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Cannot save image of blog")

				fmt.Fprint(ctx.Writer, "Image of blog cannot be saved")
			}
			oldImagePath := "./static/images/blogs/" + oldTitle + ".png"

			err = os.Remove(oldImagePath)

			if err != nil {
				if !os.IsNotExist(err) {
					log.Println("Error: ", err)
					log.Println("Description: Cannot delete ", oldImagePath)
				}
				// No need to return
			}

			fetchBlogsByTag(ctx, "All", sessionToken)
		} else if res.StatusCode == 201 {
			RenderUpdateBlogPage(ctx, blog, "Title, Description, Tag is not in format")
		} else if res.StatusCode == 202 {
			RenderUpdateBlogPage(ctx, blog, "User can't update blog")
		} else if res.StatusCode == 203 {
			fmt.Fprint(ctx.Writer, BLOG_NOT_FOUND_MESSAGE)
		} else if res.StatusCode == 205 {
			RenderLoginPage(ctx, "Session Timed Out")
		} else if res.StatusCode == 206 {
			RenderUpdateBlogPage(ctx, blog, "Title already used")
		} else if res.StatusCode == 500 {
			log.Println("Error: Back-end server has Internal Server Error")
			RenderUpdateBlogPage(ctx, blog, "Internal Server Error")
		} else {
			log.Println("Bug: Unexpected Status Code ", res.StatusCode)
			RenderUpdateBlogPage(ctx, blog, "Internal Server Error")
		}
	}
}

func DeleteBlog(ctx *gin.Context) {
	// AVOID USING 204 BECAUSE IT DOESN'T SEND ANY CONTENT OR BODY

	// 200 => Blog Deleted
	// 201 => Data is not in correct format
	// 202 => User cannot delete this blog
	// 203 => Blog not found
	// 205 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	title := ctx.Param("title")

	ctxTimeout, cancelFunction := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunction()

	sessionToken := getSessionTokenFromCookie(ctx.Request)
	serviceURL := BASE_URL + "blogs?sessionToken=" + sessionToken + "&title=" + title
	serviceURL = strings.Replace(serviceURL, " ", "%20", -1)

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "DELETE", serviceURL, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create DELETE Request with Context")

		getBlogByTitle(ctx, title, "")
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}
	requestToBackend_Server.Header.Set("Content-Type", "application/json")

	// Send request(with timeout) to back-end server
	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")

			getBlogByTitle(ctx, title, "")
			fmt.Fprint(ctx.Writer, REQUEST_TIMED_OUT_MESSAGE)
			return
		}
		log.Println("Error: ", err)
		log.Println("Description: Cannot send POST request(with timeout(context)) to back-end server")

		getBlogByTitle(ctx, title, "")
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		sessionToken := getSessionTokenFromCookie(ctx.Request)

		oldImagePath := "./static/images/blogs/" + title + ".png"

		err = os.Remove(oldImagePath)

		if err != nil {
			if !os.IsNotExist(err) {
				log.Println("Error: ", err)
				log.Println("Description: Cannot delete ", oldImagePath)
			}
			// No need to return
		}

		fetchBlogsByTag(ctx, "All", sessionToken)
	} else if res.StatusCode == 201 {
		getBlogByTitle(ctx, title, "")
		fmt.Fprint(ctx.Writer, "Data not in format")
	} else if res.StatusCode == 202 {
		getBlogByTitle(ctx, title, "")
		fmt.Fprint(ctx.Writer, "Cannot delete this blog")
	} else if res.StatusCode == 203 {
		fetchBlogsByTag(ctx, "All", sessionToken)
		fmt.Fprint(ctx.Writer, BLOG_NOT_FOUND_MESSAGE)
	} else if res.StatusCode == 205 {
		getBlogByTitle(ctx, title, "")
		RenderLoginPage(ctx, "Session Timed Out")
	} else if res.StatusCode == 500 {
		log.Println("Error: Back-end server has Internal Server Error")

		getBlogByTitle(ctx, title, "")
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)

		getBlogByTitle(ctx, title, "")
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}
}
