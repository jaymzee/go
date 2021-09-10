package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.Handle("/hello", http.HandlerFunc(helloHandler))
	http.Handle("/form", http.HandlerFunc(formHandler))
	http.ListenAndServe("localhost:8080", http.DefaultServeMux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "hello, %v", r.PostForm)
}
