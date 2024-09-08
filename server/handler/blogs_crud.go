package handler

import (
	"context"
	"log"
	"net/http"
	"regexp"

	"github.com/MohitSilwal16/Nemuda/server/db"
	pb "github.com/MohitSilwal16/Nemuda/server/pb"
	"github.com/MohitSilwal16/Nemuda/server/utils"
)

type BlogsServer struct {
	pb.UnimplementedBlogsServiceServer
}

func (s *BlogsServer) PostBlog(ctx context.Context, req *pb.PostBlogRequest) (*pb.PostBlogResponse, error) {
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

	username, err := db.GetUsernameBySessionToken(req.SessionToken)

	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, ErrInvalidSessionToken
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot get Username from SessionToken\nSource: PostBlog()")

		return nil, ErrInternalServerError
	}

	isTagValid := utils.Contains(tagsList, req.Tag)
	if !isTagValid {
		return nil, ErrInvalidBlogTag
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9 ,'"&]*$`).MatchString(req.Title) {
		return nil, ErrInvalidBlogTitleFormat
	}
	if len(req.Description) < 4 || len(req.Description) > 50 {
		return nil, ErrInvalidBlogDescriptionFormat
	}

	isMalicious := utils.IsMessageMalicious(req.Description)

	if isMalicious {
		return nil, ErrXSSDetected
	}

	image := req.GetImageData()
	contentType := http.DetectContentType(image)

	// Check if the content type is an allowed image type
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
	}

	if !allowedTypes[contentType] {
		return nil, ErrInvalidBlogImageFileType
	}

	if len(image) > MAX_BLOG_IMAGE_SIZE {
		return nil, ErrInvalidBlogImageSize
	}

	filename := req.Title + ".png"

	// Save file in S3 Bucket
	imagePath, err := db.UploadImageToAWS(utils.BytesToMultipartFile(image, filename), filename)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot upload images to S3\nSource: PostBlog()")

		return nil, ErrInternalServerError
	}

	blog := pb.Blog{
		Username:      username,
		Title:         req.Title,
		Tag:           req.Tag,
		Description:   req.Description,
		Likes:         0,
		LikedUsername: []string{},
		Comments:      []*pb.Comment{},
		ImagePath:     imagePath,
	}

	err = db.AddBlog(&blog)

	if err != nil {
		if err.Error() == "TITLE IS ALREADY USED" {
			return nil, ErrInvalidBlogTitleAlreadyUsed
		}

		log.Println("Error:", err)
		if err.Error() == "result.InsertedID is empty" {
			log.Println("Description: result.InsertedID is Empty after Blog is added\nSource: PostBlog()")
		} else {
			log.Println("Description: Cannot Add Blog in DB\nSource: PostBlog()")
		}
		return nil, ErrInternalServerError
	}
	return &pb.PostBlogResponse{IsBlogAdded: true}, nil
}

func (s *BlogsServer) GetBlogsByTagWithPagination(ctx context.Context, req *pb.GetBlogsRequest) (*pb.GetBlogsResponse, error) {
	// Errors:
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR
	// INVALID BLOG TAG
	// INVALID OFFSET

	if req.Offset < 0 {
		return nil, ErrInvalidOffset
	}

	isSessionTokenValid, err := db.IsSessionTokenValid(req.SessionToken)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Validate Session Token\nSource: GetBlogsByTagWithPagination()")

		return nil, ErrInternalServerError
	}
	if !isSessionTokenValid {
		return nil, ErrInvalidSessionToken
	}

	if req.Tag != "All" {
		isTagValid := utils.Contains(tagsList, req.Tag)
		if !isTagValid {
			return nil, ErrInvalidBlogTag
		}
	}

	blogs, err := db.GetBlogsByTagWithOffset(req.Tag, int(req.Offset), BLOG_LIMIT)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot fetch blogs from db with offset\nSource: GetBlogsByTagWithPagination()")

		return nil, ErrInternalServerError
	}

	if len(blogs) == 0 {
		return &pb.GetBlogsResponse{Blogs: nil, NextOffset: -1}, nil
	}
	if len(blogs) < BLOG_LIMIT {
		return &pb.GetBlogsResponse{Blogs: blogs, NextOffset: -1}, nil
	}

	return &pb.GetBlogsResponse{Blogs: blogs, NextOffset: req.Offset + BLOG_LIMIT}, nil
}

