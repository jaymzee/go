package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go/raspberrypi/gpio"
	"strings"
)

func main() {
	var msg string
	pinFlag := flag.Int("p", 6, "gpio pin number")
	flag.Parse()
	led := gpio.NewLED(*pinFlag)
	if flag.NArg() > 0 {
		msg = strings.Join(flag.Args(), " ")
		fmt.Printf("Sending morse code on gpio pin %d\n", *pinFlag)
		for {
			sendString(led, msg)
		}
	} else {
		fmt.Println("Usage: morse [-p pin] message")
	}
}
