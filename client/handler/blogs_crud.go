package handler

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"text/template"

	"github.com/Nemuda/client/pb"
	"github.com/Nemuda/client/utils"
	"github.com/gin-gonic/gin"
)

func PostBlog(ctx *gin.Context) {
	// Errors:
	// TITLE MUST BE B'TWIN 5-20 CHARS & ALPHANUMERIC
	// BLOG TITLE IS ALREADY USED
	// DESCRIPTION MUST BE B'TWIN 4-50 CHARS
	// INVALID BLOG TAG
	// INVALID FILE TYPE, ONLY JPG, JPEG & PNG ARE ACCEPTED
	// IMAGE SIZE EXCEEDS 2 MB
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR
	// XSS DETECTED

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	sessionToken := getSessionTokenFromCookie(ctx.Request)
	if sessionToken == "" {
		RenderLoginPage(ctx, "Session Timed Out")
		return
	}

	var blog pb.Blog
	blog.Title = ctx.Request.PostFormValue("title")
	blog.Tag = ctx.Request.PostFormValue("tag")
	blog.Description = ctx.Request.PostFormValue("description")

	if !regexp.MustCompile(`^[a-zA-Z0-9 ,'"&]*$`).MatchString(blog.Title) {
		response := "Title should be alphanumeric b'twin 5-20 chars"
		RenderUpdateBlogPage(ctx, &blog, response)
		return
	}
	if !utils.Contains(tagsList, blog.Tag) {
		response := "Unknown tag"
		RenderUpdateBlogPage(ctx, &blog, response)
		return
	}
	if len(blog.Description) < 4 || len(blog.Description) > 50 {
		response := "Min 4 letters & Max 50 letters"
		RenderUpdateBlogPage(ctx, &blog, response)
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot fetch image from user")

		RenderUpdateBlogPage(ctx, &blog, "Failed to fetch image")
		return
	}

	if image.Header.Get("Content-Type") != "image/jpeg" &&
		image.Header.Get("Content-Type") != "image/png" &&
		image.Header.Get("Content-Type") != "image/jpg" {
		RenderUpdateBlogPage(ctx, &blog, "Invalid file type, upload an image")
		return
	}

	maxSize := 2 * 1024 * 1024 // 2MB
	if image.Size > int64(maxSize) {
		RenderUpdateBlogPage(ctx, &blog, "Image Size Exceeds 2 MB")
		return
	}

	// Read image file into a byte slice
	imageData, err := utils.FileHeaderToBytes(image)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Convert Image into Bytes\nSource: UpdateBlog()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	req := &pb.PostBlogRequest{
		SessionToken: sessionToken,
		Title:        blog.Title,
		Tag:          blog.Tag,
		Description:  blog.Description,
		ImageData:    imageData,
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), LONG_CONTEXT_TIMEOUT)
	defer cancelFunc()

	_, err = blogClient.PostBlog(ctxTimeout, req)
	if err != nil {
		msg := utils.TrimGrpcErrorMessage(err.Error())
		RenderPostBlogPage(ctx, msg)
		return
	}
	RenderHomePage(ctx, sessionToken)
}

func GetMoreBlogsByTagWithOffset(ctx *gin.Context) {
	// Errors:
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR
	// INVALID BLOG TAG
	// INVALID OFFSET

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	tag := ctx.Param("tag")

	sessionToken := getSessionTokenFromCookie(ctx.Request)
	if sessionToken == "" {
		RenderLoginPage(ctx, "Session Timed Out")
		return
	}

	offset := ctx.Query("offset")
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		fmt.Fprint(ctx.Writer, "<script>alert('Offset must be non negative integer');</script>")
		return
	}
	if offset == "0" {
		RenderHomePage(ctx, sessionToken)
		return
	}

	if offsetInt < 0 {
		html := `
			<!-- No More Blogs Container -->
			<div class="flex items-center justify-center cursor-not-allowed">
			<div class="px-6 py-4 text-center bg-blue-600 rounded-lg shadow-md">
				<p class="text-lg font-semibold text-gray-100">No more Blogs</p>
			</div>
			</div>
		`
		fmt.Fprint(ctx.Writer, html)
		return
	}

	res, err := fetchBlogsByTag(sessionToken, tag, offsetInt)
	if err != nil {
		message := utils.ReturnAlertMessage(err.Error())
		fmt.Fprint(ctx.Writer, message)
		return
	}

	if len(res.Blogs) == 0 {
		html := `
		<!-- No More Blogs Container -->
		<div class="flex items-center justify-center cursor-not-allowed">
		<div class="px-6 py-4 text-center bg-blue-600 rounded-lg shadow-md">
			<p class="text-lg font-semibold text-gray-100">No more Blogs</p>
		</div>
		</div>
	`
		fmt.Fprint(ctx.Writer, html)
		return
	}

	data := map[string]interface{}{
		"Blogs":        res.Blogs,
		"RequestedTag": tag,
		"Offset":       res.NextOffset,
	}

	tmpl := template.Must(template.ParseFiles("./views/more_blogs.html"))
	err = tmpl.Execute(ctx.Writer, data)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error in tmpl.Execute() in GetMoreBlogsByTagWithOffset()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
	}
}

