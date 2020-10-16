package main

import (
	"time"
	"github.com/jaymzee/gpio0"
)

func main() {
	led := gpio0.NewLED(6)

	for i := 0; i < 4; i++ {
		led.On()
		time.Sleep(500 * time.Millisecond)
		led.Off()
		time.Sleep(500 * time.Millisecond)
	}
}
