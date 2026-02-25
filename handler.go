package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler() {
	http.HandleFunc("GET /about/", handleAbout)
	http.HandleFunc("GET /users/", handleUsers)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	aboutHTML, err := os.ReadFile("./html/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error 500"))
		return
	}
	w.Write(aboutHTML)

	fmt.Printf("Received request from %s\n", r.RemoteAddr)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userId := query.Get("userId")
	w.Write([]byte("Users " + userId))
	fmt.Printf("Received request from %s\n", r.RemoteAddr)
}
