package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go/raspberrypi/gpio"
	"log"
	"os"
	"strings"
)

func main() {
	pinFlag := flag.Int("p", 6, "gpio pin number")
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Usage: morse [-p pin] message")
		os.Exit(1)
	}

	log.SetFlags(log.Ltime | log.Lmicroseconds)
	led := gpio.NewLED(*pinFlag)
	msg := strings.Join(flag.Args(), " ")
	setupCtrlCHandler(led)

	fmt.Printf("Sending morse code on gpio pin %d\n", *pinFlag)
	for {
		sendString(led, msg)
	}
}
