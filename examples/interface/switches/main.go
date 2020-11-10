package main

func toggle(s Switch, n int) {
	for i := 0; i < n; i++ {
		s.On()
		s.Off()
	}
}

// what is cool about this is that all different receiver types can
// implement the same interface: structs, ints, strings, etc.

func main() {
	s1 := LightSwitch(15)
	s2 := MotorSwitch("5 hp induction")
	s3 := &ComplexSwitch{Name: "Widget", Param1: 42.0}
	s4 := &FancyLightSwitch{15, "white", "gold"}

	toggle(s1, 1)
	toggle(s2, 2)
	toggle(s3, 3)
	toggle(s4, 4)
}
