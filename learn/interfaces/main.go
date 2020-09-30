package main

type Magnitude interface {
	Abs() float64
}

const PI = 3.141592654

func main() {
	var m Magnitude
	x := &Point{3, 4}
	m = x
	mag := m.Abs()
	m = &Point3{3, 4, 5}
	mag += m.Abs()
	m = Polar{2.0, PI / 2}
	mag += m.Abs()
}
