package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MohitSilwal16/Nemuda/controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve static files from the 'static' directory
	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", controller.RenderInitPage).Methods("GET")
	r.HandleFunc("/register", controller.RenderRegsiterPage).Methods("GET")
	r.HandleFunc("/login", controller.RenderLoginPage).Methods("GET")
	r.HandleFunc("/home", controller.ServeHomePage).Methods("GET")

	r.HandleFunc("/register", controller.Register).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")

	r.HandleFunc("/tweets", controller.GetTweets).Methods("GET")
	r.HandleFunc("/tweets", controller.CreateTweet).Methods("POST")

	r.HandleFunc("/contacts/username", controller.SearchUser).Methods("POST")

	fmt.Println("Listening on 8080 port ...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
