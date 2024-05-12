package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Nemuda/client/handler"
	"github.com/Nemuda/client/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// Clear the terminal
	utils.ClearScreen()

	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin engine without the default middleware
	r := gin.New()

	// Add Logger and Recovery middleware
	r.Use(gin.Logger(), gin.Recovery())

	// Serve static files from the "static" directory
	r.Static("/static", "./static")

	r.GET("/", func(ctx *gin.Context) {
		handler.DefaultRoute(ctx)
	})
	r.GET("/nemuda", func(ctx *gin.Context) {
		handler.RenderInitPage(ctx)
	})

	r.GET("/register", func(ctx *gin.Context) {
		handler.RenderRegsiterPage(ctx)
	})

	r.POST("/register", func(ctx *gin.Context) {
		handler.Register(ctx)
	})

	r.GET("/login", func(ctx *gin.Context) {
		handler.RenderLoginPage(ctx)
	})

	r.POST("/login", func(ctx *gin.Context) {
		handler.Login(ctx)
	})

	r.DELETE("/login/", func(ctx *gin.Context) {
		handler.Logout(ctx)
	})

	r.GET("/home", func(ctx *gin.Context) {
		handler.RenderHomePage(ctx)
	})

	go func() {
		log.Println("Running Server on http://localhost:4200")
		r.Run("localhost:4200")
	}()

	var choi string

	for {
		fmt.Scanln(&choi)

		switch choi {
		case "h":
			log.Println("h - help")
			log.Println("c - clear")
			log.Println("q - quit")
		case "c":
			utils.ClearScreen()
		case "q":
			os.Exit(0)
		default:
			log.Println("Enter h for help")
		}
	}
}
