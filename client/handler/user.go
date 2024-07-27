package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Nemuda/client/pb"
	"github.com/Nemuda/client/utils"
	"github.com/gin-gonic/gin"
)

func SearchUserForRegistration(ctx *gin.Context) {
	// Errors:
	// INTERNAL SERVER ERROR

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	username := ctx.Query("username")
	if len(username) < 5 {
		ctx.Status(http.StatusNoContent)
		return
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), SHORT_CONTEXT_TIMEOUT)
	defer cancelFunc()

	req := &pb.UserExistsRequest{
		Username: username,
	}

	res, err := userClient.DoesUserExists(ctxTimeout, req)
	if err != nil {
		ctx.Status(http.StatusNoContent)
		return
	}

	if res.DoesUserExists {
		fmt.Fprint(ctx.Writer, "Username is already used")
	}
}

func SearchUsersByPattern(ctx *gin.Context) {
	// Errors:
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	searchPattern := ctx.Query("searchPattern")

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), SHORT_CONTEXT_TIMEOUT)
	defer cancelFunc()

	sessionToken := getSessionTokenFromCookie(ctx.Request)

	req := &pb.SearchUsersByStartingPatternRequest{
		SessionToken:  sessionToken,
		SearchPattern: searchPattern,
	}

	res, err := userClient.SearchUsersByStartingPattern(ctxTimeout, req)
	if err != nil {
		msg := utils.ReturnAlertMessage(err.Error())
		fmt.Fprint(ctx.Writer, msg)
		return
	}
	RenderSearchUsersContainer(ctx, res.UsersAndLastMessage)
}

func GetMessagesWithOffset(ctx *gin.Context) {
	// Errors:
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR
	// USER NOT FOUND
	// INVALID OFFSET

	// Set the Content-Type header to "text/html"
	ctx.Header("Content-Type", "text/html")

	user1 := ctx.Param("user")
	offset := ctx.Query("offset")

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		log.Println("Error:", err)
		fmt.Fprint(ctx.Writer, "<script>alert('Offset must be non negative integer');</script>")
		return
	}
	if offsetInt < 0 {
		RenderOlderMessage(ctx, -1, nil, user1)
		return
	}
	sessionToken := getSessionTokenFromCookie(ctx.Request)

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	req := &pb.GetMessagesRequest{
		SessionToken: sessionToken,
		Offset:       int32(offsetInt),
		User1:        user1,
	}

	res, err := userClient.GetMessagesWithPagination(ctxTimeout, req)
	if err != nil {
		msg := utils.ReturnAlertMessage(err.Error())
		fmt.Fprint(ctx.Writer, msg)
		return
	}

	if offsetInt == 0 {
		if len(res.Messages) == 0 {
			RenderMessageBodyContainer(ctx, nil, user1, -2)
		} else {
			RenderMessageBodyContainer(ctx, res.Messages, user1, int(res.NextOffset))
		}
		return
	}

	if len(res.Messages) == 0 {
		RenderOlderMessage(ctx, -1, nil, user1)
		return
	}
	RenderOlderMessage(ctx, int(res.NextOffset), res.Messages, user1)
}
