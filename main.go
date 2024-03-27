package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MohitSilwal16/Nemuda/controller"
	"github.com/MohitSilwal16/Nemuda/utils"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve static files from the 'static' directory
	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", controller.RenderInitPage).Methods("GET")

	r.HandleFunc("/register", controller.RenderRegsiterPage).Methods("GET")
	r.HandleFunc("/register", controller.Register).Methods("POST")

	r.HandleFunc("/login", controller.RenderLoginPage).Methods("GET")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/login", controller.Logout).Methods("DELETE")

	r.HandleFunc("/home", controller.ServeHomePage).Methods("GET")

	r.HandleFunc("/tweets", controller.GetTweets).Methods("GET")
	r.HandleFunc("/tweets", controller.CreateTweet).Methods("POST")

	r.HandleFunc("/contacts/username", controller.SearchUser).Methods("POST")

	fmt.Println("Listening on 8080 port ...")

	go func() {
		log.Fatal(http.ListenAndServe(":8080", r))
	}()

	var choi string

	for {
		fmt.Scanln(&choi)

		switch choi {
		case "h":
			fmt.Println("h - help")
			fmt.Println("c - clear")
			fmt.Println("q - quit")
		case "c":
			utils.ClearScreen()
		case "q":
			os.Exit(0)
		default:
			fmt.Println("Enter h for help")
		}
	}
}
