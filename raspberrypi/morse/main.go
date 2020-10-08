package main

import (
	"os"
	"log"
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
	if flag.NArg() < 1 {
		fmt.Println("Usage: morse [-p pin] message")
		os.Exit(1)
	}
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	msg = strings.Join(flag.Args(), " ")
	fmt.Printf("Sending morse code on gpio pin %d\n", *pinFlag)
	for {
		sendString(led, msg)
	}
}
