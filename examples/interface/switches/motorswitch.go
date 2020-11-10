package main

import "fmt"

// receiver does not need to be a struct to implement an interface
// in this example MotorSwitch is just a string to hold the name
// of the motor

type MotorSwitch string

func (s MotorSwitch) On() {
	fmt.Printf("motor %s turned on\n", s)
}

func (s MotorSwitch) Off() {
	fmt.Printf("motor %s turned off\n", s)
}
