// idiomatic way of implementing set in go is to use map

package main

import "fmt"

type void struct{}

func main() {
	// map from string to boolean
	basket := make(map[string]bool)
	basket["apple"] = true
	basket["banana"] = true

	// test for membership
	fmt.Println(basket["apple"])
	fmt.Println(basket["potato"])

	// alternative (avoids memory allocation used by booleans)
	var member void
	animals := make(map[string]void)
	animals["monkey"] = member
	animals["bear"] = member
	// test for membership
	if _, ok := animals["monkey"]; ok {
		fmt.Println("monkey")
	}
	if _, ok := animals["cheetah"]; ok {
		fmt.Println("cheetah")
	}
}
