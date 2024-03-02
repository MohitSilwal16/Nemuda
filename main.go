package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/MohitSilwal16/Nemuda/views"
	"github.com/gorilla/mux"
)

func serveRegister(w http.ResponseWriter, r *http.Request) {
	views.Register().Render(context.TODO(), w)
}

func register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	pass := r.PostFormValue("pass")

	fmt.Println(username, pass)

	fmt.Fprint(w, "<p class='flex justify-center'>Register Successful<p>")
}

func clearScreen() {
	var cmd *exec.Cmd

	// Check the operating system to determine the appropriate clear command
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear") // for Unix-like systems
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") // for Windows
	default:
		fmt.Println("Unsupported platform.")
		return
	}

	// Execute the clear command
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clearScreen()

	r := mux.NewRouter()

	// Serve static files from the 'static' directory
	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", serveRegister)
	r.HandleFunc("/register", register)

	fmt.Println("Listening on 8080 port ...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