func (s *BlogsServer) UpdateBlog(ctx context.Context, req *pb.UpdateBlogRequest) (*pb.UpdateBlogResponse, error) {
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

	username, err := db.GetUsernameBySessionToken(req.SessionToken)

	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, ErrInvalidSessionToken
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot get Username from SessionToken\nSource: UpdateBlog()")

		return nil, ErrInternalServerError
	}

	isNewBlogTitleAlreadyUsed, err := db.SearchBlogByTitle(req.NewTitle)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Check whether new blog title was already used or not\nSource: UpdateBlog()")

		return nil, ErrInternalServerError
	}

	if isNewBlogTitleAlreadyUsed {
		return nil, ErrInvalidBlogTitleAlreadyUsed
	}

	isTagValid := utils.Contains(tagsList, req.NewTag)
	if !isTagValid {
		return nil, ErrInvalidBlogTag
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9 ,'"&]*$`).MatchString(req.NewTitle) {
		return nil, ErrInvalidBlogTitleFormat
	}
	if len(req.NewDescription) < 4 || len(req.NewDescription) > 50 {
		return nil, ErrInvalidBlogDescriptionFormat
	}

	isMalicious := utils.IsMessageMalicious(req.NewDescription)

	if isMalicious {
		return nil, ErrXSSDetected
	}

	image := req.GetNewImageData()
	contentType := http.DetectContentType(image)

	// Check if the content type is an allowed image type
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
	}

	if !allowedTypes[contentType] {
		return nil, ErrInvalidBlogImageFileType
	}

	if len(image) > MAX_BLOG_IMAGE_SIZE {
		return nil, ErrInvalidBlogImageSize
	}

	err = db.DeleteImageFromAWS(req.OldTitle + ".png")
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Delete Blog Image from S3\nSource: UpdateBlog()")

		return nil, ErrInternalServerError
	}

	filename := req.NewTitle + ".png"
	newImagePath, err := db.UploadImageToAWS(utils.BytesToMultipartFile(image, filename), filename)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Upload Blog Image to S3\nSource: UpdateBlog()")

		return nil, ErrInternalServerError
	}

	err = db.UpdateBlog(req.OldTitle, username, req.NewTitle, req.NewDescription, newImagePath, req.NewTag)
	if err != nil {
		if err.Error() == "USER CANNOT UPDATE THIS BLOG" {
			return nil, err
		}
		if err.Error() == "BLOG NOT FOUND" {
			return nil, ErrBlogNotFound
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Update Blog in DB\nSource: UpdateBlog()")

		return nil, ErrInternalServerError
	}

	return &pb.UpdateBlogResponse{IsBlogUpdated: true}, nil
}

func (s *BlogsServer) DeleteBlog(ctx context.Context, req *pb.DeleteBlogRequest) (*pb.DeleteBlogResponse, error) {
	// Errors:
	// USER CANNOT DELETE THIS BLOG
	// BLOG NOT FOUND
	// INTERNAL SERVER ERROR
	// INVALID SESSION TOKEN

	username, err := db.GetUsernameBySessionToken(req.SessionToken)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Error while fetching username from session token\n Source: DeleteBlog()")

		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, ErrInvalidSessionToken
		}
		return nil, ErrInternalServerError
	}

	err = db.DeleteBlog(req.Title, username)
	if err != nil {
		if err.Error() == "USER CANNOT DELETE THIS BLOG" {
			return nil, err
		}
		if err.Error() == "BLOG NOT FOUND" {
			return nil, ErrBlogNotFound
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Delete Blog from DB\nSource: DeleteBlog()")

		return nil, ErrInternalServerError
	}

	err = db.DeleteImageFromAWS(req.Title + ".png")

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Delete Blog Image from S3\nSource: DeleteBlog()")

		return nil, ErrInternalServerError
	}
	return &pb.DeleteBlogResponse{IsBlogDeleted: true}, nil
}
