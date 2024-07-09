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
		ClientsMap:  map[*Client]bool{},
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		SendMessage: make(chan Message),
	}
}

func (router *Router) Run() {
	for {
		// Logger(router)
		select {
		case client := <-router.Register:
			router.Lock()
			router.ClientsMap[client] = true
			router.Unlock()

			log.Println("Client Registerd:", client.Username)
			log.Println("Users:", len(router.ClientsMap))

		case client := <-router.Unregister:
			router.Lock()
			if _, ok := router.ClientsMap[client]; ok {
				close(client.Send)
				delete(router.ClientsMap, client)
			}
			router.Unlock()
			log.Println("Client Unregisterd:", client.Username)
			log.Println("Users:", len(router.ClientsMap))

		case msg := <-router.SendMessage:

			for client := range router.ClientsMap {
				if client.Username == msg.Receiver && client.Username == msg.Sender {
					// AddMessageInDB(client, msg)

					msg.SelfMessage = true

					msgJSON, err := json.Marshal(msg)

					if err != nil {
						log.Println("Error:", err)
						log.Println("Description: Data isn't in JSON format")
						return
					}

					select {
					case client.Send <- msgJSON:
					default:
						close(client.Send)
						delete(router.ClientsMap, client)
					}
				} else if client.Username == msg.Receiver {
					msg.SelfMessage = false

					msgJSON, err := json.Marshal(msg)

					if err != nil {
						log.Println("Error:", err)
						log.Println("Description: Data isn't in JSON format")
						return
					}

					select {
					case client.Send <- msgJSON:
					default:
						close(client.Send)
						delete(router.ClientsMap, client)
					}
				} else if client.Username == msg.Sender {
					// AddMessageInDB(client, msg)

					msg.SelfMessage = true

					msgJSON, err := json.Marshal(msg)

					if err != nil {
						log.Println("Error:", err)
						log.Println("Description: Data isn't in JSON format")
						return
					}

					select {
					case client.Send <- msgJSON:
					default:
						close(client.Send)
						delete(router.ClientsMap, client)
					}
				}
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

	username, err := utils.GetUsernameBySessionToken(ctx)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot get username from session token")
		return
	}

	sessionToken := utils.GetSessionTokenFromCookie(ctx.Request)

	client := &Client{
		Username:     username,
		Connection:   conn,
		Router:       router,
		Send:         make(chan []byte),
		SessionToken: sessionToken,
	}

	router.Register <- client

	go client.readMessages()
	go client.writeMessage()
}

func Logger(router *Router) {
	log.Println("Username:")
	for client := range router.ClientsMap {
		log.Print(client.Username)
	}
}
