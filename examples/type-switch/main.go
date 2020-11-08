package main

import "fmt"

func Print(slice interface{}) {
	// type switch
	switch s := slice.(type) {
	case []int:
		fmt.Printf("s is a []int with len=%d, s[1] = %d\n", len(s), s[1])
	case []float64:
		fmt.Printf("s is a []float64 with len=%d, s[1] = %g\n", len(s), s[1])
	default:
		fmt.Printf("s is a %#v\n", s)
	}
}

func main() {
	ints := []int{1, 2, 3}
	floats := []float64{1.1, 2.7, 3.14}
	strings := []string{"apple", "orange"}

	fmt.Println("type switch:")
	Print(ints)
	Print(floats)
	Print(strings)
}
