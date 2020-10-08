package main

import (
	"fmt"
	"strings"
	"flag"
	"github.com/jaymzee/go/raspberrypi/gpio"
)

func main() {
	const pin = 6
	msg := "Hello World"
	led := gpio.NewLED(pin)
	flag.Parse()
	if flag.NArg() > 0 {
		msg = strings.Join(flag.Args(), " ")
	}
	fmt.Printf("Sending %q as morse code on gpio pin %d\n", msg, pin)
	for {
		sendString(led, msg)
	}
}
