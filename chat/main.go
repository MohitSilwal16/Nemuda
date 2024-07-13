package main

import (
	"log"

	"github.com/MohitSilwal16/Nemuda/chat/chatwebsocket"
	"github.com/MohitSilwal16/Nemuda/chat/constants"
	"github.com/gin-gonic/gin"
)

const BASE_URL = constants.BASE_URL

func main() {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin engine without the default middleware
	r := gin.New()

	// Add Logger and Recovery middleware
	r.Use(gin.Logger(), gin.Recovery())

	router := chatwebsocket.NewRouter()

	go router.Run()

	r.GET("/ws/chat/:sessionToken", router.ServeWS)

	log.Println("Running Messaging Server on", BASE_URL)
	
	err := r.Run(BASE_URL)

	if err != nil {
		log.Println("Error:", err)
	}
}
