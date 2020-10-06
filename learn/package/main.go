package main

import (
	"fmt"
	"jaymzee/learn/point"
)

func main() {
	p := point.Point2{X: 2, Y: 3}
	p.X++
	fmt.Printf("p = %v\n", p)
}