func UpdateBlog(ctx *gin.Context) {
	// Errors:
	// TITLE MUST BE B'TWIN 5-20 CHARS & ALPHANUMERIC
	// BLOG TITLE IS ALREADY USED
	// DESCRIPTION MUST BE B'TWIN 4-50 CHARS
	// INVALID BLOG TAG
	// INVALID FILE TYPE, ONLY JPG, JPEG & PNG ARE ACCEPTED
	// IMAGE SIZE EXCEEDS 2 MB
	// BLOG NOT FOUND // Old Blog not found
	// USER CANNOT UPDATE THIS BLOG
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR
	// XSS DETECTED

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	sessionToken := getSessionTokenFromCookie(ctx.Request)
	if sessionToken == "" {
		RenderLoginPage(ctx, "Session Timed Out")
		return
	}

	oldTitle := ctx.Param("title")

	var blog pb.Blog
	blog.Title = ctx.Request.PostFormValue("title")
	blog.Tag = ctx.Request.PostFormValue("tag")
	blog.Description = ctx.Request.PostFormValue("description")

	if !regexp.MustCompile(`^[a-zA-Z0-9 ,'"&]*$`).MatchString(blog.Title) {
		response := "Title should be alphanumeric b'twin 5-20 chars"
		RenderUpdateBlogPage(ctx, &blog, response)
		return
	}
	if !utils.Contains(tagsList, blog.Tag) {
		response := "Unknown tag"
		RenderUpdateBlogPage(ctx, &blog, response)
		return
	}
	if len(blog.Description) < 4 || len(blog.Description) > 50 {
		response := "Min 4 letters & Max 50 letters"
		RenderUpdateBlogPage(ctx, &blog, response)
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot fetch image from user")

		RenderUpdateBlogPage(ctx, &blog, "Failed to fetch image")
		return
	}

	if image.Header.Get("Content-Type") != "image/jpeg" &&
		image.Header.Get("Content-Type") != "image/png" &&
		image.Header.Get("Content-Type") != "image/jpg" {
		RenderUpdateBlogPage(ctx, &blog, "Invalid file type, upload an image")
		return
	}

	maxSize := 2 * 1024 * 1024 // 2MB
	if image.Size > int64(maxSize) {
		RenderUpdateBlogPage(ctx, &blog, "Image Size Exceeds 2 MB")
		return
	}

	// Read image file into a byte slice
	imageData, err := utils.FileHeaderToBytes(image)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Convert Image into Bytes\nSource: UpdateBlog()")

		fmt.Fprint(ctx.Writer, INTERNAL_SERVER_ERROR_MESSAGE)
		return
	}

	req := &pb.UpdateBlogRequest{
		SessionToken:   sessionToken,
		OldTitle:       oldTitle,
		NewTitle:       blog.Title,
		NewTag:         blog.Tag,
		NewDescription: blog.Description,
		NewImageData:   imageData,
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), LONG_CONTEXT_TIMEOUT)
	defer cancelFunc()

	_, err = blogClient.UpdateBlog(ctxTimeout, req)
	if err != nil {
		msg := utils.TrimGrpcErrorMessage(err.Error())
		RenderUpdateBlogPage(ctx, &blog, msg)
		return
	}
	RenderHomePage(ctx, sessionToken)
}

func DeleteBlog(ctx *gin.Context) {
	// Errors:
	// USER CANNOT DELETE THIS BLOG
	// BLOG NOT FOUND
	// INTERNAL SERVER ERROR
	// INVALID SESSION TOKEN

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	title := ctx.Param("title")
	sessionToken := getSessionTokenFromCookie(ctx.Request)

	ctxTimeout, cancelFunction := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunction()

	req := &pb.DeleteBlogRequest{
		SessionToken: sessionToken,
		Title:        title,
	}

	_, err := blogClient.DeleteBlog(ctxTimeout, req)

	if err != nil {
		msg := utils.ReturnAlertMessage(err.Error())
		fmt.Fprint(ctx.Writer, msg)
		return
	}
	RenderHomePage(ctx, sessionToken)
}
