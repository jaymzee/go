package main

import "fmt"

// receiver does not need to be a struct to implement an interface
// in this example LightSwitch is just an int to hold the current rating
// for the light switch

type LightSwitch int

func (s LightSwitch) On() {
	fmt.Printf("light turned on (current is less than %d amps)\n", s)
}

func (s LightSwitch) Off() {
	fmt.Printf("light turned off\n")
}
