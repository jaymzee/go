package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done")
				return
			case t := <-ticker.C:
				fmt.Println("tick at:", t)
			}
		}
	}()

	time.Sleep(2600 * time.Millisecond)
	ticker.Stop()
	done <- true // prevents leaked goroutine
	fmt.Println("Ticker stopped")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("finished")
}
