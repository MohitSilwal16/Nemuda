package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	Sender         string `json:"sender"`
	Receiver       string `json:"receiver"`
	MessageContent string `json:"messageContent"`
	Status         string `json:"status"` // Send, Delivered, Read
	DateTime       string `json:"dateTime"`
}

type Client struct {
	Username   string
	Connection *websocket.Conn
	Router     *Router
	Send       chan []byte
}

type Router struct {
	sync.RWMutex
	ClientsMap  map[*Client]bool
	Register    chan *Client
	Unregister  chan *Client
	SendMessage chan Message
}
