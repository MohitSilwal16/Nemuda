package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/MohitSilwal16/Nemuda/server/db"
	"github.com/MohitSilwal16/Nemuda/server/models"
	"github.com/MohitSilwal16/Nemuda/server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// Tags' slice
var tagsList = []string{"Political", "Technical", "Educational", "Geographical", "Programming", "Other"}

const BLOG_LIMIT = 3

func VerifySessionToken(ctx *gin.Context) {
	// 200 => Session Token is Valid
	// 201 => Session Token is Invalid
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Param("sessionToken")

	if sessionToken == "" {
		ctx.JSON(200, gin.H{
			"message": "Session Token is Valid",
		})
		return
	}

	isTokenValid, err := db.IsSessionTokenValid(sessionToken)

	if err != nil {
		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	if isTokenValid {
		ctx.JSON(200, gin.H{
			"message": "Session Token is Valid",
		})
		return
	}
	ctx.JSON(201, gin.H{
		"message": "Session Token is Invalid",
	})
}

func Register(ctx *gin.Context) {
	// AVOID USING 204 BECAUSE IT DOESN'T SEND ANY CONTENT OR BODY

	// 200 => Registered Successfully
	// 201 => User data is not in format
	// 202 => Username or Password is Empty
	// 203 => Username is not in required format
	// 205 => Password is not in required format
	// 206 => Username is already used
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(ctx.Request.Body).Decode(&user)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot read JSON data in user object")

		ctx.JSON(201, gin.H{
			"message": "Provided data is not in correct format(JSON).",
		})
		return
	}

	if user.Username == "" || user.Password == "" {
		ctx.JSON(202, gin.H{
			"message": "Username or Password is Empty",
		})
	} else if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(user.Username) {
		ctx.JSON(203, gin.H{
			"message": "Username must have between 5 and 20 alphanumeric characters.",
		})
	} else if !utils.IsPasswordInFormat(user.Password) {
		ctx.JSON(205, gin.H{
			"message": "Password must be between 8 and 20 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character",
		})
	} else {
		// Generate session token
		sessionToken := utils.TokenGenerator()

		// Checking whether the generated session token is already used or not
		// If so then generate another session token & keep this loop until we find a unique session token
		for {
			isSessionTokenDuplicate, err := db.IsSessionTokenValid(sessionToken)
			if err != nil {
				log.Println("Error: ", err)
				ctx.JSON(500, gin.H{
					"message": "Internal Server Error",
				})
				return
			}
			if !isSessionTokenDuplicate {
				break
			}
			sessionToken = utils.TokenGenerator()
			log.Println("Duplicate Token " + sessionToken)
		}

		user.Token = sessionToken

		// Add user to database
		err := db.AddUser(user)
		if err != nil {
			if err.Error() == "USERNAME IS ALREADY USED" {
				ctx.JSON(206, gin.H{
					"message": "Username is already used",
				})
				return
			}

			log.Print("Error: ", err)
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message":      "Registered Successfully",
			"sessionToken": user.Token,
		})
	}
}

func Login(ctx *gin.Context) {
	// AVOID USING 204 BECAUSE IT DOESN'T SEND ANY CONTENT OR BODY

	// 200 => Login Successful
	// 201 => User data is not in format
	// 202 => Username or Password is Empty
	// 203 => Username is not in required format
	// 205 => User doesn't exists
	// 206 => Invalid Credentials
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(ctx.Request.Body).Decode(&user)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot read JSON data in user object")

		ctx.JSON(201, gin.H{
			"message": "Provided data is not in correct format(JSON).",
		})
		return
	}

	if user.Username == "" || user.Password == "" {
		ctx.JSON(202, gin.H{
			"message": "Username or Password is Empty.",
		})
		return
	} else if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(user.Username) {
		ctx.JSON(203, gin.H{
			"message": "Username must have between 5 and 20 alphanumeric characters.",
		})
		return
	} else if !utils.IsPasswordInFormat(user.Password) {
		// Don't want to give idea to anonymous user idea about password pattern
		ctx.JSON(206, gin.H{
			"message": "Invalid Credentials",
		})
		return
	}

	// Verify Username & Password
	userVerified, err := db.VerifyIdPass(user)

	if err != nil {
		if err.Error() == "USER DOESN'T EXISTS" {
			ctx.JSON(205, gin.H{
				"message": "User doesn't exists.",
			})
			return
		}
		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !userVerified {
		ctx.JSON(206, gin.H{
			"message": "Invalid Credentials.",
		})
		return
	}

	// Now generate new session token, update session token in db & return it to user
	sessionToken, err := db.UpdateTokenInDBAndReturn(user.Username)

	if err != nil {
		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":      "Login Successful",
		"sessionToken": sessionToken,
	})
}

