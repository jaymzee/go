package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, World!\n")

	for i := 0; i < 16; i++ {
		fmt.Println(rand.Float64())
	}
	for i := 0; i < 16; i++ {
		fmt.Println(rand.Intn(100))
	}
}
