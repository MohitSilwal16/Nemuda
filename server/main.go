package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MohitSilwal16/Nemuda/db"
	"github.com/MohitSilwal16/Nemuda/handler"
	"github.com/MohitSilwal16/Nemuda/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	// Clear the terminal
	utils.ClearScreen()

	err := db.Init_MariaDB()

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	err = db.Init_Mongo()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func main() {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin engine without the default middleware
	r := gin.New()

	// CORS NOT WORKING OR IDK IF IT'S WORKING
	// Define the CORS middleware with specific IP addresses allowed
	// allowedOrigins := []string{
	// 	// "http://localhost:4200",
	// 	// "http://127.0.0.1:4200",
	// 	// "http://192.168.1.100",
	// }
	// config := cors.DefaultConfig()
	// config.AllowOrigins = allowedOrigins

	// r.Use(cors.New(config))

	// Add Logger and Recovery middleware
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/:sessionToken", handler.VerifySessionToken)

	r.POST("/register", handler.Register)

	r.POST("/login", handler.Login)

	r.DELETE("/login/:sessionToken", handler.Logout)

	r.GET("/users/:username", handler.SearchUser)

	r.GET("/blogs/:tag", handler.GetBlogsByTag)

	r.GET("/blogs/title/:title", handler.SearchBlogByTitle)

	r.POST("/blogs", handler.PostBlog)

	go func() {
		log.Println("Running Server on http://localhost:8080")
		r.Run("localhost:8080")
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
