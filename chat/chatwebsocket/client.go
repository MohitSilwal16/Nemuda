package chatwebsocket

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// Define the date format
const dateFormat = "2006-01-02 15:04:05"

func (client *Client) readMessages() {
	defer func() {
		client.Connection.Close()
		client.Router.Unregister <- client
	}()

	client.Connection.SetReadLimit(maxMessageSize)
	client.Connection.SetReadDeadline(time.Now().Add(pongWaitTime))
	client.Connection.SetPongHandler(func(appData string) error {
		client.Connection.SetReadDeadline(time.Now().Add(pongWaitTime))
		return nil
	})

	for {
		_, msg, err := client.Connection.ReadMessage()
		if err != nil {
			log.Println("Error:", err)
			break
		}

		log.Println("Message:", string(msg))

		var wsMessage WSMessage
		reader := bytes.NewReader(msg)
		err = json.NewDecoder(reader).Decode(&wsMessage)

		if err != nil {
			log.Println("Error:", err)
			log.Println("Error while decoding")
			return
		}

		currentTime := time.Now().Format(dateFormat)

		client.Router.SendMessage <- Message{
			Sender:         client.Username,
			MessageContent: wsMessage.Message,
			Receiver:       wsMessage.Receiver,
			DateTime:       currentTime,
			Status:         "Sent",
		}
	}
}

func (client *Client) writeMessage() {
	ticker := time.NewTicker(pingWaitTime)
	defer func() {
		client.Connection.Close()
		client.Router.Unregister <- client
	}()

	for {
		select {
		case msg, ok := <-client.Send:
			if !ok {
				// Hub is closed
				client.Connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.Connection.NextWriter(websocket.TextMessage)

			if err != nil {
				log.Println("Error:", err)
				log.Println("Description: Error from client.Connection.NextWriter()")
				return
			}

			w.Write(msg)

			if err = w.Close(); err != nil {
				log.Println("Error:", err)
				log.Println("Description: Error from client.Connection.NextWriter()")
				return
			}
		case <-ticker.C:
			if err := client.Connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("No response of ping from client")
				return
			}
		}
	}
}
