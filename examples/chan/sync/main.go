package main

import (
	"fmt"
	"time"
)

func worker(done chan<- bool) {
	fmt.Print("working...")
	time.Sleep(1 * time.Second)
	fmt.Println("done")
	done <- true // send message
}

func main() {
	done := make(chan bool) // create channel
	go worker(done)         // pass channel to worker
	<-done                  // wait for receipt of message from worker
}
