package main

import (
	"log"
	"os"
)

const logFlags = log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile

func sum(x int, y int) int {
	result := x + y
	log.Printf("sum(%d, %d) %d\n", x, y, result)
	return result
}

func main() {
	log.SetFlags(logFlags)
	log.SetOutput(os.Stdout)
	log.SetPrefix("log: ")

	log.Println("hello")
	z := 0
	z = sum(z, 5)
	z = sum(z, 8)
	log.Println("z =", z)
	log.Println("goodbye")
}
