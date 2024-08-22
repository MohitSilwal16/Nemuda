package handler

import (
	"context"
	"log"

	"github.com/MohitSilwal16/Nemuda/server/db"
	pb "github.com/MohitSilwal16/Nemuda/server/pb"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) DoesUserExists(ctx context.Context, req *pb.UserExistsRequest) (*pb.UserExistsResponse, error) {
	// Errors:
	// INTERNAL SERVER ERROR

	if len(req.Username) < 5 {
		return &pb.UserExistsResponse{DoesUserExists: false}, nil
	}

	usernameFound, err := db.DoesUserExists(req.Username)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot check whether username is already used or not\nSource: DoesUserExists()")

		return nil, ErrInternalServerError
	}

	if usernameFound {
		return &pb.UserExistsResponse{DoesUserExists: true}, nil
	}
	return &pb.UserExistsResponse{DoesUserExists: false}, nil
}

func (s *UserServer) SearchUsersByStartingPattern(ctx context.Context, req *pb.SearchUsersByStartingPatternRequest) (*pb.SearchUsersByStartingPatternResponse, error) {
	// Errors:
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	username, err := db.GetUsernameBySessionToken(req.SessionToken)
	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, ErrInvalidSessionToken
		}
		log.Println("Error:", err)
		log.Println("Description: Error while fetching username from session token\nSource: SearchUsersByStartingPattern()")

		return nil, ErrInternalServerError
	}

	users, err := db.SearchUsersByPattern(req.SearchPattern)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Search Users with some pattern\nSource: SearchUsersByStartingPattern()")

		return nil, ErrInternalServerError
	}

	if len(users) == 0 {
		return &pb.SearchUsersByStartingPatternResponse{UsersAndLastMessage: nil}, nil
	}

	var users_And_LastMessage []*pb.UserAndLastMessage
	var user_And_LastMessage *pb.UserAndLastMessage

	// Now fetch last message for each user
	for _, user := range users {
		message, err := db.FetchLastMessage(user, username)
		if err != nil {
			log.Println("Error:", err)
			log.Println("Description: Cannot Fetch Last Message\nSource: SearchUsersByStartingPattern()")

			return nil, ErrInternalServerError
		}

		user_And_LastMessage = &pb.UserAndLastMessage{
			Username:    user,
			LastMessage: message,
		}

		users_And_LastMessage = append(users_And_LastMessage, user_And_LastMessage)
	}

	return &pb.SearchUsersByStartingPatternResponse{UsersAndLastMessage: users_And_LastMessage}, nil
}

func (s *UserServer) GetMessagesWithPagination(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	// Errors:
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR
	// USER NOT FOUND
	// INVALID OFFSET

	if req.Offset < 0 {
		return nil, ErrInvalidOffset
	}

	user, err := db.GetUsernameBySessionToken(req.SessionToken)
	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, ErrInvalidSessionToken
		}
		log.Println("Error:", err)
		log.Println("Description: Error while fetching username from session token\nSource: GetMessagesWithPagination()")
		return nil, ErrInternalServerError
	}

	isReceiverValid, err := db.DoesUserExists(req.User1)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot check whether user exists or not\nSource: GetMessagesWithPagination()")

		return nil, ErrInternalServerError
	}

	if !isReceiverValid {
		return nil, ErrUserNotFound
	}

	messages, err := db.GetMessagesWithOffset(user, req.User1, int(req.Offset), MESSAGE_LIMIT)
	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot fetch messages\nSource: GetMessagesWithPagination()")

		return nil, ErrInternalServerError
	}

	if len(messages) == 0 {
		return &pb.GetMessagesResponse{
			Messages:   nil,
			NextOffset: -1,
		}, nil
	}

	if len(messages) < MESSAGE_LIMIT {
		return &pb.GetMessagesResponse{Messages: messages, NextOffset: -1}, nil
	}

	return &pb.GetMessagesResponse{Messages: messages, NextOffset: req.Offset + MESSAGE_LIMIT}, nil
}
