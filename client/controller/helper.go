package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Nemuda/client/model"
	"github.com/gin-gonic/gin"
)

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

func isBlogLiked(ctx *gin.Context, title string, sessionToken string) (bool, error) {
	// 200 => Blog liked
	// 201 => Blog not liked
	// 202 => Blog not found
	// 203 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	if len(title) < 5 {
		fmt.Fprint(ctx.Writer, BLOG_NOT_FOUND_MESSAGE)
		return false, errors.New("LENGHT OF TITLE IS < 5")
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	serviceURL := BASE_URL + "blogs/like/" + title + "?sessionToken=" + sessionToken

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "GET", serviceURL, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create GET Request with Context")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return false, errors.New("CANNOT CREATE GET REQUEST WITH CONTEXT")
	}

	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")
			fmt.Fprint(ctx.Writer, REQUEST_TIMED_OUT_MESSAGE)
			return false, errors.New("REQUEST TIMED OUT")
		} else {
			log.Println("Error: ", err)
			log.Println("Description: Cannot send GET request(with timeout(context)) to back-end server")
			fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
			return false, errors.New("CANNOT SEND GET REQUEST WITH CONTEXT")
		}
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		return true, nil
	} else if res.StatusCode == 201 {
		return false, nil
	} else if res.StatusCode == 202 {
		fmt.Fprint(ctx.Writer, BLOG_NOT_FOUND_MESSAGE)
		return false, errors.New("BLOG NOT FOUND")
	} else if res.StatusCode == 203 {
		RenderLoginPage(ctx, "Session Timed Out")
		return false, errors.New("SESSION TIMED OUT")
	} else if res.StatusCode == 500 {
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return false, errors.New("INTERNAL SERVER ERROR")
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return false, errors.New("BUG: UNEXPECTED STATUS CODE")
	}
}

func fetchBlogsByTag(ctx *gin.Context, tag string, sessionToken string) {
	// 200 => Blogs found
	// 201 => No blog found for the specific tag
	// 202 => Invalid Session Token
	// 500 => Internal Server Error

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	serviceURL := BASE_URL + "blogs/" + tag + "?sessionToken=" + sessionToken
	requestToBackend_server, err := http.NewRequestWithContext(ctxTimeout, "GET", serviceURL, nil)

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
			return
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

func getBlogByTitle(ctx *gin.Context, title string, message string) {
	// 200 => Blog found
	// 201 => Blog not found
	// 202 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	serviceURL := BASE_URL + "blogs/title/" + title + "?sessionToken=" + sessionToken

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
		responseJSONData, err := io.ReadAll(res.Body)

		defer res.Body.Close()

		if err != nil {
			log.Println("Error: ", err)
			log.Println("Decription: Cannot read body of response as JSON data")

			fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
			return
		}

		var blog model.Blog
		err = json.Unmarshal(responseJSONData, &blog)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Decription: Cannot read body of response as JSON data")

			fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
			return
		}
		RenderGetBlogPage(ctx, blog, message)
	} else if res.StatusCode == 201 {
		log.Println("Error: Blog not found")
		fmt.Fprint(ctx.Writer, BLOG_NOT_FOUND_MESSAGE)
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

func isBlogEditableDeletable(ctx *gin.Context, title string, sessionToken string) (bool, error) {
	// 200 => User can update/delete blog
	// 201 => User cannot update/delete blog
	// 202 => Blog not found
	// 203 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	if len(title) < 5 {
		fmt.Fprint(ctx.Writer, BLOG_NOT_FOUND_MESSAGE)
		return false, errors.New("LENGHT OF TITLE IS < 5")
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	serviceURL := BASE_URL + "blogs/updatable_deletable?sessionToken=" + sessionToken + "&title=" + title
	serviceURL = strings.Replace(serviceURL, " ", "%20", -1)

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "GET", serviceURL, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create GET Request with Context")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return false, errors.New("CANNOT CREATE GET REQUEST WITH CONTEXT")
	}

	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")
			fmt.Fprint(ctx.Writer, REQUEST_TIMED_OUT_MESSAGE)
			return false, errors.New("REQUEST TIMED OUT")
		} else {
			log.Println("Error: ", err)
			log.Println("Description: Cannot send GET request(with timeout(context)) to back-end server")
			fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
			return false, errors.New("CANNOT SEND GET REQUEST WITH CONTEXT")
		}
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		return true, nil
	} else if res.StatusCode == 201 {
		return false, nil
	} else if res.StatusCode == 202 {
		fmt.Fprint(ctx.Writer, BLOG_NOT_FOUND_MESSAGE)
		return false, errors.New("BLOG NOT FOUND")
	} else if res.StatusCode == 203 {
		RenderLoginPage(ctx, "Session Timed Out")
		return false, errors.New("SESSION TIMED OUT")
	} else if res.StatusCode == 500 {
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return false, errors.New("INTERNAL SERVER ERROR")
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return false, errors.New("BUG: UNEXPTEC STATUS CODE")
	}
}

func fetchBlogByTitleAndReturn(ctx *gin.Context, title string) (model.Blog, error) {
	// 200 => Blog found
	// 201 => Blog not found
	// 202 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	var b model.Blog

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	serviceURL := BASE_URL + "blogs/title/" + title + "?sessionToken=" + sessionToken

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "GET", serviceURL, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create GET Request with Context")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return b, err
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
		return b, err
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		responseJSONData, err := io.ReadAll(res.Body)

		defer res.Body.Close()

		if err != nil {
			log.Println("Error: ", err)
			log.Println("Decription: Cannot read body of response as JSON data")

			fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
			return b, err
		}

		var blog model.Blog
		err = json.Unmarshal(responseJSONData, &blog)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Decription: Cannot read body of response as JSON data")

			fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
			return b, err
		}
		return blog, nil
	} else if res.StatusCode == 201 {
		log.Println("Error: Blog not found")
		fmt.Fprint(ctx.Writer, BLOG_NOT_FOUND_MESSAGE)
		return b, errors.New("BLOG NOT FOUND")
	} else if res.StatusCode == 202 {
		RenderLoginPage(ctx, "Session Timed Out")
		return b, errors.New("SESSION TIMED OUT")
	} else if res.StatusCode == 500 {
		log.Println("Error: Back-end server has Internal Server Error")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return b, errors.New("INTERNAL SERVER ERROR")
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return b, errors.New("UNEXPECTED STATUS CODE")
	}
}
