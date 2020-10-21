package main

import (
	"fmt"
	. "github.com/jaymzee/go/examples/point"
)

func main() {
	p := Point2{X: 2, Y: 3}
	p.X++
	fmt.Printf("p = %v\n", p)
}
