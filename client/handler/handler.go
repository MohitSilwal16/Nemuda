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
	"text/template"
	"time"

	"github.com/Nemuda/client/model"
	"github.com/Nemuda/client/utils"
	"github.com/gin-gonic/gin"
)

const BASE_URL = "http://localhost:8080/"

func setCookie(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:  "sessionToken",
		Value: token,
		Path:  "/",
	}

	http.SetCookie(w, cookie)
}

func getCookie(r *http.Request) string {
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

func DefaultRoute(ctx *gin.Context) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	http.ServeFile(ctx.Writer, ctx.Request, "./views/index.html")
}

func RenderInitPage(ctx *gin.Context) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	sessionToken := getCookie(ctx.Request)

	// if sessionToken != "" && checkDuplicateToken(sessionToken) {
	if sessionToken != "" {
		RenderHomePage(ctx)
	} else {
		RenderLoginPage(ctx)
	}
}

func RenderRegsiterPage(ctx *gin.Context) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	tmpl := template.Must(template.ParseFiles("./views/register.html"))
	err := tmpl.Execute(ctx.Writer, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RengerRegisterPage()")

		fmt.Fprint(ctx.Writer, "<script>Internal Server Error</script>")
	}
}

func RenderLoginPage(ctx *gin.Context) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	tmpl := template.Must(template.ParseFiles("./views/login.html"))
	err := tmpl.Execute(ctx.Writer, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RenderLoginPage()")

		fmt.Fprint(ctx.Writer, "<script>Internal Server Error</script>")
	}
}

