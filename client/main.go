package main

import (
	"log"

	"github.com/Nemuda/client/constants"
	"github.com/Nemuda/client/handler"
	"github.com/Nemuda/client/pb"
	"github.com/gin-gonic/gin"
)

const BASE_URL = constants.BASE_URL

func main() {
	err := handler.NewGRPCClients(constants.SERVICE_BASE_URL)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: ")
	}

	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin engine without the default middleware
	r := gin.New()

	// Add Logger and Recovery middleware
	r.Use(gin.Logger(), gin.Recovery())

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Serve static files from the "static" directory
	r.Static("/static", "./static")

	// Render Pages
	r.GET("/", handler.DefaultRoute)
	r.GET("/nemuda", handler.RenderInitPage)
	r.GET("/register", func(ctx *gin.Context) {
		handler.RenderRegsiterPage(ctx, "")
	})
	r.GET("/login", func(ctx *gin.Context) {
		handler.RenderLoginPage(ctx, "")
	})
	r.GET("/home/:tag", handler.RenderInitPage)
	r.GET("/post-blog", func(ctx *gin.Context) {
		handler.RenderPostBlogPage(ctx, "")
	})
	r.GET("/update-blog/:title", func(ctx *gin.Context) {
		handler.RenderUpdateBlogPage(ctx, &pb.Blog{}, "")
	})
	r.GET("/chats", handler.RenderChatPage)

	// Auth
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.DELETE("/login", handler.Logout)

	// User
	r.GET("/users", handler.SearchUserForRegistration)
	r.GET("/search-users", handler.SearchUsersByPattern)
	r.GET("/message/:user", handler.GetMessagesWithOffset)

	// // CRUD Blogs
	r.POST("/blogs", handler.PostBlog)
	r.GET("/blogs/:tag", handler.GetMoreBlogsByTagWithOffset)
	r.PUT("/blogs/:title", handler.UpdateBlog)
	r.DELETE("/blogs/:title", handler.DeleteBlog)

	// // Other Blog Operations
	r.GET("/blogs/title/:title", handler.GetBlogByTitle)
	r.POST("/blogs/like/:title", handler.LikeBlog)
	r.DELETE("/blogs/like/:title", handler.DislikeBlog)
	r.GET("/blogs/comment/:title", handler.AddComment)
	r.GET("/blogs/search_title/:method", handler.SearcBlogTitle_BeforePosting)

	r.NoRoute(handler.RenderPageNotFound)
	log.Println("Running Server on", BASE_URL)
	err = r.Run(BASE_URL)

	if err != nil {
		log.Println("Error:", err)
	}
}
