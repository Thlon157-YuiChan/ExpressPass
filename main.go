package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	server := http.Server{
		Addr: ":8080", Handler: nil}

	http.HandleFunc("GET /about/", handleAbout)
	http.HandleFunc("GET /users/", handleUsers)

	server.ListenAndServe()
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
	w.Write([]byte("Users Page"))
	fmt.Printf("Received request from %s\n", r.RemoteAddr)
}
