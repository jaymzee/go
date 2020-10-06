// slices demonstrates using slices in go
package main

import (
	"fmt"
)

func main() {
	var x []float64 // empty slice (actually points to a null pointer)
	y := []float64{1, 2, 3}
	z := make([]float64, 10, 100) // length=10, capacity=100
	y[1] = 3.14
	z[3] = 42.123
	printSlice("x", x)
	printSlice("y", y)
	printSlice("z", z)
	fmt.Println()

	fmt.Println("appending")
	s := append(z, 2.7, 5.3)
	printSlice("z", z)
	printSlice("s", s)
	fmt.Println()

	fmt.Println("slicing")
	printSlice("z", z)
	printSlice("z", z[:15])
	printSlice("z", z[:5])
	t := z[:20]
	printSlice("t", t)
	t[18] = 82.3
	printSlice("t", t)
	fmt.Println()

	fmt.Println("z, t, and s point to the same underlying array")
	fmt.Printf("z.data=%p\n", z)
	fmt.Printf("t.data=%p\n", t)
	fmt.Printf("s.data=%p\n", s)
	fmt.Println()

	fmt.Println("more slicing (be careful of aliasing)")
	u := z[2:8]
	z[3] = 88
	printSlice("u", u)
	fmt.Printf("u.data=%p\n", u)
	fmt.Println()
}

// printSlice prints a slice, it's length and it's capacity
func printSlice(name string, s []float64) {
	fmt.Printf("%s=%v, len=%d, cap=%d\n", name, s, len(s), cap(s))
}
