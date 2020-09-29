package main

import (
	"log"
	"os"
)

// log2 is a custom logger
var log2 = log.New(os.Stdout, "log2: ", log.Ldate | log.Ltime |
										log.Lmicroseconds |
										log.Lshortfile)

func sum(x int, y int) int {
	result := x + y
	log2.Printf("sum(%d, %d) %d\n", x, y, result)
	return result
}

func main() {
	log.Println("hello world")
	z := 0
	z = sum(z, 5)
	z = sum(z, 8)
	log2.Println("z =", z)
	log2.Println("goodbye")
}
