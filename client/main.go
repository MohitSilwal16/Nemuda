package main

import (
	"log"

	"github.com/Nemuda/client/constants"
	"github.com/Nemuda/client/controller"
	"github.com/Nemuda/client/models"
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

	r.GET("/home", controller.RenderInitPage)

	r.GET("/users", controller.SearchUserForRegistration)

	r.GET("/blogs/:tag", controller.GetMoreBlogsByTagWithOffset)

	r.GET("/post-blog", func(ctx *gin.Context) {
		controller.RenderPostBlogPage(ctx, "")
	})

	r.GET("/update-blog/:title", func(ctx *gin.Context) {
		controller.RenderUpdateBlogPage(ctx, models.Blog{}, "")
	})

	r.POST("/blogs", controller.PostBlog)
	r.PUT("/blogs/:title", controller.UpdateBlog)
	r.DELETE("/blogs/:title", controller.DeleteBlog)

	r.GET("/blogs/search_title/:method", controller.SearcBlogTitle_BeforePosting)
	r.GET("/blogs/title/:title", controller.GetBlogByTitle)

	r.POST("/blogs/like/:title", controller.LikeBlog)
	r.DELETE("/blogs/like/:title", controller.DislikeBlog)

	r.GET("/blogs/comment/:title", controller.AddComment)

	r.GET("/chats", controller.RenderChatPage)

	r.GET("/message/:user", controller.GetMessages)

	r.GET("/search-users", controller.SearchUsersByPattern)

	r.NoRoute(controller.RenderPageNotFound)
	log.Println("Running Server on", BASE_URL)
	r.Run(BASE_URL)
}
