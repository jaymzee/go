package main

import (
	"net/http"
)

func main() {
	http.Handle("/hello", http.HandlerFunc(helloHandler))
	http.ListenAndServe("localhost:8080", http.DefaultServeMux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
