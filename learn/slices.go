// slices is a program that demonstrates slices in go
package main

import (
	"fmt"
)

func main() {
	var x []float64 // empty slice (actually points to a null pointer)
	y := []float64{1, 2, 3}
	z := make([]float64, 10, 100)

	y[1] = 3.14
	z[3] = 42.123
	print_slice("x", x)
	print_slice("y", y)
	print_slice("z", z)
	fmt.Println()

	fmt.Println("appending")
	s := append(z, 2.7, 5.3)
	print_slice("z", z)
	print_slice("s", s)
	fmt.Println()

	fmt.Println("slicing")
	print_slice("z", z)
	print_slice("z", z[:15])
	print_slice("z", z[:5])
	t := z[:20]
	print_slice("t", t)
	t[18] = 82.3
	print_slice("t", t)
	fmt.Println()

	fmt.Println("z, t, and s point to the same underlying array")
	fmt.Printf("z.data=%p\n", z)
	fmt.Printf("t.data=%p\n", t)
	fmt.Printf("s.data=%p\n", s)
	fmt.Println()

	fmt.Println("more slicing (be careful of aliasing)")
	u := z[2:8]
	z[3] = 88
	print_slice("u", u)
	fmt.Printf("u.data=%p\n", u)
	fmt.Println()
}

// print_slice prints a slice, it's length and it's capacity
func print_slice(name string, s []float64) {
	fmt.Printf("%v=%v, len=%d, cap=%d\n", name, s, len(s), cap(s))
}
