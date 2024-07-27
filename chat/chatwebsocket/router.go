package chatwebsocket

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	maxMessageSize = 512
	pongWaitTime   = time.Second * 30
	pingWaitTime   = (pongWaitTime * 9) / 10
)

var INTERNAL_SERVER_ERROR_MESSAGE = Message{
	Error: "Internal Server Error",
}

var REQUEST_TIMED_OUT_MESSAGE = Message{
	Error: "Request Timed Out",
}

var SESSION_TIMED_OUT_MESSAGE = Message{
	Error: "Session Timed Out",
}

var REQUEST_TIMED_OUT_MESSAGE_BYTES, _ = json.Marshal(REQUEST_TIMED_OUT_MESSAGE)
var SESSION_TIMED_OUT_ERROR_BYTES, _ = json.Marshal(SESSION_TIMED_OUT_MESSAGE)
var INTERNAL_SERVER_ERROR_BYTES, _ = json.Marshal(SESSION_TIMED_OUT_MESSAGE)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin, modify as needed
	},
}

func NewRouter() *Router {
	return &Router{
		ClientsMap:         map[string]*Client{},
		Register:           make(chan *Client),
		Unregister:         make(chan *Client),
		SendMessage:        make(chan Message),
		UsersWhoAreWaiting: map[string][]string{},
	}
}

func (router *Router) Run() {
	defer log.Println("Closing Router ...")
	for {
		// Logger(router)
		select {
		case client, ok := <-router.Register:
			if !ok {
				log.Println("Register Channel is closed")
			}
			router.Lock()

			// If user is connecting same account from two devices, tabs or windows then disconnect both accounts
			if oldClient, ok := router.ClientsMap[client.Username]; ok {
				close(oldClient.SendMessage)
				if oldClient.Connection != nil {
					oldClient.Connection.Close()
				}
				delete(router.ClientsMap, oldClient.Username)
				close(client.SendMessage)
				if client.Connection != nil {
					client.Connection.Close()
				}
				router.Unlock()
				log.Println("Double dipping,", oldClient.Username, "removed")
				continue
			}

			router.ClientsMap[client.Username] = client

			usersWhoAreWaiting, ok := router.UsersWhoAreWaiting[client.Username]

			if !ok {
				router.Unlock()

				log.Println("Client Registerd:", client.Username)
				log.Println("Users:", len(router.ClientsMap))
				continue
			}

			if len(usersWhoAreWaiting) == 0 {
				router.Unlock()

				log.Println("Client Registerd:", client.Username)
				log.Println("Users:", len(router.ClientsMap))
				continue
			}

			for _, cli := range usersWhoAreWaiting {
				if waitingClient, ok := router.ClientsMap[cli]; ok {
					waitingClient.SendMessage <- Message{
						Sender:      client.Username,
						Receiver:    cli,
						MessageType: "Delivered",
					}
				}
			}
			// Change status of messages which're sent to this user
			changeStatusOfMessage(client, client.Username, "", "Delivered")

			router.UsersWhoAreWaiting[client.Username] = []string{}

			router.Unlock()

			log.Println("Client Registerd:", client.Username)
			log.Println("Users:", len(router.ClientsMap))

		case client, ok := <-router.Unregister:
			if !ok {
				log.Println("Unregister Channel is closed")
			}

			router.Lock()

			if _, ok := router.ClientsMap[client.Username]; ok {
				close(client.SendMessage)
				if client.Connection != nil {
					client.Connection.Close()
				}
				delete(router.ClientsMap, client.Username)
			}

			router.Unlock()

			log.Println("Client Unregisterd:", client.Username)
			log.Println("Users:", len(router.ClientsMap))

		case msg, ok := <-router.SendMessage:
			if !ok {
				log.Println("Send Message Channel of Router is closed")
			}

			// Giving feedback that I've read your message
			if msg.MessageType == "Read" {
				// To Receiver
				router.Lock()
				receiver, ok := router.ClientsMap[msg.Receiver]
				router.Unlock()

				if ok {
					receiver.SendMessage <- msg
				}

				router.Lock()
				sender, okSender := router.ClientsMap[msg.Sender]
				router.Unlock()

				if okSender {
					changeStatusOfMessage(sender, sender.Username, msg.Receiver, "Read")
				}

				continue
			}

			// Self message
			router.Lock()
			sender, okSender := router.ClientsMap[msg.Sender]
			router.Unlock()

			if msg.Sender == msg.Receiver && okSender {
				msg.Status = "Read"
				msg.SelfMessage = true
				sender.SendMessage <- msg

				addMessageInDB(sender, msg)
				continue
			}

			// To Receiver
			router.Lock()
			receiver, ok := router.ClientsMap[msg.Receiver]
			router.Unlock()

			var messageDelivered bool

			if ok {
				msg.SelfMessage = false
				receiver.SendMessage <- msg
				messageDelivered = true
			}

			// To Sender
			if okSender {
				msg.SelfMessage = true
				if messageDelivered {
					msg.Status = "Delivered"
				} else {
					router.Lock()
					router.UsersWhoAreWaiting[msg.Receiver] = append(router.UsersWhoAreWaiting[msg.Receiver], sender.Username)
					router.Unlock()
				}
				sender.SendMessage <- msg
			}

			addMessageInDB(sender, msg)
		}
	}
}

func (router *Router) ServeWS(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot upgrade HTTP to Websocket")
		return
	}

	sessionToken := ctx.Param("sessionToken")
	if sessionToken == "" {
		log.Println("Error: Empty Session Token")
		conn.WriteMessage(websocket.TextMessage, SESSION_TIMED_OUT_ERROR_BYTES)
		conn.WriteMessage(websocket.CloseMessage, []byte{})
		return
	}

	username, err := getUsernameBySessionToken(sessionToken)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot get username from session token")
		if err == context.DeadlineExceeded {
			conn.WriteMessage(websocket.TextMessage, REQUEST_TIMED_OUT_MESSAGE_BYTES)
		} else {
			conn.WriteMessage(websocket.TextMessage, INTERNAL_SERVER_ERROR_BYTES)
		}
		conn.WriteMessage(websocket.CloseMessage, []byte{})
		return
	}

	if username == "" {
		log.Println("Error: Cannot get username from session token")
		conn.WriteMessage(websocket.TextMessage, SESSION_TIMED_OUT_ERROR_BYTES)
		conn.WriteMessage(websocket.CloseMessage, []byte{})
		return
	}

	client := &Client{
		Username:     username,
		Connection:   conn,
		Router:       router,
		SendMessage:  make(chan Message),
		SessionToken: sessionToken,
	}

	router.Register <- client

	go client.readMessages()
	go client.writeMessage()
}

// Only for logging purpose
func Logger(router *Router) {
	log.Println("Users: ")
	router.Lock()
	for username := range router.ClientsMap {
		log.Println(username)
	}
	router.Unlock()
}
