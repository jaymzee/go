package main

import (
	"fmt"
	"io"
	"os"
	"log"
)

func main() {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "Hello there\n")
		w.Close()
	}()

	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}
