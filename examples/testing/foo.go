package testing

import "fmt"

// Abs computes the absolute value
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Rand returns a random number
func Rand() int {
	return 42
}

// Hello prints Hello
func Hello() {
	fmt.Println("Hello")
}
