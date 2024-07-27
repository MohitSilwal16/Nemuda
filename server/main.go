package main

import (
	"log"
	"net"
	"os"

	"github.com/MohitSilwal16/Nemuda/server/db"
	"github.com/MohitSilwal16/Nemuda/server/handler"
	pb "github.com/MohitSilwal16/Nemuda/server/pb"
	"github.com/MohitSilwal16/Nemuda/server/utils"
	"google.golang.org/grpc"
)

const BASE_URL = "0.0.0.0:8080"

func init() {
	err := db.Init_MariaDB()

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	err = db.Init_Mongo()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	err = db.Init_S3()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func main() {
	lis, err := net.Listen("tcp", BASE_URL)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Listen TCP at 8080 PORT")
		return
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(utils.StructuredLoggerInterceptor()),
	)

	pb.RegisterAuthServiceServer(s, &handler.AuthServer{})
	pb.RegisterUserServiceServer(s, &handler.UserServer{})
	pb.RegisterBlogsServiceServer(s, &handler.BlogsServer{})

	log.Println("Running GRPC Server on", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Serve at 8080 PORT")
		return
	}
}
