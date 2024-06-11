package controller

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/Nemuda/client/model"
	"github.com/gin-gonic/gin"
)

// Handlers
func DefaultRoute(ctx *gin.Context) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	http.ServeFile(ctx.Writer, ctx.Request, "./views/index.html")
}

func RenderInitPage(ctx *gin.Context) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	isSessionTokenValid, err := isSessionTokenValid(ctx)

	if err != nil {
		fmt.Fprint(ctx.Writer, "")
		return
	}

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	if isSessionTokenValid {
		fetchBlogsByTag(ctx, "All", sessionToken)
	} else {
		RenderLoginPage(ctx, "Session Timed Out")
	}
}

func RenderRegsiterPage(ctx *gin.Context, message string) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	tmpl := template.Must(template.ParseFiles("./views/register.html"))
	err := tmpl.Execute(ctx.Writer, message)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RengerRegisterPage()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}
}

func RenderLoginPage(ctx *gin.Context, message string) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	tmpl := template.Must(template.ParseFiles("./views/login.html"))
	err := tmpl.Execute(ctx.Writer, message)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RenderLoginPage()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}
}

func RenderHomePage(ctx *gin.Context, blogs []model.Blog, tag string) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	data := map[string]interface{}{
		"Blogs":        blogs,
		"RequestedTag": tag,
		"TagsList":     tagsList,
	}

	tmpl := template.Must(template.ParseFiles("./views/home.html", "./views/blog.html"))
	err := tmpl.Execute(ctx.Writer, data)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RenderHomePage()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}
}

func RenderPostBlogPage(ctx *gin.Context, message string) {
	temp, err := template.ParseFiles("./views/post_blog.html")
	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RenderPostBlogPage()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}
	isSessionTokenValid, err := isSessionTokenValid(ctx)

	if err != nil {
		fmt.Fprint(ctx.Writer, "")
		return
	}

	if !isSessionTokenValid {
		RenderLoginPage(ctx, "Session Timed Out")
		return
	}

	data := map[string]interface{}{
		"TagsList": tagsList,
		"Message":  message,
	}
	temp.Execute(ctx.Writer, data)
}

func RenderUpdateBlogPage(ctx *gin.Context, blog model.Blog, message string) {
	temp, err := template.ParseFiles("./views/update_blog.html")
	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RenderUpdateBlogPage()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	title := ctx.Param("title")

	if blog.Title == "" {
		blog, err = fetchBlogByTitleAndReturn(ctx, title)
		if err != nil {
			return
		}
	}

	isSessionTokenValid, err := isSessionTokenValid(ctx)

	if err != nil {
		fmt.Fprint(ctx.Writer, "")
		return
	}

	if !isSessionTokenValid {
		RenderLoginPage(ctx, "Session Timed Out")
		return
	}

	data := map[string]interface{}{
		"TagsList": tagsList,
		"Message":  message,
		"Blog":     blog,
		"OldTitle": title,
	}
	temp.Execute(ctx.Writer, data)
}

func RenderGetBlogPage(ctx *gin.Context, blog model.Blog, message string) {
	temp, err := template.ParseFiles("./views/view_blog.html")
	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RenderGetBlogPage()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	isBlogLiked, err := isBlogLiked(ctx, blog.Title, sessionToken)

	if err != nil {
		fmt.Fprint(ctx.Writer, "")
		return
	}

	isEditableDeletable, err := isBlogEditableDeletable(ctx, blog.Title, sessionToken)

	if err != nil {
		fmt.Fprint(ctx.Writer, "")
		return
	}

	data := map[string]interface{}{
		"Blog":                blog,
		"IsBlogLiked":         isBlogLiked,
		"Message":             message,
		"IsEditableDeletable": isEditableDeletable,
	}

	temp.Execute(ctx.Writer, data)
}
