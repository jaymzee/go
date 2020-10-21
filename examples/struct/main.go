package main

import "fmt"

type T struct {
	a, b int
	c    float64
	d    string
}

func main() {
	s := T{a: 3, b: 5, c: 3.14, d: "joe"}
	s.b = 42
	t := &T{a: 5, b: 6, c: 7, d: "george"}
	fmt.Printf("%#v\n", s)
	fmt.Printf("%#v\n", t)
}
