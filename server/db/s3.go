package db

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

var uploader *manager.Uploader
var s3Client *s3.Client

func Init_S3() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	// Setup S3 Uploader
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Println("Connection with AWS S3 isn't established")
		return err
	}

	s3Client = s3.NewFromConfig(cfg)

	uploader = manager.NewUploader(s3Client)

	log.Println("Connection with AWS S3 is established")
	return nil
}

func UploadImageToAWS(image multipart.File, fileName string) (string, error) {
	err := godotenv.Load()

	if err != nil {
		return "", err
	}

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("nemuda"),
		Key:    aws.String(fileName),
		Body:   image,
		ACL:    "public-read",
	})

	if err != nil {
		return "", err
	}

	return result.Location, nil
}

func DeleteImageFromAWS(fileName string) error {
	_, err := s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String("nemuda"),
		Key:    aws.String(fileName),
	})

	if err != nil {
		return err
	}

	return nil
}
