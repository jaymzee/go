package main

import "fmt"

// demonstrates embedding
type FancyLightSwitch struct {
	LightSwitch
	Color   string
	Plating string
}

// since LightSwitch is embedded
// not all of the methods need to be reimplemented

func (s *FancyLightSwitch) On() {
	fmt.Printf("fancy switch on current is less than %v A\n", s.LightSwitch)
}
