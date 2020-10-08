package main

import (
	"github.com/jaymzee/go/raspberrypi/gpio"
	"os"
	"os/signal"
	"syscall"
)

// CTRL-C handler
func setupCtrlCHandler(led *gpio.LED) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		led.Out(1) // turn off LED
		os.Exit(0)
	}()
}
