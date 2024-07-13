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
)

const SERVICE_BASE_URL = constants.SERVICE_BASE_URL

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
		log.Println("Error: ", err, "\nSource: getUsernameBySessionToken()")
		log.Println("Description: Cannot Create DELETE Request with Context")
		return "", err
	}

	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err, "\nSource: getUsernameBySessionToken()")
			log.Println("Description: Back-end server didn't responsed in given time")
			return "", err
		}
		log.Println("Error: ", err, "\nSource: getUsernameBySessionToken()")
		log.Println("Description: Cannot send DELETE request(with timeout(context)) to back-end server")
		return "", err
	}

	if res.StatusCode == 200 {
		var responseDataStructure map[string]string

		err = json.NewDecoder(res.Body).Decode(&responseDataStructure)

		if err != nil {
			log.Println("Error:", err, "\nSource: getUsernameBySessionToken()")
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
		log.Println("Error: Internal Server Error from Server\nSource: getUsernameBySessionToken()")
		return "", errors.New("INTERNAL SERVER ERROR")
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode, "\nSource: getUsernameBySessionToken()")
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

	if client == nil {
		log.Println("Null client\nSource: addMessageInDB()")
		return
	}

	messageJSON, err := json.Marshal(message)

	if err != nil {
		log.Println("Error: ", err, "\nSource: addMessageInDB()")
		log.Println("Description: Cannot parse Client's data to JSON")

		client.SendMessage <- INTERNAL_SERVER_ERROR_MESSAGE
		return
	}

	ctxTimeout, cancelFunction := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunction()

	serviceURL := SERVICE_BASE_URL + "messages" + "?sessionToken=" + client.SessionToken

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "POST", serviceURL, bytes.NewBuffer(messageJSON))

	if err != nil {
		log.Println("Error: ", err, "\nSource: addMessageInDB()")
		log.Println("Description: Cannot Create POST Request with Context")

		client.SendMessage <- INTERNAL_SERVER_ERROR_MESSAGE
		return
	}
	requestToBackend_Server.Header.Set("Content-Type", "application/json")

	// Send request(with timeout) to back-end server
	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err, "\nSource: addMessageInDB()")
			log.Println("Description: Back-end server didn't responsed in given time")

			client.SendMessage <- REQUEST_TIMED_OUT_MESSAGE
			return
		}
		log.Println("Error: ", err, "\nSource addMessageInDB()")
		log.Println("Description: Cannot send POST request(with timeout(context)) to back-end server")

		client.SendMessage <- INTERNAL_SERVER_ERROR_MESSAGE
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		return
	} else if res.StatusCode == 201 {
		log.Println("Error Response From Server: Message sent data isn't in proper format\nSource: addMessageInDB()")
		client.SendMessage <- INTERNAL_SERVER_ERROR_MESSAGE
	} else if res.StatusCode == 202 {
		client.SendMessage <- SESSION_TIMED_OUT_MESSAGE
		client.Router.Unregister <- client
	} else if res.StatusCode == 500 {
		log.Println("Error: Back-end server has Internal Server Error\nSource: addMessageInDB()")
		client.SendMessage <- INTERNAL_SERVER_ERROR_MESSAGE
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode, "Source: addMessageInDB()")
		client.SendMessage <- INTERNAL_SERVER_ERROR_MESSAGE
	}
}

func isSessionTokenValid(client *Client, sessionToken string) bool {
	// 200 => Session Token is Valid
	// 201 => Session Token is Invalid
	// 500 => Internal Server Error

	if client == nil {
		log.Println("Null client\nSource: isSessionTokenValid()")
		return false
	}

	if sessionToken == "" {
		client.SendMessage <- SESSION_TIMED_OUT_MESSAGE
		return false
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelFunc()

	serviceURL := SERVICE_BASE_URL + sessionToken

	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "GET", serviceURL, nil)

	if err != nil {
		log.Println("Error: ", err, "\nSource: isSessionTokenValid()")
		log.Println("Description: Cannot Create DELETE Request with Context while validating session")
		client.SendMessage <- INTERNAL_SERVER_ERROR_MESSAGE
		return false
	}

	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err, "\nSource: isSessionTokenValid()")
			log.Println("Description: Back-end server didn't responsed in given time")
			client.SendMessage <- REQUEST_TIMED_OUT_MESSAGE
			return false
		}
		log.Println("Error: ", err, "\nSource: isSessionTokenValid()")
		log.Println("Description: Cannot send DELETE request(with timeout(context)) to back-end server")
		client.SendMessage <- INTERNAL_SERVER_ERROR_MESSAGE
		return false
	}

	if res.StatusCode == 200 {
		return true
	} else if res.StatusCode == 201 {
		client.SendMessage <- SESSION_TIMED_OUT_MESSAGE
		return false
	} else if res.StatusCode == 500 {
		log.Println("Error: Back-end server has Internal Server Error\nSource: isSessionTokenValid()")
		client.SendMessage <- INTERNAL_SERVER_ERROR_MESSAGE
		return false
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode, "\nSource: isSessionTokenValid()")
		client.SendMessage <- INTERNAL_SERVER_ERROR_MESSAGE
		return false
	}
}

func changeStatusOfMessage(client *Client, newStatus string, receiver string) {
	// 200 => Status Changed
	// 201 => NOTHING
	// 202 => Invalid Session Token
	// 203 => Invalid Status
	// 500 => Internal Server Error

	if client == nil {
		log.Println("Null client\nSource: changeStatusOfMessage()")
		return
	}

	ctxTimeout, cancelFunction := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunction()

	serviceURL := SERVICE_BASE_URL + "messages/" + client.SessionToken + "?newStatus=" + newStatus + "&receiver=" + receiver

	// Create Request with Timeout
	requestToBackend_Server, err := http.NewRequestWithContext(ctxTimeout, "PUT", serviceURL, nil)

	if err != nil {
		log.Println("Error: ", err, "\nSource: changeStatusOfMessage()")
		log.Println("Description: Cannot Create POST Request with Context")

		client.SendMessage <- INTERNAL_SERVER_ERROR_MESSAGE
		return
	}
	requestToBackend_Server.Header.Set("Content-Type", "application/json")

	// Send request(with timeout) to back-end server
	res, err := http.DefaultClient.Do(requestToBackend_Server)

	if err != nil {
		if ctxTimeout.Err() == context.DeadlineExceeded {
			log.Println("Error: ", err, "\nSource: changeStatusOfMessage()")
			log.Println("Description: Back-end server didn't responsed in given time")

			client.SendMessage <- REQUEST_TIMED_OUT_MESSAGE
			return
		}
		log.Println("Error: ", err, "\nSource changeStatusOfMessage()")
		log.Println("Description: Cannot send POST request(with timeout(context)) to back-end server")

		client.SendMessage <- INTERNAL_SERVER_ERROR_MESSAGE
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
	} else if res.StatusCode == 202 {
		client.SendMessage <- SESSION_TIMED_OUT_MESSAGE
		client.Router.Unregister <- client
	} else if res.StatusCode == 500 {
		log.Println("Error: Back-end server has Internal Server Error\nSource: changeStatusOfMessage()")
		client.SendMessage <- INTERNAL_SERVER_ERROR_MESSAGE
	} else {
		log.Println("Bug: Unexpected Status Code ", res.StatusCode, "Source: changeStatusOfMessage()")
		client.SendMessage <- INTERNAL_SERVER_ERROR_MESSAGE
	}
}
