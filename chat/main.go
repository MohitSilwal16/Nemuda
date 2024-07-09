package main

import (
	"log"

	"github.com/MohitSilwal16/Nemuda/chat/chatwebsocket"
	"github.com/gin-gonic/gin"
)

const BASE_URL = "127.0.0.1:3000"

func main() {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin engine without the default middleware
	r := gin.New()

	// Add Logger and Recovery middleware
	r.Use(gin.Logger(), gin.Recovery())

	// Serve static files from the "static" directory
	r.Static("/static", "./static")

	router := chatwebsocket.NewRouter()

	go router.Run()

	r.GET("/ws/chat", router.ServeWS)

	log.Println("Running Messaging Server on", BASE_URL)
	r.Run(BASE_URL)
}
