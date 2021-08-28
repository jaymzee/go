package main

import (
	"fmt"
	"log"
)

func main() {
	err := fmt.Errorf("something failed")
	log.Fatal(err)
	fmt.Println("ureachable code")
}
