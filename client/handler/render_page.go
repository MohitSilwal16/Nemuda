package handler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/Nemuda/client/pb"
	"github.com/Nemuda/client/utils"

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

	sessionToken := getSessionTokenFromCookie(ctx.Request)
	isSessionTokenValid, err := isSessionTokenValid(sessionToken)

	if err != nil {
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	if isSessionTokenValid {
		RenderHomePage(ctx, sessionToken)
		return
	}
	RenderLoginPage(ctx, "")
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

func RenderHomePage(ctx *gin.Context, sessionToken string) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")
	tag := ctx.Param("tag")

	if tag == "" {
		tag = "All"
	}

	res, err := fetchBlogsByTag(sessionToken, tag, 0)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot fetch blogs by tag\nSource: RenderHomePage()")

		message := utils.ReturnAlertMessage(err.Error())
		fmt.Fprint(ctx.Writer, message)
		return
	}

	data := map[string]interface{}{
		"Blogs":        res.Blogs,
		"RequestedTag": tag,
		"TagsList":     tagsList,
		"Offset":       res.NextOffset,
	}

	tmpl := template.Must(template.ParseFiles("./views/home.html", "./views/blog.html"))
	err = tmpl.Execute(ctx.Writer, data)

	if err != nil {
		log.Println("Error:", err)
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
	sessionToken := getSessionTokenFromCookie(ctx.Request)
	isSessionTokenValid, err := isSessionTokenValid(sessionToken)

	if err != nil {
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
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

func RenderUpdateBlogPage(ctx *gin.Context, blog *pb.Blog, message string) {
	temp, err := template.ParseFiles("./views/update_blog.html")
	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RenderUpdateBlogPage()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	title := ctx.Param("title")

	if blog.Title == "" {
		sessionToken := getSessionTokenFromCookie(ctx.Request)
		res, err := getBlogByTitle(sessionToken, title)
		blog = res.Blog
		if err != nil {
			return
		}
	}

	sessionToken := getSessionTokenFromCookie(ctx.Request)
	isSessionTokenValid, err := isSessionTokenValid(sessionToken)

	if err != nil {
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
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

func RenderGetBlogPage(ctx *gin.Context, title string, message string) {
	temp, err := template.ParseFiles("./views/view_blog.html")
	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RenderGetBlogPage()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	sessionToken := getSessionTokenFromCookie(ctx.Request)
	res, err := getBlogByTitle(sessionToken, title)
	if err != nil {
		msg := utils.ReturnAlertMessage(err.Error())
		fmt.Fprint(ctx.Writer, msg)
	}

	data := map[string]interface{}{
		"Blog":                res.Blog,
		"IsBlogLiked":         res.IsBlogLiked,
		"Message":             message,
		"IsEditableDeletable": res.IsBlogUpdatableDeletable,
	}

	temp.Execute(ctx.Writer, data)
}

func RenderPageNotFound(ctx *gin.Context) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	tmpl := template.Must(template.ParseFiles("./views/page_not_found.html"))
	err := tmpl.Execute(ctx.Writer, nil)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RengerRegisterPage()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}
}

func RenderChatPage(ctx *gin.Context) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	sessionToken := getSessionTokenFromCookie(ctx.Request)
	isSessionTokenValid, err := isSessionTokenValid(sessionToken)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Validate Session Token\nSource: RenderChatPage()")
		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	if !isSessionTokenValid {
		RenderLoginPage(ctx, "Session Timed Out")
		return
	}

	tmpl, err := template.ParseFiles("./views/chats.html", "./views/messaging_page.html", "./views/search_users.html")

	data := map[string]interface{}{
		"Messages": nil,
		"Receiver": "Nemu Chat",
		"Users":    nil,
		"Offset":   "0",
	}

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RenderChatPage()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}

	tmpl.Execute(ctx.Writer, data)
}

func RenderMessageBodyContainer(ctx *gin.Context, messages []*pb.Message, user1 string, nextOffset int) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("./views/messaging_page.html")

	data := map[string]interface{}{
		"Messages": messages,
		"Receiver": user1,
		"Offset":   nextOffset,
	}

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in RenderChatPage()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}

	tmpl.Execute(ctx.Writer, data)
}

func RenderSearchUsersContainer(ctx *gin.Context, usersAndLastMessage []*pb.UserAndLastMessage) {
	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("./views/search_users.html")

	data := map[string][]*pb.UserAndLastMessage{
		"Users": usersAndLastMessage,
	}

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Error in tmpl.Execute() in RenderChatPage()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}

	tmpl.Execute(ctx.Writer, data)
}

func RenderOlderMessage(ctx *gin.Context, offset int, messages []*pb.Message, receiver string) {
	tmpl, err := template.ParseFiles("./views/message.html")
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Error in tmpl.Execute() in GetMessagesWithOffset() for message.html")

		return
	}

	data := map[string]interface{}{
		"Messages": messages,
		"Offset":   offset,
		"Receiver": receiver,
	}

	tmpl.Execute(ctx.Writer, data)
}
