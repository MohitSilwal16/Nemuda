package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/MohitSilwal16/Nemuda/db"
	"github.com/MohitSilwal16/Nemuda/models"
	"github.com/MohitSilwal16/Nemuda/utils"
	"github.com/gin-gonic/gin"
)

func VerifySessionToken(ctx *gin.Context) {
	// 200 => Session Token is Valid
	// 201 => Session Token is Invalid
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Param("sessionToken")

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

func GetBlogsByTags(ctx *gin.Context) {
	// 200 => Blogs found
	// 201 => No blog found for the specific tag
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

	tag := ctx.Param("tag")
	tag = strings.Title(tag)

	blogs, err := db.GetBlogsByTags(tag)

	if err != nil {
		log.Println("Error: ", err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if blogs == nil {
		ctx.JSON(201, gin.H{
			"message": fmt.Sprintf("No Blogs for '%s' tag", tag),
		})
		return
	}
	ctx.JSON(200, blogs)
}

func AddBlog(ctx *gin.Context) {
	// 200 => Blog added
	// 201 => Title or description is empty
	// 202 => Invalid Session Token
	// 500 => Internal Server Error

	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	sessionToken := ctx.Query("sessionToken")

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

	if blog.Title == "" {
		ctx.JSON(201, gin.H{
			"message": "Provided data is not in correct format(JSON).",
		})
		return
	}
	// Incomplete method
}
