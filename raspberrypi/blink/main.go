package main

import (
	"time"
	"github.com/jaymzee/go/raspberrypi/gpio"
)

func main() {
	led := gpio.NewLED(6)

	for i := 0; i < 4; i++ {
		led.Out(0) // turn on LED
		time.Sleep(500 * time.Millisecond)
		led.Out(1) // turn off LED
		time.Sleep(500 * time.Millisecond)
	}
}
