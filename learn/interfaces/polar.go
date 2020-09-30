package main

type Polar struct {
	R float64
	ϴ float64
}

func (p Polar) Abs() float64 {
	return p.R
}
