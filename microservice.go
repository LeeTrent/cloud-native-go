package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/LeeTrent/cloud-native-go/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)

	////////////////////////////////////////////////////
	// Used to
	// 1. Retrieve all books (GET)
	// 2. Create and persist a book (POST)
	http.HandleFunc("/api/books", api.BooksHandleFunc)
	////////////////////////////////////////////////////
	// Used to:
	// 1. Retrieve an individual book by ISBN (GET)
	// 2. Update a book by ISBN (PUT)
	// 3. Delete a book by ISBN (DELETE)
	http.HandleFunc("/api/books/", api.BooksHandleFunc)
	////////////////////////////////////////////////////

	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Cloud Native Go!")
}

func echo(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query()["message"][0]
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, msg)
}
