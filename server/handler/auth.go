package handler

import (
	"context"
	"log"
	"regexp"

	"github.com/MohitSilwal16/Nemuda/server/db"
	pb "github.com/MohitSilwal16/Nemuda/server/pb"
	"github.com/MohitSilwal16/Nemuda/server/utils"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Login(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	// Errors:
	// USERNAME MUST BE B'TWIN 5-20 CHARS & ALPHANUMERIC
	// INVALID CREDENTIALS
	// USER DOESN'T EXISTS
	// INTERNAL SERVER ERROR

	if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(req.Username) {
		return nil, ErrInvalidUsernameFormat
	} else if !utils.IsPasswordInFormat(req.Password) {
		// No need to give idea about password pattern to user
		return nil, ErrInvalidCredentials
	}

	isUserVerified, err := db.VerifyIdPass(&pb.User{Username: req.Username, Password: req.Password})

	if err != nil {
		if err.Error() == "USER DOESN'T EXISTS" {
			return nil, err
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Verify Pass\nSource: Login()")

		return nil, ErrInternalServerError
	}

	if !isUserVerified {
		return nil, ErrInvalidCredentials
	}

	// Now generate new session token, update session token in db & return it to user
	sessionToken, err := db.UpdateTokenInDBAndReturn(req.Username)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Update Session Token in DB & Return\nSource: Login()")

		return nil, ErrInternalServerError
	}

	return &pb.AuthResponse{SessionToken: sessionToken}, nil
}

func (s *AuthServer) Register(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	// Errors:
	// USERNAME MUST BE B'TWIN 5-20 CHARS & ALPHANUMERIC
	// PASSWORD MUST BE B'TWIN 8-20 CHARS, ATLEAST 1 UPPER, LOWER CASE & SPECIAL CHARACTER
	// USERNAME IS ALREADY USED
	// INTERNAL SERVER ERROR

	if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(req.Username) {
		return nil, ErrInvalidUsernameFormat
	} else if !utils.IsPasswordInFormat(req.Password) {
		return nil, ErrInvalidPasswordFormat
	}

	sessionToken := utils.TokenGenerator()

	// Checking whether the generated session token is already used or not
	// If so then generate another session token & keep this loop until we find a unique session token
	for {
		isSessionTokenDuplicate, err := db.IsSessionTokenValid(sessionToken)
		if err != nil {
			log.Println("Error:", err)
			log.Println("Description: Cannot Validate Session Token\nSource: Register()")
			return nil, ErrInternalServerError
		}
		if !isSessionTokenDuplicate {
			break
		}
		sessionToken = utils.TokenGenerator()
		log.Println("Duplicate Token:", sessionToken)
	}

	user := &pb.User{Username: req.Username, Password: req.Password, Token: sessionToken}
	err := db.AddUser(user)
	if err != nil {
		if err.Error() == "USERNAME IS ALREADY USED" {
			return nil, ErrUsernameAlreadyUsedError
		}

		log.Print("Error:", err)
		log.Println("Description: Cannot add user in db\nSource: Register()")
		return nil, ErrInternalServerError
	}

	return &pb.AuthResponse{SessionToken: sessionToken}, nil
}

func (s *AuthServer) VerifySessionToken(ctx context.Context, req *pb.ValidationRequest) (*pb.ValidationResponse, error) {
	// Errors:
	// INTERNAL SERVER ERROR

	if req.SessionToken == "" {
		return &pb.ValidationResponse{IsUserVerified: false}, nil
	}

	isTokenValid, err := db.IsSessionTokenValid(req.SessionToken)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Validate Session Token\nSource: VerifySessionToken()")
		return nil, ErrInternalServerError
	}
	if isTokenValid {
		return &pb.ValidationResponse{IsUserVerified: true}, nil
	}
	return &pb.ValidationResponse{IsUserVerified: false}, nil
}

func (s *AuthServer) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	// Errors:
	// INTERNAL SERVER ERROR
	// INVALID SESSION TOKEN

	if req.SessionToken == "" {
		return nil, ErrInvalidSessionToken
	}

	err := db.DeleteTokenInDB(req.SessionToken)
	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, ErrInvalidSessionToken
		}

		log.Println("Error:", err)
		log.Println("Description: Cannot delete Session Token from DB\nSource: Logout()")

		return nil, ErrInternalServerError
	}

	return &pb.LogoutResponse{IsUserLoggedOut: true}, nil
}
