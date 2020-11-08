package main

import "fmt"

func Print(it interface{}) {
	// type assertions are different than type conversions
	// they bring out the concrete type underlying the interface
	if ints, ok := it.([]int); ok {
		fmt.Printf("ints: %#v\n", ints)
	} else if floats, ok := it.([]float64); ok {
		fmt.Printf("floats: %#v\n", floats)
	} else {
		fmt.Printf("unknown type: %#v\n", it)
	}
}

func main() {
	ints := []int{1, 2, 3}
	floats := []float64{1.1, 2.7, 3.14}
	strings := []string{"apple", "orange"}

	fmt.Println("type assertion:")
	Print(ints)
	Print(floats)
	Print(strings)
	Print("Hi")
	Print(42)
}
