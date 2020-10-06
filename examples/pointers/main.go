package main

import "fmt"

func inc(x *int) {
	*x++
}

func main() {
	x := 42
	y := new(int)

	*y = 88
	inc(&x)
	inc(y)
	fmt.Printf("x = %d\n", x)
	fmt.Printf("y = %d\n", *y)
}
