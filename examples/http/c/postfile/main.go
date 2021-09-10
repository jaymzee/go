package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s url mimetype file\n", os.Args[0])
		os.Exit(2)
	}

	file, err := os.Open(os.Args[3])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	resp, err := http.Post(os.Args[1], os.Args[2], file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", body)
}
