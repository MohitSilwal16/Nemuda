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

	r.GET("/", handler.DefaultRoute)
	r.GET("/nemuda", handler.RenderInitPage)

	r.GET("/register", func(ctx *gin.Context) {
		handler.RenderRegsiterPage(ctx, "")
	})
	r.POST("/register", handler.Register)

	r.GET("/login", func(ctx *gin.Context) {
		handler.RenderLoginPage(ctx, "")
	})
	r.POST("/login", handler.Login)
	r.DELETE("/login", handler.Logout)

	r.GET("/users", handler.SearchUserForRegistration)

	r.GET("/blogs/:tag", handler.GetBlogsByTag)
	r.GET("/post_blog", func(ctx *gin.Context) {
		handler.RenderPostBlogPage(ctx, "")
	})
	r.POST("/blogs", handler.PostBlog)
	r.GET("/blogs/title", handler.SearcBlogTitle_BeforePosting)
	r.GET("/blogs/title/:title", handler.GetBlogByTitle)

	r.POST("/blogs/like/:title", handler.LikeBlog)
	r.DELETE("/blogs/like/:title", handler.DislikeBlog)

	r.GET("/blogs/comment/:title", handler.AddComment)

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
