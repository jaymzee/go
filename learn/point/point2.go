package point

import "math"

type Point2 struct {
	X, Y float64
}

func (p *Point2) Scale(s float64) {
	p.X *= s
	p.Y *= s
}

func (p *Point2) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
