package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	server := http.Server{
		Addr: ":8080", Handler: nil}

	http.HandleFunc("GET /about/", handleRoot)

	server.ListenAndServe()
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	rootHTML, err := os.ReadFile("./html/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error 500"))
		return
	}
	w.Write(rootHTML)

	fmt.Printf("Received request from %s\n", r.RemoteAddr)
}
