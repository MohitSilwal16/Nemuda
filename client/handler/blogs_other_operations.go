package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Nemuda/client/pb"
	"github.com/Nemuda/client/utils"
	"github.com/gin-gonic/gin"
)

func GetBlogByTitle(ctx *gin.Context) {
	// Errors:
	// BLOG NOT FOUND
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	title := ctx.Param("title")
	RenderGetBlogPage(ctx, title, "")
}

func LikeBlog(ctx *gin.Context) {
	// Errors:
	// BLOG ALREADY LIKED
	// BLOG NOT FOUND
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	title := ctx.Param("title")

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	req := &pb.LikeBlogRequest{
		SessionToken: sessionToken,
		Title:        title,
	}

	_, err := blogClient.LikeBlog(ctxTimeout, req)

	if err != nil {
		msg := utils.ReturnAlertMessage(err.Error())
		fmt.Fprint(ctx.Writer, msg)
		return
	}

	RenderGetBlogPage(ctx, title, "")
}

func DislikeBlog(ctx *gin.Context) {
	// Errors:
	// BLOG ALREADY DISLIKED
	// BLOG NOT FOUND
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	title := ctx.Param("title")

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	req := &pb.DislikeBlogRequest{
		SessionToken: sessionToken,
		Title:        title,
	}

	_, err := blogClient.DislikeBlog(ctxTimeout, req)

	if err != nil {
		msg := utils.ReturnAlertMessage(err.Error())
		fmt.Fprint(ctx.Writer, msg)
		return
	}

	RenderGetBlogPage(ctx, title, "")
}

func AddComment(ctx *gin.Context) {
	// Errors:
	// BLOG NOT FOUND
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR
	// XSS DETECTED

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	title := ctx.Param("title")
	sessionToken := getSessionTokenFromCookie(ctx.Request)

	commentDescription := ctx.Query("comment")
	if len(commentDescription) < 5 || len(commentDescription) > 50 {
		RenderGetBlogPage(ctx, title, "Comment: Min 5 & Max 50 letters")
		return
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	req := &pb.AddCommentRequest{
		SessionToken:       sessionToken,
		Title:              title,
		CommentDescription: commentDescription,
	}

	_, err := blogClient.AddComment(ctxTimeout, req)
	if err != nil {
		msg := utils.ReturnAlertMessage(err.Error())
		fmt.Fprint(ctx.Writer, msg)
		return
	}
	RenderGetBlogPage(ctx, title, "")
}

func SearcBlogTitle_BeforePosting(ctx *gin.Context) {
	// Errors:
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	title := ctx.Query("title")

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), SHORT_CONTEXT_TIMEOUT)
	defer cancelFunc()

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	req := &pb.SearchBlogRequest{
		SessionToken: sessionToken,
		Title:        title,
	}

	res, err := blogClient.SearchBlogByTitle(ctxTimeout, req)
	if err != nil {
		ctx.Status(http.StatusNoContent)
		return
	}

	if res.DoesBlogExists {
		method := ctx.Param("method")
		if method == "post" {
			RenderPostBlogPage(ctx, "Title is already used")
		} else if method == "update" {
			RenderUpdateBlogPage(ctx, &pb.Blog{}, "Title is already used")
		}
	}
}
