package chatwebsocket

import (
	"encoding/json"
	"log"
	"time"

	"github.com/MohitSilwal16/Nemuda/chat/utils"
	"github.com/gorilla/websocket"
)

// Define the date format
const dateFormat = "2006-01-02 15:04:05"

func (client *Client) readMessages() {
	defer func() {
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
			log.Println("Error:", err, " User:", client.Username)
			return
		}

		var wsMessage WSMessage
		err = json.Unmarshal(msg, &wsMessage)
		if err != nil {
			log.Println("Error:", err)
			log.Println("Error decoding message while reading message from user")
			return
		}

		isSessionTokenValid(client)

		if wsMessage.MessageType == "Read" {
			if client.Username == wsMessage.Receiver {
				continue
			}
			client.Router.SendMessage <- Message{
				Sender:      client.Username,
				Receiver:    wsMessage.Receiver,
				MessageType: "Read",
			}
			continue
		}
		if wsMessage.Message == "" {
			client.SendMessage <- Message{
				Error: "Empty Message",
			}
			continue
		}

		var isMalicious bool
		wsMessage.Message, isMalicious = utils.SanitizeMessage(wsMessage.Message)

		if isMalicious {
			log.Println(wsMessage.Message)
			client.SendMessage <- Message{
				Error: "Whoops! Looks like someone tried to sprinkle some XSS magic. Nice attempt son",
			}
			continue
		}

		if len(wsMessage.Message) > 100 {
			client.SendMessage <- Message{
				Error: "Message should be less than 100 chars",
			}
			continue
		}

		currentTime := time.Now().Format(dateFormat)
		client.Router.SendMessage <- Message{
			Sender:         client.Username,
			MessageContent: wsMessage.Message,
			Receiver:       wsMessage.Receiver,
			DateTime:       currentTime,
			Status:         "Sent",
			MessageType:    "Message",
		}
	}
}

func (client *Client) writeMessage() {
	ticker := time.NewTicker(pingWaitTime)
	defer func() {
		client.Router.Unregister <- client
	}()

	for {
		select {
		case msg, ok := <-client.SendMessage:

			if !ok {
				// Router is closed
				log.Println("Channel is closed for user:", client.Username)
				return
			}

			msgJSON, err := json.Marshal(msg)

			if err != nil {
				log.Println("Error:", err)
				log.Println("Description: Cannot convert Message struct into JSON while writing")
				return
			}
			client.Connection.WriteMessage(websocket.TextMessage, msgJSON)
		case <-ticker.C:
			if err := client.Connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("No response of ping from client:", client.Username)
				return
			}
		}
	}
}
