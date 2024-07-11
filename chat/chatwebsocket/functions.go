package chatwebsocket

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/MohitSilwal16/Nemuda/chat/constants"
	"github.com/gin-gonic/gin"
)

const SERVICE_BASE_URL = constants.SERVICE_BASE_URL

var INTERNAL_SERVER_ERROR_JSON = gin.H{
	"error": "Internal Server Error",
}

var REQUEST_TIMED_OUT_JSON = gin.H{
	"error": "Request Timed Out",
}

var SESSION_TIMED_OUT_JSON = gin.H{
	"error": "Invalid Session Token",
}

// Message when a single user uses websocket from two devices or tabs
var DOUBLE_CONNECTION = gin.H{
	"error": "Oops! Looks like you're double-dipping. Disconnect from one device or tab to keep things smooth",
}

var INTERNAL_SERVER_ERROR_BYTES, _ = json.Marshal(INTERNAL_SERVER_ERROR_JSON)
var REQUEST_TIMED_OUT_ERROR_BYTES, _ = json.Marshal(REQUEST_TIMED_OUT_JSON)
var SESSION_TIMED_OUT_ERROR_BYTES, _ = json.Marshal(SESSION_TIMED_OUT_JSON)
var DOUBLE_CONNECTION_ERROR_BYTES, _ = json.Marshal(DOUBLE_CONNECTION)

func getUsernameBySessionToken(sessionToken string) (string, error) {
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

// Send request to backend server to add message in DB
func addMessageInDB(client *Client, message Message) {
	// 200 => Messages added
	// 201 => Message data isn't in proper format(JSON)
	// 202 => Invalid Session Token
	// 203 => Invalid Receiver
	// 500 => Internal Server Error

	messageJSON, err := json.Marshal(message)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot parse Client's data to JSON")

		client.SendMessage <- INTERNAL_SERVER_ERROR_BYTES
		return
	}

	ctxTimeout, cancelFunction := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunction()

	serviceURL := SERVICE_BASE_URL + "messages" + "?sessionToken=" + client.SessionToken

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "POST", serviceURL, bytes.NewBuffer(messageJSON))

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot Create POST Request with Context")

		client.SendMessage <- INTERNAL_SERVER_ERROR_BYTES
		return
	}
	requestToBackend_Server.Header.Set("Content-Type", "application/json")

	// Send request(with timeout) to back-end server
	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")

			client.SendMessage <- REQUEST_TIMED_OUT_ERROR_BYTES
			return
		}
		log.Println("Error: ", err)
		log.Println("Description: Cannot send POST request(with timeout(context)) to back-end server")

		client.SendMessage <- INTERNAL_SERVER_ERROR_BYTES

		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		return
	} else if res.StatusCode == 201 {
		log.Println("Error: Message data isn't in proper format")
		client.SendMessage <- INTERNAL_SERVER_ERROR_BYTES
	} else if res.StatusCode == 202 {
		client.SendMessage <- SESSION_TIMED_OUT_ERROR_BYTES
		client.Router.Unregister <- client
	} else if res.StatusCode == 500 {
		log.Println("Error: Back-end server has Internal Server Error")
		client.SendMessage <- INTERNAL_SERVER_ERROR_BYTES
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)
		client.SendMessage <- INTERNAL_SERVER_ERROR_BYTES
	}
}
