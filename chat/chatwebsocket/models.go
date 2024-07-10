package chatwebsocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WSMessage struct {
	Message  string `json:"message"`
	Receiver string `json:"receiver"`
}

type Client struct {
	Username     string
	Connection   *websocket.Conn
	Router       *Router
	SendMessage  chan []byte
	SessionToken string
}

type Router struct {
	sync.RWMutex
	ClientsMap  map[string]*Client // Key: Username
	Register    chan *Client
	Unregister  chan *Client
	SendMessage chan Message
	SendError   chan ErrorMessage
}

type Message struct {
	Sender         string `json:"sender"`
	Receiver       string `json:"receiver"`
	MessageContent string `json:"messageContent"`
	Status         string `json:"status"` // Sent, Delivered, Read
	DateTime       string `json:"dateTime"`
	SelfMessage    bool   `json:"selfMessage"` // Sending message to himself/herself
}

type ErrorMessage struct {
	Username string `json:"username"`
	Error    string `json:"error"`
}
