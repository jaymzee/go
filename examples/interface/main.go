package main

import . "github.com/jaymzee/go/examples/point"

type Magnitude interface {
	Abs() float64
}

const PI = 3.141592654

func main() {
	x := &Point2{X: 3, Y: 4}
	var m Magnitude = x
	mag := m.Abs()

	m = &Point3{X: 3, Y: 4, Z: 5}
	mag += m.Abs()

	m = Polar{R: 2.0, ϴ: PI / 2}
	mag += m.Abs()
}