func RenderHomePage(ctx *gin.Context) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	tmpl := template.Must(template.ParseFiles("./views/home.html", "./views/blog.html"))
	err := tmpl.Execute(ctx.Writer, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RenderHomePage()")

		fmt.Fprint(ctx.Writer, "<script>Internal Server Error</script>")
	}
}

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

		tmpl := template.Must(template.ParseFiles("./views/register.html"))
		err := tmpl.Execute(ctx.Writer, response)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in Register()")
		}
		return
	}

	if user.Username == "" || user.Password == "" {
		response := "Username or Password is Empty"

		tmpl := template.Must(template.ParseFiles("./views/register.html"))
		err := tmpl.Execute(ctx.Writer, response)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in Register()")
			return
		}
	} else if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(user.Username) {
		response := "Username should be alphanumeric b'twin 5-20 chars"

		tmpl := template.Must(template.ParseFiles("./views/register.html"))
		err := tmpl.Execute(ctx.Writer, response)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in Register()")
			return
		}
	} else if !utils.IsPasswordInFormat(user.Password) {
		response := "Password: 8+ chars, lower & upper case, digit, symbol"

		tmpl := template.Must(template.ParseFiles("./views/register.html"))
		err := tmpl.Execute(ctx.Writer, response)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in Register()")
			return
		}
	} else {
		userJSON, err := json.Marshal(user)

		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Cannot parse Client's data to JSON")
			response := "Cannot Parse data to JSON"

			tmpl := template.Must(template.ParseFiles("./views/register.html"))
			err := tmpl.Execute(ctx.Writer, response)
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Error in tmpl.Execute() in Register()")
			}
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
			response := "Internal Server Error"

			tmpl := template.Must(template.ParseFiles("./views/register.html"))
			err := tmpl.Execute(ctx.Writer, response)
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Error in tmpl.Execute() in Register()")
			}
			return
		}
		requestToBackend_Server.Header.Set("Content-Type", "application/json")

		client := http.Client{
			Timeout: 3 * time.Second,
		}

		// Send request(with timeout) to back-end server
		res, err := client.Do(requestToBackend_Server)

		if err != nil {
			if ctxTimeout.Err() == context.DeadlineExceeded {
				log.Println("Error: ", err)
				log.Println("Description: Back-end server didn't responsed in given time")
			} else {
				log.Println("Error: ", err)
				log.Println("Description: Cannot send POST request(with timeout(context)) to back-end server")
			}
			response := "Internal Server Error"

			tmpl := template.Must(template.ParseFiles("./views/register.html"))
			tmpl.Execute(ctx.Writer, response)
			return
		}

		defer res.Body.Close()

		if res.StatusCode == 200 {
			responseJSONData, err := io.ReadAll(res.Body)

			if err != nil {
				log.Println("Error: ", err)
				log.Println("Decription: Cannot read body of response as JSON data")
				response := "Internal Server Error"

				tmpl := template.Must(template.ParseFiles("./views/register.html"))
				err := tmpl.Execute(ctx.Writer, response)
				if err != nil {
					log.Println("Error: ", err)
					log.Println("Description: Error in tmpl.Execute() in Register()")
				}
				return
			}

			var responseDataStructure map[string]string
			err = json.Unmarshal(responseJSONData, &responseDataStructure)
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Decription: Cannot read body of response")
				response := "Internal Server Error"

				tmpl := template.Must(template.ParseFiles("./views/register.html"))
				err := tmpl.Execute(ctx.Writer, response)
				if err != nil {
					log.Println("Error: ", err)
					log.Println("Description: Error in tmpl.Execute() in Register()")
				}
				return
			}

			sessionToken := responseDataStructure["sessionToken"]
			responseMessage := responseDataStructure["message"]

			if responseMessage != "Registered Successfully" {
				log.Println("Error: Back-end server didn't send the message of 'Registered Successfully'")
				log.Println("There might be some problem in back-end server")
				response := "Internal Server Error"

				tmpl := template.Must(template.ParseFiles("./views/register.html"))
				err := tmpl.Execute(ctx.Writer, response)
				if err != nil {
					log.Println("Error: ", err)
					log.Println("Description: Error in tmpl.Execute() in Register()")
				}
				return
			}

			// No Session Token is sent from back-end server
			if sessionToken == "" {
				log.Println("Error: Session Token not provided by back-end server")
				response := "Internal Server Error"

				tmpl := template.Must(template.ParseFiles("./views/register.html"))
				err := tmpl.Execute(ctx.Writer, response)
				if err != nil {
					log.Println("Error: ", err)
					log.Println("Description: Error in tmpl.Execute() in Register()")
				}
				return
			}
			// Save session token in cookie of user
			setCookie(ctx.Writer, sessionToken)

			// TODO: Fetch blogs from server
			tmpl := template.Must(template.ParseFiles("./views/home.html"))
			err = tmpl.Execute(ctx.Writer, nil)
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Error in tmpl.Execute() in Register()")
				return
			}
		} else if res.StatusCode == 206 {
			tmpl := template.Must(template.ParseFiles("./views/register.html"))
			err := tmpl.Execute(ctx.Writer, "Username is already used")
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Error in tmpl.Execute() in Register()")
				return
			}
		} else if res.StatusCode == 500 {
			log.Println("Error: Back-end server has Internal Server Error")

			tmpl := template.Must(template.ParseFiles("./views/register.html"))
			err := tmpl.Execute(ctx.Writer, "Internal Server Error")
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Error in tmpl.Execute() in Register()")
				return
			}
		} else {
			// We already filtered data above & checked whether
			// Username, Password is in correct format & not empty
			// Also we ensured that data is in correct format(JSON)
			log.Println("Error: Client side already filtered data")
			log.Println("But still we get response of inconsistent data or unfiltered data")
			log.Println("Either there's some logic issue in back-end server or issue in filtering data in client server(front-end server)")

			tmpl := template.Must(template.ParseFiles("./views/register.html"))
			err := tmpl.Execute(ctx.Writer, "Some Problem Occurred")
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Error in tmpl.Execute() in Register()")
				return
			}
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
		response := "Data isn't in JSON format"

		tmpl := template.Must(template.ParseFiles("./views/login.html"))
		err := tmpl.Execute(ctx.Writer, response)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in Login()")
		}
		return
	}

	if user.Username == "" || user.Password == "" {
		response := "Username or Password is Empty"

		tmpl := template.Must(template.ParseFiles("./views/login.html"))
		err := tmpl.Execute(ctx.Writer, response)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in Login()")
			return
		}
	} else if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(user.Username) {
		response := "Username should be alphanumeric b'twin 5-20 chars"

		tmpl := template.Must(template.ParseFiles("./views/login.html"))
		err := tmpl.Execute(ctx.Writer, response)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in Login()")
			return
		}
	} else if !utils.IsPasswordInFormat(user.Password) {
		// Don't want to give idea to anonymous user idea about password pattern
		response := "Invalid Credentials"

		tmpl := template.Must(template.ParseFiles("./views/login.html"))
		err := tmpl.Execute(ctx.Writer, response)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in Login()")
			return
		}
	} else {
		userJSON, err := json.Marshal(user)

		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Cannot parse Client's data to JSON")
			response := "Cannot Parse data to JSON"

			tmpl := template.Must(template.ParseFiles("./views/login.html"))
			err := tmpl.Execute(ctx.Writer, response)
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Error in tmpl.Execute() in Login()")
			}
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
			response := "Internal Server Error"

			tmpl := template.Must(template.ParseFiles("./views/login.html"))
			err := tmpl.Execute(ctx.Writer, response)
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Error in tmpl.Execute() in Login()")
			}
			return
		}
		requestToBackend_Server.Header.Set("Content-Type", "application/json")

		client := http.Client{
			Timeout: 3 * time.Second,
		}

		// Send request(with timeout) to back-end server
		res, err := client.Do(requestToBackend_Server)

		if err != nil {
			if ctxTimeout.Err() == context.DeadlineExceeded {
				log.Println("Error: ", err)
				log.Println("Description: Back-end server didn't responsed in given time")
			} else {
				log.Println("Error: ", err)
				log.Println("Description: Cannot send POST request(with timeout(context)) to back-end server")
			}
			response := "Internal Server Error"

			tmpl := template.Must(template.ParseFiles("./views/login.html"))
			err := tmpl.Execute(ctx.Writer, response)
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Error in tmpl.Execute() in Login()")
			}
			return
		}

		defer res.Body.Close()

		if res.StatusCode == 200 {
			responseJSONData, err := io.ReadAll(res.Body)

			if err != nil {
				log.Println("Error: ", err)
				log.Println("Decription: Cannot read body of response")
				response := "Internal Server Error"

				tmpl := template.Must(template.ParseFiles("./views/login.html"))
				err := tmpl.Execute(ctx.Writer, response)
				if err != nil {
					log.Println("Error: ", err)
					log.Println("Description: Error in tmpl.Execute() in Login()")
				}
				return
			}

			var responseDataStructure map[string]string
			err = json.Unmarshal(responseJSONData, &responseDataStructure)
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Decription: Cannot read body of response as JSON data")
				response := "Internal Server Error"

				tmpl := template.Must(template.ParseFiles("./views/login.html"))
				err := tmpl.Execute(ctx.Writer, response)
				if err != nil {
					log.Println("Error: ", err)
					log.Println("Description: Error in tmpl.Execute() in Login()")
				}
				return
			}

			sessionToken := responseDataStructure["sessionToken"]
			responseMessage := responseDataStructure["message"]

			if responseMessage != "Login Successful" {
				log.Println("Error: Back-end server didn't send the message of 'Login Successful'")
				log.Println("There might be some problem in back-end server")
				response := "Internal Server Error"

				tmpl := template.Must(template.ParseFiles("./views/login.html"))
				err := tmpl.Execute(ctx.Writer, response)
				if err != nil {
					log.Println("Error: ", err)
					log.Println("Description: Error in tmpl.Execute() in Login()")
				}
				return
			}

			// No Session Token is sent from back-end server
			if sessionToken == "" {
				log.Println("Error: Session Token not provided by back-end server")
				response := "Internal Server Error"

				tmpl := template.Must(template.ParseFiles("./views/login.html"))
				err := tmpl.Execute(ctx.Writer, response)
				if err != nil {
					log.Println("Error: ", err)
					log.Println("Description: Error in tmpl.Execute() in Login()")
				}
				return
			}
			// Save session token in cookie of user
			setCookie(ctx.Writer, sessionToken)

			// TODO: Fetch blogs from server
			tmpl := template.Must(template.ParseFiles("./views/home.html"))
			err = tmpl.Execute(ctx.Writer, nil)
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Error in tmpl.Execute() in Login()")
				return
			}

		} else if res.StatusCode == 205 {
			tmpl := template.Must(template.ParseFiles("./views/login.html"))
			err = tmpl.Execute(ctx.Writer, "User doesn't exists")
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Error in tmpl.Execute() in Login()")
				return
			}
		} else if res.StatusCode == 206 {
			tmpl := template.Must(template.ParseFiles("./views/login.html"))
			err = tmpl.Execute(ctx.Writer, "Invalid Credentials")
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Error in tmpl.Execute() in Login()")
				return
			}
		} else if res.StatusCode == 500 {
			log.Println("Error: Back-end server has Internal Server Error")

			tmpl := template.Must(template.ParseFiles("./views/login.html"))
			err = tmpl.Execute(ctx.Writer, "Internal Server Error")
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Error in tmpl.Execute() in Login()")
				return
			}
		} else {
			// We already filtered data above & checked whether
			// Username, Password is in correct format & not empty
			// Also we ensured that data is in correct format(JSON)
			log.Println("Error: Client side already filtered data")
			log.Println("But still we get response of inconsistent data or unfiltered data")
			log.Println("Either there's some logic issue in back-end server or issue in filtering data in client server(front-end server)")

			tmpl := template.Must(template.ParseFiles("./views/login.html"))
			err = tmpl.Execute(ctx.Writer, "Some Problem Occurred")
			if err != nil {
				log.Println("Error: ", err)
				log.Println("Description: Error in tmpl.Execute() in Login()")
				return
			}
		}
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

	sessionToken := getCookie(ctx.Request)

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	URL := BASE_URL + "login/" + sessionToken

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "DELETE", URL, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create DELETE Request with Context")
		response := "<script>Internal Server Error</script>"

		tmpl := template.Must(template.ParseFiles("./views/home.html"))
		err := tmpl.Execute(ctx.Writer, response)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in Logout()")
		}
		return
	}

	client := http.Client{
		Timeout: 3 * time.Second,
	}

	// Send request(with timeout) to back-end server
	res, err := client.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")
		} else {
			log.Println("Error: ", err)
			log.Println("Description: Cannot send DELETE request(with timeout(context)) to back-end server")
		}
		response := "<script>Internal Server Error</script>"

		tmpl := template.Must(template.ParseFiles("./views/home.html"))
		err := tmpl.Execute(ctx.Writer, response)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in Logout()")
			return
		}
	}

	defer res.Body.Close()

	if res.StatusCode == 200 || res.StatusCode == 201 {
		if res.StatusCode == 201 {
			log.Println("Invalid Session Token")
		}
		deleteSessionTokenFromCookie(ctx.Writer)

		// Even if session token is invalid log out user
		tmpl := template.Must(template.ParseFiles("./views/login.html"))
		err = tmpl.Execute(ctx.Writer, nil)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in Logout()")
			return
		}
	} else if res.StatusCode == 500 {
		response := "<script>Internal Server Error</script>"

		tmpl := template.Must(template.ParseFiles("./views/home.html"))
		err = tmpl.Execute(ctx.Writer, response)
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in Logout()")
			return
		}
	}
}

