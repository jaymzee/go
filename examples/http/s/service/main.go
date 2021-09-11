package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "hello world!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "hello, %v", r.PostForm)
}
