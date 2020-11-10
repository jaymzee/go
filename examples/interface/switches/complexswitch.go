package main

import "fmt"

// receiver does not need to be a struct
// but in this case it is

type ComplexSwitch struct {
	Name   string
	Param1 float64
	Param2 int
}

func (s *ComplexSwitch) On() {
	fmt.Printf("complex switch %s on, Param1=%f\n", s.Name, s.Param1)
}

func (s *ComplexSwitch) Off() {
	fmt.Printf("complex switch %s off\n", s.Name)
}