func SearchUserForRegistration(ctx *gin.Context) {
	// 200 => User found (Username is already used)
	// 201 => User not found (Username is not used yet)
	// 500 => Internal Server Error

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	username := ctx.Query("username")
	log.Println("Username: ", username)

	if len(username) < 5 {
		log.Println("Length of username < 5")
		tmpl, err := template.New("t").Parse("")
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in SearchUserForRegistration()")
			return
		}
		tmpl.Execute(ctx.Writer, nil)
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

		tmpl, err := template.New("t").Parse("")
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in SearchUserForRegistration")
			return
		}
		tmpl.Execute(ctx.Writer, nil)
		return
	}

	client := http.Client{
		Timeout: time.Second,
	}

	res, err := client.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")
		} else {
			log.Println("Error: ", err)
			log.Println("Description: Cannot send GET request(with timeout(context)) to back-end server")
		}

		tmpl, err := template.New("t").Parse("")
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in SearchUserForRegistration")
			return
		}
		tmpl.Execute(ctx.Writer, nil)
		return
	}

	if res.StatusCode == 200 {
		tmpl, err := template.New("t").Parse("Username is already used")
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in SearchUserForRegistration")
			return
		}
		tmpl.Execute(ctx.Writer, nil)
	} else if res.StatusCode == 201 {
		log.Println("User not found")
		tmpl, err := template.New("t").Parse("")
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in SearchUserForRegistration")
			return
		}
		tmpl.Execute(ctx.Writer, nil)
	} else if res.StatusCode == 500 {
		log.Println("Error: Back-end server has Internal Server Error")

		tmpl, err := template.New("t").Parse("")
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Error in tmpl.Execute() in SearchUserForRegistration")
			return
		}
		tmpl.Execute(ctx.Writer, nil)
	} else {
		tmpl, err := template.New("t").Parse("")
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Back-end Server is sending some other status code")
			return
		}
		tmpl.Execute(ctx.Writer, nil)
	}
}
