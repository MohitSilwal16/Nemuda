package chatwebsocket

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MohitSilwal16/Nemuda/chat/constants"
	"github.com/gin-gonic/gin"
)

const SERVICE_BASE_URL = constants.SERVICE_BASE_URL

var INTERNAL_SERVER_ERROR_JSON = gin.H{
	"message": "Internal Server Error",
}

var REQUEST_TIMED_OUT_JSON = gin.H{
	"message": "Request Timed Out",
}

var INTERNAL_SERVER_ERROR_BYTES, _ = json.Marshal(INTERNAL_SERVER_ERROR_JSON)
var REQUEST_TIMED_OUT_ERROR_BYTES, _ = json.Marshal(REQUEST_TIMED_OUT_JSON)

// Send request to backend server to add message in DB
func AddMessageInDB(client *Client, message Message) {
	// 200 => Messages added
	// 201 => Message data isn't in proper format(JSON)
	// 202 => Invalid Session Token
	// 203 => Invalid Receiver
	// 500 => Internal Server Error

	messageJSON, err := json.Marshal(message)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot parse Client's data to JSON")

		client.Send <- INTERNAL_SERVER_ERROR_BYTES
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

		client.Send <- INTERNAL_SERVER_ERROR_BYTES
		return
	}
	requestToBackend_Server.Header.Set("Content-Type", "application/json")

	// Send request(with timeout) to back-end server
	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err)
			log.Println("Description: Back-end server didn't responsed in given time")

			client.Send <- REQUEST_TIMED_OUT_ERROR_BYTES
			return
		}
		log.Println("Error: ", err)
		log.Println("Description: Cannot send POST request(with timeout(context)) to back-end server")

		client.Send <- INTERNAL_SERVER_ERROR_BYTES

		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		return
	} else if res.StatusCode == 201 {
		log.Println("Error: Message data isn't in proper format")
		client.Send <- INTERNAL_SERVER_ERROR_BYTES
	} else if res.StatusCode == 500 {
		log.Println("Error: Back-end server has Internal Server Error")
		client.Send <- INTERNAL_SERVER_ERROR_BYTES
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode)
		client.Send <- INTERNAL_SERVER_ERROR_BYTES
	}
}
