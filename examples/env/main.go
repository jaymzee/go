package main

import (
	"fmt"
	"os"
)

func main() {
	term := os.Getenv("TERM")
	fmt.Printf("TERM=%s\n", term)
}
