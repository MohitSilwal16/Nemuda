package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/MohitSilwal16/Nemuda/models"
	"github.com/MohitSilwal16/Nemuda/utils"
	"github.com/gorilla/mux"
)

func ShowFiles() {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}
}

func setCookie(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:  "sessionToken",
		Value: token,
	}

	http.SetCookie(w, cookie)
}

func getCookie(r *http.Request) string {
	cookie, err := r.Cookie("sessionToken")

	if err == http.ErrNoCookie {
		return ""
	} else if err != nil {
		panic(err)
	}
	return cookie.Value
}

func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/index.html")
}

func RenderInitPage(w http.ResponseWriter, r *http.Request) {
	sessionToken := getCookie(r)
	if sessionToken != "" && checkDuplicateToken(sessionToken) {
		RenderHomePage(w, r)
	} else {
		RenderLoginPage(w, r)
	}
}

func RenderRegsiterPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/register.html"))

	err := tmpl.Execute(w, nil)

	if err != nil {
		panic(err)
	}
}

func RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/login.html"))

	err := tmpl.Execute(w, nil)

	if err != nil {
		panic(err)
	}
}

func RenderHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/homepage.html", "views/blog.html"))

	err := tmpl.Execute(w, fakeBlogDB)

	if err != nil {
		panic(err)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		panic(err)
	}

	var data map[string]string
	page := "views/register.html"

	if user.Username == "" || user.Password == "" {
		data = map[string]string{"Data": "All fields're manadatory"}
	} else if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(user.Username) {
		data = map[string]string{"Data": "Username must be alphanumeric"}
	} else if !utils.IsPasswordInFormat(user.Password) {
		data = map[string]string{"Data": "Password: 8+ chars, lower & upper case, digit, symbol"}
	} else {
		page = "views/homepage.html"
		data = map[string]string{"Data": user.Username}

		// Generate session token
		sessionToken := utils.TokenGenerator()

		for checkDuplicateToken(sessionToken) {
			sessionToken = utils.TokenGenerator()
			fmt.Println("Duplicate Token")
		}
		user.Token = sessionToken

		fmt.Printf("Form Data: %#v\n", user)
		AddUser(user)
		setCookie(w, user.Token)
	}

	tmpl := template.Must(template.ParseFiles(page))
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		panic(err)
	}
	var data map[string]string
	page := "views/login.html"

	if user.Username == "" || user.Password == "" {
		data = map[string]string{"Data": "All fields're manadatory"}
	} else if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(user.Username) || !utils.IsPasswordInFormat(user.Password) {
		data = map[string]string{"Data": "Invalid Credentials"}
	} else if VerifyIdPass(user) {
		user.Token = UpdateTokenInDBAndReturn(user.Username)
		setCookie(w, user.Token)

		RenderHomePage(w, r)
		return
	} else {
		data = map[string]string{"Data": "Invalid Credentials"}
	}

	tmpl := template.Must(template.ParseFiles(page))
	err = tmpl.Execute(w, data)

	if err != nil {
		fmt.Println(err)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	DeleteTokenInDB(getCookie(r))
	setCookie(w, "")

	RenderLoginPage(w, r)
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Println(err)
		return
	}

	if SearchUserByName(user.Username) {
		tmpl := template.Must(template.New("t").Parse("Username is already used"))
		tmpl.Execute(w, nil)
	}
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

func GetBlogsByTags(w http.ResponseWriter, r *http.Request) {
	tag := mux.Vars(r)["tag"]
	fmt.Println(tag)

	tmpl := template.Must(template.ParseFiles("views/blog.html"))
	if tag == "All" {
		tmpl.Execute(w, fakeBlogDB)
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
	fmt.Println(filtredBlogs)

	tmpl.Execute(w, filtredBlogs)
}
