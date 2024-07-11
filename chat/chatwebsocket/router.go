package chatwebsocket

import (
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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin, modify as needed
	},
}

func NewRouter() *Router {
	return &Router{
		ClientsMap:  map[string]*Client{},
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		SendMessage: make(chan Message),
	}
}

func (router *Router) Run() {
	defer log.Println("Closing Router ...")
	for {
		// logger(router)
		select {
		case client := <-router.Register:
			router.Lock()

			if _, ok := router.ClientsMap[client.Username]; ok {
				client.SendMessage <- DOUBLE_CONNECTION_ERROR_BYTES
			}

			router.ClientsMap[client.Username] = client
			router.Unlock()

			log.Println("Client Registerd:", client.Username)
			log.Println("Users:", len(router.ClientsMap))

		case client := <-router.Unregister:
			router.Lock()

			if _, ok := router.ClientsMap[client.Username]; ok {
				close(client.SendMessage)
				delete(router.ClientsMap, client.Username)
			}

			router.Unlock()

			log.Println("Client Unregisterd:", client.Username)
			log.Println("Users:", len(router.ClientsMap))

		case msg := <-router.SendMessage:
			// Self message
			router.Lock()
			sender, okSender := router.ClientsMap[msg.Sender]
			router.Unlock()

			if msg.Sender == msg.Receiver && okSender {
				msg.SelfMessage = true

				msgJSON, err := json.Marshal(msg)
				if err != nil {
					log.Println("Error:", err)
					log.Println("Description: Data isn't in JSON format")
					continue
				}
				sender.SendMessage <- msgJSON
				continue
			}

			// To Receiver
			router.Lock()
			receiver, ok := router.ClientsMap[msg.Receiver]
			router.Unlock()

			var messageDelivered bool

			if ok {
				msg.SelfMessage = false
				msgJSON, err := json.Marshal(msg)

				if err != nil {
					log.Println("Error:", err)
					log.Println("Description: Data isn't in JSON format")
					continue
				}
				receiver.SendMessage <- msgJSON
				messageDelivered = true
			}

			// To Sender
			if okSender {
				msg.SelfMessage = true
				if messageDelivered {
					msg.Status = "Delivered"
				}

				msgJSON, err := json.Marshal(msg)
				if err != nil {
					log.Println("Error:", err)
					log.Println("Description: Data isn't in JSON format")
					continue
				}
				sender.SendMessage <- msgJSON
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
		return
	}
	username, err := getUsernameBySessionToken(sessionToken)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot get username from session token")

		conn.WriteMessage(websocket.TextMessage, INTERNAL_SERVER_ERROR_BYTES)
		return
	}

	if username == "" {
		log.Println("Error: Cannot get username from session token")

		conn.WriteMessage(websocket.TextMessage, SESSION_TIMED_OUT_ERROR_BYTES)
		return
	}

	client := &Client{
		Username:     username,
		Connection:   conn,
		Router:       router,
		SendMessage:  make(chan []byte),
		SessionToken: sessionToken,
	}

	router.Register <- client

	go client.readMessages()
	go client.writeMessage()
}

// Only for logging purpose
func logger(router *Router) {
	log.Println("Users: ")
	router.Lock()
	for username := range router.ClientsMap {
		log.Println(username)
	}
	router.Unlock()
}
