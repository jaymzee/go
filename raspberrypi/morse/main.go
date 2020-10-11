package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go/raspberrypi/gpio"
	"log"
	"os"
	"os/signal"
	"strings"
)

func main() {
	pinFlag := flag.Int("p", -1, "gpio pin number")
	flag.Parse()
	if *pinFlag < 0 || flag.NArg() < 1 {
		fmt.Println("Usage: morse -p pin message")
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

// CTRL-C handler
func setupCtrlCHandler(led *gpio.LED) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		led.Off()
		os.Exit(0)
	}()
}
