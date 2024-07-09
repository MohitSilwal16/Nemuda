package utils

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/MohitSilwal16/Nemuda/chat/constants"
)

const SERVICE_BASE_URL = constants.SERVICE_BASE_URL

// Cookie handling
func SetSessionTokenInCookie(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:  "sessionToken",
		Value: token,
		Path:  "/",
	}

	http.SetCookie(w, cookie)
}

func GetSessionTokenFromCookie(r *http.Request) string {
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

func DeleteSessionTokenFromCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "sessionToken",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}

const INTERNAL_SERVER_ERROR_MESSAGE = "<script>alert('Internal Server Error');</script>"

func IsSessionTokenValid(sessionToken string) (bool, error) {
	// 200 => Session Token is Valid
	// 201 => Session Token is Invalid
	// 500 => Internal Server Error

	if sessionToken == "" {
		return false, nil
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	serviceURL := SERVICE_BASE_URL + sessionToken

	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "GET", serviceURL, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create DELETE Request with Context")

		return false, err
	}

	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")
			return false, err
		}
		log.Println("Error: ", err)
		log.Println("Description: Cannot send DELETE request(with timeout(context)) to back-end server")
		return false, err
	}

	if res.StatusCode == 200 {
		return true, nil
	} else if res.StatusCode == 201 {
		return false, nil
	} else if res.StatusCode == 500 {
		return false, errors.New("INTERNAL SERVER ERROR")
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)
		return false, errors.New("INTERNAL SERVER ERROR")
	}
}

func GetUsernameBySessionToken(sessionToken string) (string, error) {
	// 200 => Session Token is Valid
	// 201 => Session Token is Invalid
	// 500 => Internal Server Error

	if sessionToken == "" {
		return "", nil
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	serviceURL := SERVICE_BASE_URL + "get-users-by-sessionToken/" + sessionToken

	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "GET", serviceURL, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create DELETE Request with Context")
		return "", err
	}

	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")
			return "", err
		}
		log.Println("Error: ", err)
		log.Println("Description: Cannot send DELETE request(with timeout(context)) to back-end server")
		return "", err
	}

	if res.StatusCode == 200 {
		var responseDataStructure map[string]string

		err = json.NewDecoder(res.Body).Decode(&responseDataStructure)

		if err != nil {
			log.Println("Error:", err)
			log.Println("Description: Data from server is not in JSON format")
			return "", nil
		}
		username, ok := responseDataStructure["username"]

		if ok {
			return username, nil
		}
		return "", nil
	} else if res.StatusCode == 201 {
		return "", nil
	} else if res.StatusCode == 500 {
		return "", errors.New("INTERNAL SERVER ERROR")
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)
		return "", errors.New("INTERNAL SERVER ERROR")
	}
}