func Logout(ctx *gin.Context) {
	// 200 => Log out Successful
	// 201 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Param("sessionToken")
	err := db.DeleteTokenInDB(sessionToken)

	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			ctx.JSON(201, gin.H{
				"message": "Invalid Session Token",
			})
			return
		}

		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Log out Sucessful",
	})
}

func SearchUser(ctx *gin.Context) {
	// 200 => User found (Username is already used)
	// 201 => User not found (Username is not used yet)
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	username := ctx.Param("username")

	if len(username) < 5 {
		ctx.JSON(201, gin.H{
			"message": "User not found",
		})
		return
	}

	usernameFound, err := db.DoesUserExists(username)

	if err != nil {
		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if usernameFound {
		ctx.JSON(200, gin.H{
			"message": "User found",
		})
		return
	}
	ctx.JSON(201, gin.H{
		"message": "User not found",
	})
}

func SearchBlogByTitle(ctx *gin.Context) {
	// 200 => Blog found
	// 201 => Blog not found
	// 202 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Query("sessionToken")

	if sessionToken == "" {
		ctx.JSON(202, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	isSessionTokenValid, err := db.IsSessionTokenValid(sessionToken)
	if err != nil {
		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !isSessionTokenValid {
		ctx.JSON(202, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	title := ctx.Param("title")

	blog, err := db.GetBlogByTitle(title)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			ctx.JSON(201, gin.H{
				"message": "Blog not found",
			})
			return
		}
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(200, blog)
}

func PostBlog(ctx *gin.Context) {
	// AVOID USING 204 BECAUSE IT DOESN'T SEND ANY CONTENT OR BODY

	// 200 => Blog added
	// 201 => Title, Description, Tag is not in requested format(JSON)
	// 202 => Title is already used
	// 203 => Invalid Session Token
	// 205 => Image is not in proper format
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Query("sessionToken")

	if sessionToken == "" {
		ctx.JSON(203, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	isSessionTokenValid, err := db.IsSessionTokenValid(sessionToken)
	if err != nil {
		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !isSessionTokenValid {
		ctx.JSON(203, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	// Parse the multipart form
	if err := ctx.Request.ParseMultipartForm(10 << 20); err != nil { // 10 MB max
		ctx.JSON(205, gin.H{"error": "File too large"})
		return
	}
	// log.Printf("Multipart form files: %#v", ctx.Request.MultipartForm.File)

	// Retrieve the JSON data from form data
	data := ctx.Request.FormValue("data")
	var blog models.Blog
	if err := json.Unmarshal([]byte(data), &blog); err != nil {
		ctx.JSON(201, gin.H{
			"message": "Title, Description, Tag is not in proper format"})
		return
	}

	blog.Username, err = db.GetUsernameBySessionToken(sessionToken)
	blog.Comments = []models.Comment{}
	blog.LikedUsername = []string{}
	blog.Likes = 0

	log.Printf("Blog Data: %#v", blog)

	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			ctx.JSON(203, gin.H{
				"message": "Invalid Session Token",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9 ,'"&]*$`).MatchString(blog.Title) {
		response := "Title should be alphanumeric b'twin 5-20 chars"
		ctx.JSON(201, gin.H{
			"message": response},
		)
	} else if !utils.Contains(tagsList, blog.Tag) {
		response := "Unknown tag"
		ctx.JSON(201, gin.H{
			"message": response},
		)
	} else if len(blog.Description) < 4 || len(blog.Description) > 50 {
		response := "Desc: Min 4 letters & Max 50 letters"
		ctx.JSON(201, gin.H{
			"message": response},
		)
	} else {
		image, header, err := ctx.Request.FormFile("file")
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Cannot fetch image from user")

			ctx.JSON(205, gin.H{
				"message": "Cannot fetch Image",
			})
			return
		}

		// For example, save the file and log the additional data
		filename := header.Filename
		// Save the file or process it as needed
		fmt.Printf("Received file: %s\n", filename)

		// Read the first 512 bytes of the file
		buffer := make([]byte, 512)
		_, err = image.Read(buffer)
		if err != nil {
			ctx.JSON(207, gin.H{"error": "Unable to read file"})
			return
		}

		// Reset the file reader to the beginning
		image.Seek(0, 0)

		contentType := http.DetectContentType(buffer)
		log.Println("Detected Content Type:", contentType)

		// Check if the content type is an allowed image type
		allowedTypes := map[string]bool{
			"image/jpeg": true,
			"image/jpg":  true,
			"image/png":  true,
		}

		if !allowedTypes[contentType] {
			ctx.JSON(205, gin.H{
				"message": "Invalid file type, upload a JPG, JPEG, or PNG image",
			})
			return
		}

		maxSize := 2 * 1024 * 1024 // 2MB
		if ctx.Request.MultipartForm.File["file"][0].Size > int64(maxSize) {
			ctx.JSON(205, gin.H{
				"message": "Image size exceeds 2 MB",
			})
			return
		}

		out, err := os.Create(blog.ImagePath)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Unable to save the file"})
			return
		}
		defer out.Close()

		if _, err := io.Copy(out, image); err != nil {
			ctx.JSON(500, gin.H{"error": "Unable to copy the file"})
			return
		}

		err = db.AddBlog(blog)

		if err != nil {
			log.Println("Error:", err)
			if err.Error() == "TITLE IS ALREADY USED" {
				ctx.JSON(202, gin.H{
					"message": "Title is already used",
				})
				return
			} else if err.Error() == "result.InsertedID is empty" {
				log.Println("Description: result.InsertedID is empty")
			}
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"message": "Blog Added",
		})
	}
}

func IsBlogLikedByUser(ctx *gin.Context) {
	// 200 => Blog liked
	// 201 => Blog not liked
	// 202 => Blog not found
	// 203 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Query("sessionToken")

	if sessionToken == "" {
		ctx.JSON(203, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	isSessionTokenValid, err := db.IsSessionTokenValid(sessionToken)
	if err != nil {
		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !isSessionTokenValid {
		ctx.JSON(203, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	title := ctx.Param("title")

	username, err := db.GetUsernameBySessionToken(sessionToken)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error while fetching username from session token")

		if err.Error() == "INVALID SESSION TOKEN" {
			ctx.JSON(203, gin.H{
				"message": "Invalid Session Token",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}

	doesBlogExists, err := db.SearchBlogByTitle(title)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot search blog by title")

		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !doesBlogExists {
		ctx.JSON(202, gin.H{
			"message": "Blog Not Found",
		})
		return
	}

	isBlogLiked, err := db.IsBlogLikedByUser(title, username)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot like blog")

		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if isBlogLiked {
		ctx.JSON(200, gin.H{
			"message": "Blog Liked",
		})
		return
	}
	ctx.JSON(201, gin.H{
		"message": "Blog Not Liked",
	})
}

func LikeBlog(ctx *gin.Context) {
	// 200 => Blog liked
	// 201 => Blog already liked
	// 202 => Blog not found
	// 203 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Query("sessionToken")

	if sessionToken == "" {
		ctx.JSON(203, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	isSessionTokenValid, err := db.IsSessionTokenValid(sessionToken)
	if err != nil {
		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !isSessionTokenValid {
		ctx.JSON(203, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	title := ctx.Param("title")
	if title == "" {
		ctx.JSON(201, gin.H{
			"message": "Title is Empty",
		})
		return
	}

	username, err := db.GetUsernameBySessionToken(sessionToken)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error while fetching username from session token")

		if err.Error() == "INVALID SESSION TOKEN" {
			ctx.JSON(203, gin.H{
				"message": "Invalid Session Token",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}

	err = db.LikeBlog(title, username)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot like blog")

		if err.Error() == "BLOG NOT FOUND" {
			ctx.JSON(201, gin.H{
				"message": "Blog Not Found",
			})
		} else if err.Error() == "BLOG ALREADY LIKED" {
			ctx.JSON(202, gin.H{
				"message": "Blog Already Liked",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Blog Liked",
	})
}

func DislikeBlog(ctx *gin.Context) {
	// 200 => Blog disliked
	// 201 => Blog already disliked
	// 202 => Blog not found
	// 203 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Query("sessionToken")

	if sessionToken == "" {
		ctx.JSON(203, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	isSessionTokenValid, err := db.IsSessionTokenValid(sessionToken)
	if err != nil {
		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !isSessionTokenValid {
		ctx.JSON(203, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	title := ctx.Param("title")
	if title == "" {
		ctx.JSON(201, gin.H{
			"message": "Title is Empty",
		})
		return
	}

	username, err := db.GetUsernameBySessionToken(sessionToken)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error while fetching username from session token")

		if err.Error() == "INVALID SESSION TOKEN" {
			ctx.JSON(203, gin.H{
				"message": "Invalid Session Token",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}

	err = db.DislikeBlog(title, username)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot like blog")

		if err.Error() == "BLOG NOT FOUND" {
			ctx.JSON(201, gin.H{
				"message": "Blog Not Found",
			})
		} else if err.Error() == "BLOG ALREADY DISLIKED" {
			ctx.JSON(202, gin.H{
				"message": "Blog Already Disliked",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Blog Disliked",
	})
}

func AddComment(ctx *gin.Context) {
	// 200 => Comment Added
	// 201 => Comment Description or Title is Empty
	// 202 => Blog not found
	// 203 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Query("sessionToken")

	if sessionToken == "" {
		ctx.JSON(202, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	isSessionTokenValid, err := db.IsSessionTokenValid(sessionToken)
	if err != nil {
		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !isSessionTokenValid {
		ctx.JSON(202, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	title := ctx.Query("title")

	if title == "" {
		ctx.JSON(201, gin.H{
			"message": "Title is Empty",
		})
		return
	}

	username, err := db.GetUsernameBySessionToken(sessionToken)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error while fetching username from session token")

		if err.Error() == "INVALID SESSION TOKEN" {
			ctx.JSON(202, gin.H{
				"message": "Invalid Session Token",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}

	commentDescription := ctx.Param("comment")

	if commentDescription == "" {
		ctx.JSON(201, gin.H{
			"message": "Comment Description is Empty",
		})
		return
	}

	if len(commentDescription) < 5 || len(commentDescription) > 50 {
		ctx.JSON(201, gin.H{
			"message": "Comment: Min 5 & Max 50 letters",
		})
		return
	}

	comment := models.Comment{
		Username:    username,
		Description: commentDescription,
	}

	err = db.AddComment(title, comment)

	if err != nil {
		if err.Error() == "BLOG NOT FOUND" {
			ctx.JSON(202, gin.H{
				"message": "Blog not found",
			})
		} else {
			log.Println("Error :", err)
			log.Println("Description: Cannot add comment")

			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Comment Added",
	})
}

func UpdateBlog(ctx *gin.Context) {
	// AVOID USING 204 BECAUSE IT DOESN'T SEND ANY CONTENT OR BODY

	// 200 => Blog Updated
	// 201 => Data is not in correct format
	// 202 => User cannot update this blog
	// 203 => Blog not found
	// 205 => Invalid Session Token
	// 206 => Blog Title is already used
	// 207 => Image is not in proper format
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Query("sessionToken")

	if sessionToken == "" {
		ctx.JSON(205, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	isSessionTokenValid, err := db.IsSessionTokenValid(sessionToken)
	if err != nil {
		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !isSessionTokenValid {
		ctx.JSON(205, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	oldTitle := ctx.Query("title")

	if oldTitle == "" {
		ctx.JSON(201, gin.H{
			"message": "Title is Empty",
		})
		return
	}

	username, err := db.GetUsernameBySessionToken(sessionToken)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error while fetching username from session token")

		if err.Error() == "INVALID SESSION TOKEN" {
			ctx.JSON(205, gin.H{
				"message": "Invalid Session Token",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}

	// Retrieve the JSON data from form data
	data := ctx.Request.FormValue("data")
	var blog models.Blog
	if err := json.Unmarshal([]byte(data), &blog); err != nil {
		ctx.JSON(201, gin.H{
			"message": "Title, Description, Tag is not in proper format"})
		return
	}

	blog.Comments = []models.Comment{}
	blog.LikedUsername = []string{}
	blog.Likes = 0

	log.Printf("Blog Data: %#v", blog)

	if !regexp.MustCompile(`^[a-zA-Z0-9 ,'"&]*$`).MatchString(blog.Title) {
		response := "Title should be alphanumeric b'twin 5-20 chars"
		ctx.JSON(201, gin.H{
			"message": response,
		})
	} else if !utils.Contains(tagsList, blog.Tag) {
		response := "Unknown tag"
		ctx.JSON(201, gin.H{
			"message": response,
		})
	} else if len(blog.Description) < 4 || len(blog.Description) > 50 {
		response := "Desc: Min 4 letters & Max 50 letters"
		ctx.JSON(201, gin.H{
			"message": response,
		})
	} else {
		doesBlogExists, err := db.SearchBlogByTitle(blog.Title)

		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Cannot search blog by title")

			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		if doesBlogExists {
			ctx.JSON(206, gin.H{
				"message": "Blog Title is already used",
			})
			return
		}

		image, header, err := ctx.Request.FormFile("file")
		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Cannot fetch image from user")

			ctx.JSON(207, gin.H{
				"message": "Cannot fetch Image",
			})
			return
		}

		// For example, save the file and log the additional data
		filename := header.Filename
		// Save the file or process it as needed
		fmt.Printf("Received file: %s\n", filename)

		// Read the first 512 bytes of the file
		buffer := make([]byte, 512)
		_, err = image.Read(buffer)
		if err != nil {
			ctx.JSON(207, gin.H{"error": "Unable to read file"})
			return
		}

		// Reset the file reader to the beginning
		image.Seek(0, 0)

		contentType := http.DetectContentType(buffer)
		log.Println("Detected Content Type:", contentType)

		// Check if the content type is an allowed image type
		allowedTypes := map[string]bool{
			"image/jpeg": true,
			"image/jpg":  true,
			"image/png":  true,
		}

		if !allowedTypes[contentType] {
			ctx.JSON(207, gin.H{
				"message": "Invalid file type, upload a JPG, JPEG, or PNG image",
			})
			return
		}

		maxSize := 2 * 1024 * 1024 // 2MB
		if ctx.Request.MultipartForm.File["file"][0].Size > int64(maxSize) {
			ctx.JSON(207, gin.H{
				"message": "Image size exceeds 2 MB",
			})
			return
		}

		out, err := os.Create(blog.ImagePath)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Unable to save the file"})
			return
		}
		defer out.Close()

		if _, err := io.Copy(out, image); err != nil {
			ctx.JSON(500, gin.H{"error": "Unable to copy the file"})
			return
		}

		err = db.UpdateBlog(oldTitle, username, blog.Title, blog.Description, blog.ImagePath, blog.Tag)

		if err != nil {
			if err.Error() == "USER CANNOT UPDATE THIS BLOG" {
				ctx.JSON(202, gin.H{
					"message": "User cannot update this blog",
				})
			} else if err.Error() == "BLOG NOT FOUND" {
				ctx.JSON(203, gin.H{
					"message": "Blog Not Found",
				})
			} else {
				log.Println("Error: ", err)
				log.Println("Description: Cannot Update Blog")
				ctx.JSON(500, gin.H{
					"message": "Internal Server Error",
				})
			}
			return
		}
		err = ctx.SaveUploadedFile(header, blog.ImagePath)

		if err != nil {
			log.Println("Error: ", err)
			log.Println("Description: Cannot save image of blog")

			fmt.Fprint(ctx.Writer, "Image of blog cannot be saved")
		}
		oldImagePath := "./static/images/blogs/" + oldTitle + ".png"

		err = os.Remove(oldImagePath)

		if err != nil {
			if !os.IsNotExist(err) {
				log.Println("Error: ", err)
				log.Println("Description: Cannot delete ", oldImagePath)
			}
			// No need to return
		}

		ctx.JSON(200, gin.H{
			"message": "Blog Updated",
		})
	}
}

func DeleteBlog(ctx *gin.Context) {
	// AVOID USING 204 BECAUSE IT DOESN'T SEND ANY CONTENT OR BODY

	// 200 => Blog Deleted
	// 201 => Data is not in correct format
	// 202 => User cannot delete this blog
	// 203 => Blog not found
	// 205 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Query("sessionToken")

	if sessionToken == "" {
		ctx.JSON(205, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	isSessionTokenValid, err := db.IsSessionTokenValid(sessionToken)
	if err != nil {
		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !isSessionTokenValid {
		ctx.JSON(205, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	title := ctx.Query("title")

	if title == "" {
		ctx.JSON(201, gin.H{
			"message": "Title is Empty",
		})
		return
	}

	username, err := db.GetUsernameBySessionToken(sessionToken)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error while fetching username from session token")

		if err.Error() == "INVALID SESSION TOKEN" {
			ctx.JSON(205, gin.H{
				"message": "Invalid Session Token",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}

	err = db.DeleteBlog(title, username)

	if err != nil {
		if err.Error() == "USER CANNOT DELETE THIS BLOG" {
			ctx.JSON(202, gin.H{
				"message": "User cannot delete this blog",
			})
		} else if err.Error() == "BLOG NOT FOUND" {
			ctx.JSON(203, gin.H{
				"message": "Blog Not Found",
			})
		} else {
			log.Println("Error: ", err)
			log.Println("Description: Cannot Delete Blog")

			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}

	imagePath := "./static/images/blogs/" + title + ".png"

	err = os.Remove(imagePath)

	if err != nil {
		if !os.IsNotExist(err) {
			log.Println("Error: ", err)
			log.Println("Description: Cannot delete ", imagePath)
		}
		// No need to return
	}

	ctx.JSON(200, gin.H{
		"message": "Blog Deleted",
	})
}

func CanUserUpdate_DeleteBlog(ctx *gin.Context) {
	// 200 => User can update/delete blog
	// 201 => User cannot update/delete blog
	// 202 => Blog not found
	// 203 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Query("sessionToken")

	if sessionToken == "" {
		ctx.JSON(203, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	isSessionTokenValid, err := db.IsSessionTokenValid(sessionToken)
	if err != nil {
		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !isSessionTokenValid {
		ctx.JSON(203, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	title := ctx.Query("title")

	if title == "" {
		ctx.JSON(201, gin.H{
			"message": "Title is Empty",
		})
		return
	}

	username, err := db.GetUsernameBySessionToken(sessionToken)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error while fetching username from session token")

		if err.Error() == "INVALID SESSION TOKEN" {
			ctx.JSON(203, gin.H{
				"message": "Invalid Session Token",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}

	isUpdatble_Deletable, err := db.IsBlogUpdatable_Deletable(title, username)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot like blog")

		if err.Error() == "BLOG NOT FOUND" {
			ctx.JSON(201, gin.H{
				"message": "Blog Not Found",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}

	if isUpdatble_Deletable {
		ctx.JSON(200, gin.H{
			"message": "User can update or delete blog",
		})
		return
	}
	ctx.JSON(201, gin.H{
		"message": "User cannot update or delete blog",
	})
}

func GetBlogsByTag(ctx *gin.Context) {
	// AVOID USING 204 BECAUSE IT DOESN'T SEND ANY CONTENT OR BODY

	// 200 => Blogs found
	// 202 => Invalid Session Token
	// 203 => No more blogs available
	// 205 => Invalid Offset
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Query("sessionToken")

	if sessionToken == "" {
		ctx.JSON(202, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	isSessionTokenValid, err := db.IsSessionTokenValid(sessionToken)
	if err != nil {
		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !isSessionTokenValid {
		ctx.JSON(202, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	tag := ctx.Param("tag")
	tag = strings.Title(tag)

	offsetString := ctx.Query("offset")

	offsetInt, err := strconv.Atoi(offsetString)

	if err != nil {
		ctx.JSON(205, gin.H{
			"message": "Offset shouldn't be negative integer",
		})
		return
	}

	if offsetInt < 0 {
		ctx.JSON(205, gin.H{
			"message": "Offset shouldn't be negative integer",
		})
		return
	}

	blogs, err := db.GetBlogsByTagWithOffset(tag, offsetInt, BLOG_LIMIT)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot fetch blogs from db with offset")
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if len(blogs) == 0 {
		ctx.JSON(203, gin.H{
			"message": fmt.Sprintf("No more blogs for %s available", tag),
		})
		return
	}
	if len(blogs) < BLOG_LIMIT {
		ctx.JSON(200, gin.H{
			"blogs":      blogs,
			"nextOffset": "-1",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"blogs":      blogs,
		"nextOffset": strconv.Itoa(offsetInt + BLOG_LIMIT),
	})
}

func GetUsernameBySessionToken(ctx *gin.Context) {
	// 200 => Username found
	// 201 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Param("sessionToken")

	if sessionToken == "" {
		ctx.JSON(201, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	username, err := db.GetUsernameBySessionToken(sessionToken)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error while fetching username from session token")

		if err.Error() == "INVALID SESSION TOKEN" {
			ctx.JSON(201, gin.H{
				"message": "Invalid Session Token",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}
	ctx.JSON(200, gin.H{
		"username": username,
	})
}

func GetMessages(ctx *gin.Context) {
	// 200 => Messages returned
	// 201 => No messages
	// 202 => Invalid Session Token
	// 203 => Invalid Receiver
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Query("sessionToken")

	if sessionToken == "" {
		ctx.JSON(202, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	receiver := ctx.Param("user")

	sender, err := db.GetUsernameBySessionToken(sessionToken)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error while fetching username from session token")

		if err.Error() == "INVALID SESSION TOKEN" {
			ctx.JSON(202, gin.H{
				"message": "Invalid Session Token",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}

	isReceiverValid, err := db.DoesUserExists(receiver)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !isReceiverValid {
		ctx.JSON(203, gin.H{
			"message": "Invalid Receiver",
		})
		return
	}

	messages, err := db.GetMessages(sender, receiver)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Error while fetching messages")

		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if messages == nil {
		ctx.JSON(201, gin.H{
			"message": "No messages with " + receiver,
		})
		return
	}

	// Fetching all messages of receiver indirectly means I read 'em so I should mark 'em as read
	for _, message := range messages {
		if message.Sender == receiver {
			db.ChangeStatusOfMessage(message, "Read")
			message.Status = "Read"
		}
	}

	ctx.JSON(200, messages)
}

func AddMessage(ctx *gin.Context) {
	// 200 => Messages added
	// 201 => Message data isn't in proper format(JSON)
	// 202 => Invalid Session Token
	// 203 => Invalid Receiver
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Query("sessionToken")

	if sessionToken == "" {
		ctx.JSON(202, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	isSessionTokenValid, err := db.IsSessionTokenValid(sessionToken)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot validate session token")
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !isSessionTokenValid {
		ctx.JSON(202, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	var message models.Message
	err = json.NewDecoder(ctx.Request.Body).Decode(&message)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot read json data")

		ctx.JSON(201, gin.H{
			"message": "Provided data is not in correct format(JSON).",
		})
		return
	}

	doesReceiverExists, err := db.DoesUserExists(message.Receiver)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !doesReceiverExists {
		ctx.JSON(203, gin.H{
			"message": "Invalid Receiver",
		})
		return
	}

	err = db.AddMessage(message)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot validate session token")

		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Message Added",
	})
}

func SearchUsersByPattern(ctx *gin.Context) {
	// 200 => Users found
	// 201 => No users found
	// 202 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Query("sessionToken")

	if sessionToken == "" {
		ctx.JSON(202, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	isSessionTokenValid, err := db.IsSessionTokenValid(sessionToken)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot validate session token")
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !isSessionTokenValid {
		ctx.JSON(202, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	searchString := ctx.Query("searchString")

	users, err := db.SearchUsersByPattern(searchString)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Search Users with some pattern")

		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if users == nil {
		ctx.JSON(201, gin.H{
			"message": "No users found",
		})
		return
	}
	ctx.JSON(200, users)
}

// User can revert the status of message
func ChangeStatusOfMessage(ctx *gin.Context) {
	// 200 => Status Changed
	// 201 => Message not found
	// 202 => Invalid Session Token
	// 203 => Invalid New Status
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Query("sessionToken")

	if sessionToken == "" {
		ctx.JSON(202, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	isSessionTokenValid, err := db.IsSessionTokenValid(sessionToken)

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot validate session token")
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if !isSessionTokenValid {
		ctx.JSON(202, gin.H{
			"message": "Invalid Session Token",
		})
		return
	}

	var message models.Message
	err = json.NewDecoder(ctx.Request.Body).Decode(&message)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot read json data")

		ctx.JSON(201, gin.H{
			"message": "Provided data is not in correct format(JSON).",
		})
		return
	}

	newStatus := ctx.Param("newStatus")

	if newStatus != "Read" && newStatus != "Sent" && newStatus != "Delivered" {
		ctx.JSON(203, gin.H{
			"message": "New Status must be either Sent, Delivered or Read",
		})
		return
	}

	err = db.ChangeStatusOfMessage(message, newStatus)

	if err != nil {
		if err.Error() == "MESSAGE NOT FOUND" {
			ctx.JSON(201, gin.H{
				"message": "Message Not Found",
			})
			return
		}

		log.Println("Error:", err)
		log.Println("Description: Cannot Change Status of Message")

		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Status Updated",
	})
}
