package main

import (
	"fmt"
	"net/http"

	"github.com/russross/blackfriday"
)

func main() {
	fmt.Println("Heya")

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/markdown", markdown)
	http.Handle("/", http.FileServer(http.Dir("public")))

	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
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
