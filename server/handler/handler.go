package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/MohitSilwal16/Nemuda/db"
	"github.com/MohitSilwal16/Nemuda/models"
	"github.com/MohitSilwal16/Nemuda/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// Tags' slice
var tagsList = []string{"Political", "Technical", "Educational", "Geographical", "Programming", "Other"}

var BLOG_LIMIT = 1

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

	isTokenValid, err := db.IsSessionTokenDuplicate(sessionToken)

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
			isSessionTokenDuplicate, err := db.IsSessionTokenDuplicate(sessionToken)
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

	usernameFound, err := db.SearchUserByName(username)

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

	isSessionTokenValid, err := db.IsSessionTokenDuplicate(sessionToken)
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

	isSessionTokenValid, err := db.IsSessionTokenDuplicate(sessionToken)
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

	var blog models.Blog
	err = json.NewDecoder(ctx.Request.Body).Decode(&blog)

	if err != nil {
		log.Println("Error: ", err)
		ctx.JSON(201, gin.H{
			"message": "Provided data is not in correct format(JSON).",
		})
		return
	}

	blog.Username, err = db.GetUsernameBySessionToken(sessionToken)

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

	isSessionTokenValid, err := db.IsSessionTokenDuplicate(sessionToken)
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

	isSessionTokenValid, err := db.IsSessionTokenDuplicate(sessionToken)
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

	isSessionTokenValid, err := db.IsSessionTokenDuplicate(sessionToken)
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

	isSessionTokenValid, err := db.IsSessionTokenDuplicate(sessionToken)
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

	isSessionTokenValid, err := db.IsSessionTokenDuplicate(sessionToken)
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
	var blog models.Blog
	err = json.NewDecoder(ctx.Request.Body).Decode(&blog)

	if err != nil {
		log.Println("Error: ", err)
		log.Println("Description: Cannot read JSON data in blog object")

		ctx.JSON(201, gin.H{
			"message": "Provided data is not in correct format(JSON).",
		})
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

		err = db.UpdateBlog(title, username, blog.Title, blog.Description, blog.ImagePath, blog.Tag)

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

	isSessionTokenValid, err := db.IsSessionTokenDuplicate(sessionToken)
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

	isSessionTokenValid, err := db.IsSessionTokenDuplicate(sessionToken)
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
	// 201 => No blog found for the specific tag
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

	isSessionTokenValid, err := db.IsSessionTokenDuplicate(sessionToken)
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
