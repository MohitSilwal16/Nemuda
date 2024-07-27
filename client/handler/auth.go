package handler

import (
	"context"
	"encoding/json"
	"log"
	"regexp"

	pb "github.com/Nemuda/client/pb"
	"github.com/Nemuda/client/utils"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	// Errors:
	// USERNAME OR PASSWORD IS EMPTY
	// USERNAME MUST BE B'TWIN 5-20 CHARS & ALPHANUMERIC
	// PASSWORD MUST BE B'TWIN 8-20 CHARS, ATLEAST 1 UPPER, LOWER CASE & SPECIAL CHARACTER
	// USERNAME IS ALREADY USED
	// INTERNAL SERVER ERROR

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	var user pb.User
	err := json.NewDecoder(ctx.Request.Body).Decode(&user)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Client's data is not in JSON format\nSource: Register()")

		response := "Data isn't in JSON Format"
		RenderRegsiterPage(ctx, response)
		return
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(user.Username) {
		response := "Username should be alphanumeric b'twin 5-20 chars"
		RenderRegsiterPage(ctx, response)
		return
	} else if !utils.IsPasswordInFormat(user.Password) {
		response := "Password: 8+ chars, lower & upper case, digit, symbol"
		RenderRegsiterPage(ctx, response)
		return
	}

	req := &pb.AuthRequest{
		Username: user.Username,
		Password: user.Password,
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	res, err := authClient.Register(ctxTimeout, req)
	if err != nil {
		RenderLoginPage(ctx, utils.TrimGrpcErrorMessage(err.Error()))
		return
	}

	// Save session token in cookie of user
	setSessionTokenInCookie(ctx.Writer, res.SessionToken)
	RenderHomePage(ctx, res.SessionToken)
}

func Login(ctx *gin.Context) {
	// Errors:
	// USERNAME OR PASSWORD IS EMPTY
	// USERNAME MUST BE B'TWIN 5-20 CHARS & ALPHANUMERIC
	// INVALID CREDENTIALS
	// USER DOESN'T EXISTS
	// INTERNAL SERVER ERROR

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	var user pb.User
	err := json.NewDecoder(ctx.Request.Body).Decode(&user)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Client's data is not in JSON format\nSource: Login()")

		response := "Data isn't in JSON Format"
		RenderRegsiterPage(ctx, response)
		return
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(user.Username) {
		response := "Username should be alphanumeric b'twin 5-20 chars"
		RenderLoginPage(ctx, response)
		return
	} else if !utils.IsPasswordInFormat(user.Password) {
		// Don't want to give idea to anonymous user idea about password pattern or format
		response := "Invalid Credentials"
		RenderLoginPage(ctx, response)
		return
	}

	req := &pb.AuthRequest{
		Username: user.Username,
		Password: user.Password,
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	res, err := authClient.Login(ctxTimeout, req)
	if err != nil {
		RenderLoginPage(ctx, utils.TrimGrpcErrorMessage(err.Error()))
		return
	}

	// Save session token in cookie of user
	setSessionTokenInCookie(ctx.Writer, res.SessionToken)
	RenderHomePage(ctx, res.SessionToken)
}

func Logout(ctx *gin.Context) {
	// Errors:
	// INTERNAL SERVER ERROR
	// INVALID SESSION TOKEN

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	req := &pb.LogoutRequest{
		SessionToken: sessionToken,
	}
	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	_, err := authClient.Logout(ctxTimeout, req)
	if err != nil {
		if err.Error() == "INTERNAL SERVER ERROR" {
			RenderLoginPage(ctx, utils.TrimGrpcErrorMessage(err.Error()))
			return
		}
	}
	deleteSessionTokenFromCookie(ctx.Writer)
	RenderLoginPage(ctx, "")
}
