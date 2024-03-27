package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/MohitSilwal16/Nemuda/models"
	"github.com/MohitSilwal16/Nemuda/utils"
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

func RenderInitPage(w http.ResponseWriter, r *http.Request) {
	sessionToken := getCookie(r)
	var page map[string]string
	if sessionToken != "" && checkDuplicateToken(sessionToken) {
		page = map[string]string{"Page": "home"}
	} else {
		page = map[string]string{"Page": "login"}
	}

	tmpl := template.Must(template.ParseFiles("views/index.html"))
	err := tmpl.Execute(w, page)

	if err != nil {
		fmt.Println(err)
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

func ServeHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/homepage.html"))

	err := tmpl.Execute(w, nil)

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

	tmpl := template.Must(template.ParseFiles("views/homepage.html"))
	err = tmpl.Execute(w, nil)

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

	if VerifyIdPass(user) {
		user.Token = UpdateTokenInDBAndReturn(user.Username)
		setCookie(w, user.Token)

		tmpl := template.Must(template.ParseFiles("views/homepage.html"))
		err = tmpl.Execute(w, nil)

		if err != nil {
			fmt.Println(err)
		}
	} else {
		data := map[string]string{"Data": "Invalid Credentials"}

		tmpl := template.Must(template.ParseFiles("views/login.html"))
		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Println(err)
		}
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
var tweetBox = map[string][]models.Tweet{
	"Tweets": {
		{Username: "Nimesh", Content: "Hello"},
		{Username: "Nimesh", Content: "Owner of Gadhvi Airlines"},
		{Username: "Konark", Content: "Front-end god"},
	},
}

func CreateTweet(w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet

	err := json.NewDecoder(r.Body).Decode(&tweet)

	if err != nil {
		fmt.Println(err)
		return
	}
	tweet.Username = "Nimesh"
	fmt.Printf("%#v\n", tweet)
	tweetBox["Tweets"] = append(tweetBox["Tweets"], tweet)

	htmlStr := fmt.Sprintf("<div class='flex justify-end'>%s - %s</div>", tweet.Username, tweet.Content)

	tmpl, err := template.New("t").Parse(htmlStr)

	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl.Execute(w, nil)
}

func GetTweets(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/homepage.html"))

	tmpl.Execute(w, tweetBox)
}
