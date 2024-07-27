package handler

import (
	"context"
	"log"
	"net/http"

	pb "github.com/Nemuda/client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var userClient pb.UserServiceClient
var blogClient pb.BlogsServiceClient
var authClient pb.AuthServiceClient

func NewGRPCClients(address string) error {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	userClient = pb.NewUserServiceClient(conn)
	authClient = pb.NewAuthServiceClient(conn)
	blogClient = pb.NewBlogsServiceClient(conn)

	return nil
}

// Cookie handling
func setSessionTokenInCookie(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:  "sessionToken",
		Value: token,
		Path:  "/",
	}

	http.SetCookie(w, cookie)
}

func getSessionTokenFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("sessionToken")

	if err == http.ErrNoCookie {
		return ""
	} else if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error while Fetching Cookie")
		return ""
	}
	return cookie.Value
}

func deleteSessionTokenFromCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "sessionToken",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}

// Helper Methods
func isSessionTokenValid(sessionToken string) (bool, error) {
	req := &pb.ValidationRequest{
		SessionToken: sessionToken,
	}
	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	res, err := authClient.VerifySessionToken(ctxTimeout, req)
	if err != nil {
		return false, err
	}
	if res.IsUserVerified {
		return true, nil
	}
	return false, nil
}

func fetchBlogsByTag(sessionToken string, tag string, offset int) (*pb.GetBlogsResponse, error) {
	req := &pb.GetBlogsRequest{
		SessionToken: sessionToken,
		Tag:          tag,
		Offset:       int32(offset),
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	res, err := blogClient.GetBlogsByTagWithPagination(ctxTimeout, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func getBlogByTitle(sessionToken string, title string) (*pb.GetBlogResponse, error) {
	req := &pb.GetBlogRequest{
		SessionToken: sessionToken,
		Title:        title,
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	res, err := blogClient.GetBlogByTitle(ctxTimeout, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
