package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"regexp"

	"github.com/MohitSilwal16/Nemuda/db"
	"github.com/MohitSilwal16/Nemuda/models"
	"github.com/MohitSilwal16/Nemuda/utils"
	"github.com/gin-gonic/gin"
)

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

// Temp database
var fakeBlogDB = []models.Blog{
	{Username: "Nimesh", Title: "I love Hitler's wife", Tags: []string{"Political"}, Description: " I'm Nimesh Gadhvi, owner of Gadhvi Airlines. I hereby announce that I like Hilter's wife. I don't know why Hilter wanted to ..."},
	{Username: "Konark", Title: "I love Messi's wife", Tags: []string{"Educational"}, Description: " I'm Nimesh Gadhvi, owner of Gadhvi Airlines. I hereby announce that I like Hilter's wife. I don't know why Hilter wanted to ..."},
	{Username: "Nimesh", Title: "I love Hitler's wife", Tags: []string{"Political"}, Description: " I'm Nimesh Gadhvi, owner of Gadhvi Airlines. I hereby announce that I like Hilter's wife. I don't know why Hilter wanted to ..."},
	{Username: "Konark", Title: "I love Messi's wife", Tags: []string{"Educational"}, Description: " I'm Nimesh Gadhvi, owner of Gadhvi Airlines. I hereby announce that I like Hilter's wife. I don't know why Hilter wanted to ..."},
	{Username: "Nimesh", Title: "I love Hitler's wife", Tags: []string{"Political"}, Description: " I'm Nimesh Gadhvi, owner of Gadhvi Airlines. I hereby announce that I like Hilter's wife. I don't know why Hilter wanted to ..."},
	{Username: "Konark", Title: "I love Messi's wife", Tags: []string{"Educational"}, Description: " I'm Nimesh Gadhvi, owner of Gadhvi Airlines. I hereby announce that I like Hilter's wife. I don't know why Hilter wanted to ..."},
}

// var fakeBlogsDB = map[string][]models.Blog{
// 	"Blogs": {
// 		{Username: "Nimesh", Title: "I love Hitler's wife", Tags: []string{"Political"}, Description: " I'm Nimesh Gadhvi, owner of Gadhvi Airlines. I hereby announce that I like Hilter's wife. I don't know why Hilter wanted to ..."},
// 		{Username: "Konark", Title: "I love Messi's wife", Tags: []string{"Educational"}, Description: " I'm Nimesh Gadhvi, owner of Gadhvi Airlines. I hereby announce that I like Hilter's wife. I don't know why Hilter wanted to ..."},
// 		{Username: "Nimesh", Title: "I love Hitler's wife", Tags: []string{"Political"}, Description: " I'm Nimesh Gadhvi, owner of Gadhvi Airlines. I hereby announce that I like Hilter's wife. I don't know why Hilter wanted to ..."},
// 		{Username: "Konark", Title: "I love Messi's wife", Tags: []string{"Educational"}, Description: " I'm Nimesh Gadhvi, owner of Gadhvi Airlines. I hereby announce that I like Hilter's wife. I don't know why Hilter wanted to ..."},
// 		{Username: "Nimesh", Title: "I love Hitler's wife", Tags: []string{"Political"}, Description: " I'm Nimesh Gadhvi, owner of Gadhvi Airlines. I hereby announce that I like Hilter's wife. I don't know why Hilter wanted to ..."},
// 		{Username: "Konark", Title: "I love Messi's wife", Tags: []string{"Educational"}, Description: " I'm Nimesh Gadhvi, owner of Gadhvi Airlines. I hereby announce that I like Hilter's wife. I don't know why Hilter wanted to ..."},
// 	},
// }

func GetBlogsByTags(ctx *gin.Context) {
	// Set the Content-Type header to "application/json"
	ctx.Header("Content-Type", "application/json")

	tag := ctx.Param("tag")
	log.Println(tag)

	tmpl := template.Must(template.ParseFiles("views/blog.html"))
	if tag == "All" {
		tmpl.Execute(ctx.Writer, fakeBlogDB)
		return
	}

	var filtredBlogs []models.Blog

	for _, val := range fakeBlogDB {
		for _, t := range val.Tags {
			if t == tag {
				filtredBlogs = append(filtredBlogs, val)
			}
		}
	}
	log.Println(filtredBlogs)

	tmpl.Execute(ctx.Writer, filtredBlogs)
}
