package main

import (
	"fmt"
	"jaymzee/learn/point"
)

func example1() {
	fmt.Println("methods for Point struct")
	x := &point.Point2{3, 4}
	y := point.Point2{3.141592654, 4.3}

	fmt.Printf("x = %.4v\n", *x)
	fmt.Printf("y = %.4v\n", y)
	x.Scale(5)
	fmt.Println("x.Scale(5)")
	y.Scale(3)
	fmt.Println("y.Scale(3)")
	fmt.Printf("x = %.4v\n", *x)
	fmt.Printf("y = %.4v\n", y)
	fmt.Printf("x.Abs() = %v\n", x.Abs())
}
