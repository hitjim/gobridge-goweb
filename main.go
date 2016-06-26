package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

func main() {
	fmt.Println("Heya")

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/markdown", markdown)
	http.Handle("/", http.FileServer(http.Dir("public")))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":"+port, nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	body := []byte("Hello there")
	w.Write(body)
}

func markdown(w http.ResponseWriter, r *http.Request) {
	body := []byte(r.FormValue("body"))
	markdown := blackfriday.MarkdownCommon(body)
	w.Write(markdown)
}
