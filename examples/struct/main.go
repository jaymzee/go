package main

import "fmt"

type T struct {
	a, b int
	c    float64
	d    string
}

func main() {
	s := T{a:3, b:5, c:3.14, d:"joe"}
	s.b = 42
	fmt.Printf("%#v\n", s)
}
