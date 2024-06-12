package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Nemuda/client/controller"
	"github.com/Nemuda/client/model"
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

	r.GET("/", controller.DefaultRoute)
	r.GET("/nemuda", controller.RenderInitPage)

	r.GET("/register", func(ctx *gin.Context) {
		controller.RenderRegsiterPage(ctx, "")
	})
	r.POST("/register", controller.Register)

	r.GET("/login", func(ctx *gin.Context) {
		controller.RenderLoginPage(ctx, "")
	})
	r.POST("/login", controller.Login)
	r.DELETE("/login", controller.Logout)

	r.GET("/users", controller.SearchUserForRegistration)

	r.GET("/blogs/:tag", controller.GetMoreBlogsByTagWithOffset)

	r.GET("/post_blog", func(ctx *gin.Context) {
		controller.RenderPostBlogPage(ctx, "")
	})

	r.GET("/update_blog/:title", func(ctx *gin.Context) {
		controller.RenderUpdateBlogPage(ctx, model.Blog{}, "")
	})

	r.POST("/blogs", controller.PostBlog)
	r.PUT("/blogs/:title", controller.UpdateBlog)
	r.DELETE("/blogs/:title", controller.DeleteBlog)

	r.GET("/blogs/search_title/:method", controller.SearcBlogTitle_BeforePosting)
	r.GET("/blogs/title/:title", controller.GetBlogByTitle)

	r.POST("/blogs/like/:title", controller.LikeBlog)
	r.DELETE("/blogs/like/:title", controller.DislikeBlog)

	r.GET("/blogs/comment/:title", controller.AddComment)

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
