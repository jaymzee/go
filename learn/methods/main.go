package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Scale(s float64) {
	p.X *= s
	p.Y *= s
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	x := &Point{3, 4}
	y := Point{3.141592654, 4.3}

	fmt.Printf("x = %.4v\n", *x)
	fmt.Printf("y = %.4v\n", y)
	x.Scale(5)
	y.Scale(3)
	fmt.Printf("x = %.4v\n", *x)
	fmt.Printf("y = %.4v\n", y)
	fmt.Printf("abs(x) = %v\n", x.Abs())
}
