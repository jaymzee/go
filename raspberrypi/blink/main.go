package main

import (
	"time"
	"github.com/jaymzee/go/raspberrypi/gpio"
)

func main() {
	led := gpio.NewLED(6)

	for i := 0; i < 4; i++ {
		led.Set(0)
		time.Sleep(500 * time.Millisecond)
		led.Set(1)
		time.Sleep(500 * time.Millisecond)
	}
}
