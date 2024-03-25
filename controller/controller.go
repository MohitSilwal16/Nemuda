package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/MohitSilwal16/Nemuda/models"
)

func Temp() {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}
}

func RenderInitPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/index.html")
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

	fmt.Printf("Form Data: %#v\n", user)

	AddUser(user)

	fmt.Fprint(w, "<p class='flex justify-center'>Register Successful<p>")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		panic(err)
	}

	if VerifyIdPass(user) {
		GetTweets(w, r)
	} else {
		tmpl := template.Must(template.New("t").Parse("<p class='flex justify-center'>Invalid Credentials<p>"))
		tmpl.Execute(w, nil)
	}
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
