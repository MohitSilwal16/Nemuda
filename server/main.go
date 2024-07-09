package main

import (
	"log"
	"os"

	"github.com/MohitSilwal16/Nemuda/server/db"
	"github.com/MohitSilwal16/Nemuda/server/handler"
	"github.com/MohitSilwal16/Nemuda/server/utils"
	"github.com/gin-gonic/gin"
)

const BASE_URL = "127.0.0.1:8080"

func init() {
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

	r.GET("/get-users-by-sessionToken/:sessionToken", handler.GetUsernameBySessionToken)

	r.GET("/blogs/title/:title", handler.SearchBlogByTitle)

	r.GET("/blogs/updatable_deletable", handler.CanUserUpdate_DeleteBlog)

	r.GET("/blogs/:tag", handler.GetBlogsByTag)
	r.POST("/blogs", handler.PostBlog)
	r.PUT("/blogs", handler.UpdateBlog)
	r.DELETE("/blogs", handler.DeleteBlog)

	r.POST("/blogs/like/:title", handler.LikeBlog)
	r.GET("/blogs/like/:title", handler.IsBlogLikedByUser)
	r.DELETE("/blogs/like/:title", handler.DislikeBlog)

	r.GET("/blogs/comment/:comment", handler.AddComment)

	r.POST("/messages", handler.AddMessage)
	r.GET("/messages/:user", handler.GetMessages)

	r.GET("/search-users", handler.SearchUsersByPattern)

	log.Println("Running Back-end Server on", BASE_URL)
	r.Run(BASE_URL)
}
