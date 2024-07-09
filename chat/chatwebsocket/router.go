package chatwebsocket

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MohitSilwal16/Nemuda/chat/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	maxMessageSize = 512
	pongWaitTime   = time.Second * 60
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
		SendError:   make(chan ErrorMessage),
	}
}

func (router *Router) Run() {
	for {
		// Logger(router)
		select {
		case client := <-router.Register:
			router.Lock()
			if oldClient, ok := router.ClientsMap[client.Username]; ok {
				close(oldClient.SendMessage)
				oldClient.Connection.Close()
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
			sender, ok := router.ClientsMap[msg.Sender]

			if msg.Sender == msg.Receiver && ok {
				msg.SelfMessage = true

				msgJSON, err := json.Marshal(msg)
				if err != nil {
					log.Println("Error:", err)
					log.Println("Description: Data isn't in JSON format")
					return
				}
				sender.SendMessage <- msgJSON
				continue
			}

			// To Receiver
			receiver, ok := router.ClientsMap[msg.Receiver]

			var messageDelivered bool

			if ok {
				msg.SelfMessage = false
				msgJSON, err := json.Marshal(msg)

				if err != nil {
					log.Println("Error:", err)
					log.Println("Description: Data isn't in JSON format")
					return
				}
				receiver.SendMessage <- msgJSON
				messageDelivered = true
			}

			// To Sender
			sender, ok = router.ClientsMap[msg.Sender]

			if ok {
				msg.SelfMessage = true
				if messageDelivered {
					msg.Status = "Delivered"
				}

				msgJSON, err := json.Marshal(msg)
				if err != nil {
					log.Println("Error:", err)
					log.Println("Description: Data isn't in JSON format")
					return
				}
				sender.SendMessage <- msgJSON
			}

			AddMessageInDB(sender, msg)

		case errorMessage := <-router.SendError:
			log.Println("Error: ")
			log.Printf("%#v", errorMessage)

			sender, ok := router.ClientsMap[errorMessage.Username]

			if ok {
				errorMsgJSON, err := json.Marshal(errorMessage)

				if err != nil {
					log.Println("Error:", err)
					log.Println("Description: Data isn't in JSON format")
					return
				}

				sender.SendMessage <- errorMsgJSON
			}
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

	sessionToken := utils.GetSessionTokenFromCookie(ctx.Request)
	username, err := utils.GetUsernameBySessionToken(sessionToken)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot get username from session token")

		router.SendError <- ErrorMessage{}
	}

	client := &Client{
		Username:     username,
		Connection:   conn,
		Router:       router,
		SendMessage:  make(chan []byte),
		SessionToken: sessionToken,
	}

	router.Register <- client

	go client.readMessages(username)
	go client.writeMessage()
}

func Logger(router *Router) {
	log.Println("Username:")
	for username := range router.ClientsMap {
		log.Print(username)
	}
}
