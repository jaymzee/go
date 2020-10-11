package main

import (
	"fmt"
	"github.com/jaymzee/go/raspberrypi/gpio"
	"time"
)

func main() {
	btn := gpio.NewButton(19)
	for {
		fmt.Printf("%v\n", btn.Pressed())
		time.Sleep(250 * time.Millisecond)
	}
}
