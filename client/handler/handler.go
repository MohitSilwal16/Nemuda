package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/Nemuda/client/model"
	"github.com/Nemuda/client/utils"
	"github.com/gin-gonic/gin"
)

const BASE_URL = "http://localhost:8080/"

const INTERNAL_SERVER_ERROR_MESSAGE = "<script>alert('Internal Server Error');</script>"
const REQUEST_TIMED_OUT_MESSAGE = "<script>alert('Request Timed Out');</script>"

// Tags' slice
var tagsList = []string{"Political", "Technical", "Educational", "Geographical", "Programming", "Other"}

// Helper methods
func setSessionTokenInCookie(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:  "sessionToken",
		Value: token,
		Path:  "/",
	}

	http.SetCookie(w, cookie)
}

func getSessionTokenFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("sessionToken")

	if err == http.ErrNoCookie {
		return ""
	} else if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error while Fetching Cookie")
		return ""
	}
	return cookie.Value
}

func deleteSessionTokenFromCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "sessionToken",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}

func fetchBlogsByTag(ctx *gin.Context, tag string, sessionToken string) {
	// 200 => Blogs found
	// 201 => No blog found for the specific tag
	// 202 => Invalid Session Token
	// 500 => Internal Server Error

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	URL := BASE_URL + "blogs/" + tag + "?sessionToken=" + sessionToken
	requestToBackend_server, err := http.NewRequestWithContext(ctxTimeout, "GET", URL, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create GET Request with Context")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}

	res, err := http.DefaultClient.Do(requestToBackend_server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")

			fmt.Fprint(ctx.Writer, REQUEST_TIMED_OUT_MESSAGE)
		}
		log.Println("Error: ", err)
		log.Println("Description: Cannot send POST request(with timeout(context)) to back-end server")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		responseJSONData, err := io.ReadAll(res.Body)

		defer res.Body.Close()

		if err != nil {
			log.Println("Error: ", err)
			log.Println("Decription: Cannot read body of response as JSON data")

			fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		}

		var responseBlogsList []model.Blog
		err = json.Unmarshal(responseJSONData, &responseBlogsList)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Decription: Cannot read body of response as JSON data")

			fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		}
		RenderHomePage(ctx, responseBlogsList, tag)
	} else if res.StatusCode == 201 {
		response := "No Blogs found for " + tag + " tag"
		log.Println(response)

		RenderHomePage(ctx, nil, tag)

	} else if res.StatusCode == 202 {
		RenderLoginPage(ctx, "Session Timed Out")
	} else if res.StatusCode == 500 {
		log.Println("Error: Back-end server has Internal Server Error")
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}
}

// Handlers
func DefaultRoute(ctx *gin.Context) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	http.ServeFile(ctx.Writer, ctx.Request, "./views/index.html")
}

func RenderInitPage(ctx *gin.Context) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	URL := BASE_URL + sessionToken

	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "GET", URL, nil)

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
			fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
			return
		}
		log.Println("Error: ", err)
		log.Println("Description: Cannot send DELETE request(with timeout(context)) to back-end server")
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	if res.StatusCode == 200 {
		fetchBlogsByTag(ctx, "All", sessionToken)
	} else if res.StatusCode == 500 {
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	} else {
		RenderLoginPage(ctx, "")
	}
}

func RenderRegsiterPage(ctx *gin.Context, message string) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	tmpl := template.Must(template.ParseFiles("./views/register.html"))
	err := tmpl.Execute(ctx.Writer, message)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RengerRegisterPage()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}
}

func RenderLoginPage(ctx *gin.Context, message string) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	tmpl := template.Must(template.ParseFiles("./views/login.html"))
	err := tmpl.Execute(ctx.Writer, message)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RenderLoginPage()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}
}

func RenderHomePage(ctx *gin.Context, blogs []model.Blog, tag string) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	data := map[string]interface{}{
		"Blogs":        blogs,
		"RequestedTag": tag,
		"TagsList":     tagsList,
	}

	tmpl := template.Must(template.ParseFiles("./views/home.html", "./views/blog.html"))
	err := tmpl.Execute(ctx.Writer, data)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RenderHomePage()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}
}

func RenderPostBlogPage(ctx *gin.Context, message string) {
	temp, err := template.ParseFiles("./views/add_blog.html")
	if err != nil {
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}
	data := map[string]interface{}{
		"TagsList": tagsList,
		"Message":  message,
	}
	temp.Execute(ctx.Writer, data)
}

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

		URL := BASE_URL + "register"

		// Create Request with Timeout
		requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "POST", URL, bytes.NewBuffer(userJSON))

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

	URL := BASE_URL + "login"

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "POST", URL, bytes.NewBuffer(userJSON))

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

	URL := BASE_URL + "login/" + sessionToken

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "DELETE", URL, nil)

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

			RenderLoginPage(ctx, "Request Timed Out")
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
		RenderLoginPage(ctx, "Internal Server Error")
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)

		RenderLoginPage(ctx, "Internal Server Error")
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

	URL := BASE_URL + "users/" + username

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "GET", URL, nil)

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

func GetBlogByTag(ctx *gin.Context) {
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
	// 200 => Title already found
	// 201 => Title not used
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

	sessioToken := getSessionTokenFromCookie(ctx.Request)

	URL := BASE_URL + "blogs/title/" + title + "?sessionToken=" + sessioToken

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "GET", URL, nil)

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
		fmt.Fprint(ctx.Writer, "Title is already used")
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
	blog.Tag = strings.Title(blog.Tag)
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
		URL := BASE_URL + "blogs?sessionToken=" + sessionToken

		// Create Request with Timeout
		requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "POST", URL, bytes.NewBuffer(blogJSON))

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

				RenderPostBlogPage(ctx, "Internal Server Error")
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

// TODO: Also send image to back-end server instead of storing it in client-side server
