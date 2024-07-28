package handler

import (
	"context"
	"log"

	"github.com/MohitSilwal16/Nemuda/server/db"
	pb "github.com/MohitSilwal16/Nemuda/server/pb"
	"github.com/MohitSilwal16/Nemuda/server/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *BlogsServer) GetBlogByTitle(ctx context.Context, req *pb.GetBlogRequest) (*pb.GetBlogResponse, error) {
	// Errors:
	// BLOG NOT FOUND
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	username, err := db.GetUsernameBySessionToken(req.SessionToken)
	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, ErrInvalidSessionToken
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot fetch username from session token\nSource: LikeBlog()")

		return nil, ErrInternalServerError
	}

	blog, err := db.GetBlogByTitle(req.Title)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return nil, ErrBlogNotFound
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Get Blog by Title\nSource: SearchBlogByTitle()")
		return nil, ErrInternalServerError
	}

	var isBlogLiked bool
	var isBlogUpdatableDeletable bool

	for _, val := range blog.LikedUsername {
		if val == username {
			isBlogLiked = true
			break
		}
	}

	if blog.Username == username {
		isBlogUpdatableDeletable = true
	}

	return &pb.GetBlogResponse{
		Blog:                     blog,
		IsBlogLiked:              isBlogLiked,
		IsBlogUpdatableDeletable: isBlogUpdatableDeletable,
	}, nil
}

func (s *BlogsServer) LikeBlog(ctx context.Context, req *pb.LikeBlogRequest) (*pb.LikeBlogResponse, error) {
	// Errors:
	// BLOG ALREADY LIKED
	// BLOG NOT FOUND
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	username, err := db.GetUsernameBySessionToken(req.SessionToken)
	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, ErrInvalidSessionToken
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot fetch username from session token\nSource: LikeBlog()")

		return nil, ErrInternalServerError
	}

	err = db.LikeBlog(req.Title, username)
	if err != nil {
		if err.Error() == "BLOG NOT FOUND" {
			return nil, ErrBlogNotFound
		}
		if err.Error() == "BLOG ALREADY LIKED" {
			return nil, err
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot like blog\nSource: LikeBlog()")

		return nil, ErrInternalServerError
	}
	return &pb.LikeBlogResponse{IsBlogLiked: true}, nil
}

func (s *BlogsServer) DislikeBlog(ctx context.Context, req *pb.DislikeBlogRequest) (*pb.DislikeBlogResponse, error) {
	// Errors:
	// BLOG ALREADY DISLIKED
	// BLOG NOT FOUND
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	username, err := db.GetUsernameBySessionToken(req.SessionToken)
	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, ErrInvalidSessionToken
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot fetch username from session token\nSource: DislikeBlog()")

		return nil, ErrInternalServerError
	}

	err = db.DislikeBlog(req.Title, username)
	if err != nil {
		if err.Error() == "BLOG NOT FOUND" {
			return nil, ErrBlogNotFound
		}
		if err.Error() == "BLOG ALREADY DISLIKED" {
			return nil, err
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot dislike blog\nSource: DislikeBlog()")

		return nil, ErrInternalServerError
	}
	return &pb.DislikeBlogResponse{IsBlogDisliked: true}, nil
}

func (s *BlogsServer) AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.AddCommentResponse, error) {
	// Errors:
	// BLOG NOT FOUND
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR
	// XSS DETECTED

	username, err := db.GetUsernameBySessionToken(req.SessionToken)
	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, ErrInvalidSessionToken
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot fetch username from session token\nSource: AddComment()")

		return nil, ErrInternalServerError
	}

	if len(req.CommentDescription) < 5 || len(req.CommentDescription) > 50 {
		return nil, ErrInvalidBlogCommentFormat
	}
	isMalicious := utils.IsMessageMalicious(req.CommentDescription)

	if isMalicious {
		return nil, ErrXSSDetected
	}

	comment := &pb.Comment{
		Username:    username,
		Description: req.CommentDescription,
	}

	err = db.AddComment(req.Title, comment)
	if err != nil {
		if err.Error() == "BLOG NOT FOUND" {
			return nil, ErrBlogNotFound
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot add comment\nSource: AddComment()")

		return nil, ErrInternalServerError
	}
	return &pb.AddCommentResponse{IsCommentAdded: true}, nil
}

func (s *BlogsServer) SearchBlogByTitle(ctx context.Context, req *pb.SearchBlogRequest) (*pb.SearchBlogResponse, error) {
	// Errors:
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	isSessionTokenValid, err := db.IsSessionTokenValid(req.SessionToken)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Validate Session Token\nSource: SearchBlogByTitle()")

		return nil, ErrInternalServerError
	}
	if !isSessionTokenValid {
		return nil, ErrInvalidSessionToken
	}

	doesBlogExists, err := db.SearchBlogByTitle(req.Title)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return nil, ErrBlogNotFound
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Get Blog by Title\nSource: SearchBlogByTitle()")
		return nil, ErrInternalServerError
	}

	return &pb.SearchBlogResponse{
		DoesBlogExists: doesBlogExists,
	}, nil
}
