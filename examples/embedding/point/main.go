package main

import (
	"fmt"
	. "github.com/jaymzee/go/examples/point"
)

type ColoredPoint struct {
	Point2
	Color int
}

func main() {
	p1 := Point2{X: 2, Y: 3}
	p2 := ColoredPoint{Point2{3, 4}, 0xffffff}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p2.Abs())
}
