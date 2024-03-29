package main

// demonstrate cleaning up when CTRL-C is pressed

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	resource := 42
	setupCloseHandler(resource)
	for {
		log.Println("tick")
		time.Sleep(1 * time.Second)
	}
}

func setupCloseHandler(resource int) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("\nCTRL C pressed")
		log.Printf("cleanup resource %d\n", resource)
		os.Exit(0)
	}()
}
