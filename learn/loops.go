package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	x := []int{1, 2, 3, 4, 42}

	for i, v := range x {
		fmt.Println(i, v)
	}
}
