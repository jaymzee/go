package main

import (
	"log"
	"os"
)

const log2Flags = log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile

// log2 is a custom logger
var log2 = log.New(os.Stdout, "log2: ", log2Flags)

func sum(x int, y int) int {
	result := x + y
	log2.Printf("sum(%d, %d) %d\n", x, y, result)
	return result
}

func main() {
	log2.Println("hello")
	z := 0
	z = sum(z, 5)
	z = sum(z, 8)
	log2.Println("z =", z)
	log2.Println("goodbye")
}
