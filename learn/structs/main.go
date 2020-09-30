package main

import "fmt"

type T struct {
	a, b int
	c    float64
	d    string
}

func main() {
	var s T
	s.a = 3
	s.b = 5
	s.c = 3.14
	s.d = "joe"
	fmt.Printf("%v %T\n", s, s)
}
