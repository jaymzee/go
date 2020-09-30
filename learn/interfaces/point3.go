package main

import "math"

type Point3 struct{ X, Y, Z float64 }

func (p *Point3) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}
