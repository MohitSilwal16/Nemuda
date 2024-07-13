package chatwebsocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WSMessage struct {
	Message      string `json:"message"`
	Receiver     string `json:"receiver"`
	SessionToken string `json:"sessionToken"`
	MessageType  string `json:"messageType"` // Message, Read
}

type Client struct {
	Username     string
	Connection   *websocket.Conn
	Router       *Router
	SendMessage  chan Message
	SessionToken string
}

type Router struct {
	sync.RWMutex
	ClientsMap         map[string]*Client // Key: Username
	Register           chan *Client
	Unregister         chan *Client
	SendMessage        chan Message
	UsersWhoAreWaiting map[string][]string // Stores users who'll be notified when this user gets online
}

type Message struct {
	Sender         string `json:"sender"`
	Receiver       string `json:"receiver"`
	MessageContent string `json:"messageContent"`
	Status         string `json:"status"` // Sent, Delivered, Read
	DateTime       string `json:"dateTime"`
	SelfMessage    bool   `json:"selfMessage"` // Sending message to himself/herself
	Error          string `json:"error"`
	MessageType    string `json:"messageType"` // Message, Read, Delivered
}

// If msg is of type Read or Delivered then,
// Sender means the one who is reading the message or to whom the message got delivered
// Receiver means the guy who sent message
