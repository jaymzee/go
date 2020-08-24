// arrays is a program that demonstrates arrays in go
package main

import (
	"fmt"
)

const arr_size int = 3

func main() {
	var x [arr_size]float64 // x is an array that is default initialized
	y := [5]int{1, 2, 3}

	x[1] = 3.14
	foo(x)
	fmt.Println("y =", y)
	fmt.Println("len(y) =", len(y))

	for i, el := range y {
		fmt.Println(i, el)
	}
}

func foo(arr [arr_size]float64) {
	fmt.Println("x =", arr)
	fmt.Println("len(x) =", len(arr))
}
