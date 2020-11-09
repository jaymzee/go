package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.SetOutput(os.Stderr)

	ch := make(chan string)

	go func() {
		for i := 1; true; i++ {
			var msg string
			if i%2 == 1 {
				msg = fmt.Sprintf("ping %d", i)
			} else {
				msg = fmt.Sprintf("pong %d", i)
			}
			log.Printf("send %s", msg)
			ch <- msg
		}
	}()

	for i := 0; i < 8; i++ {
		msg := <-ch
		log.Printf("recv %s", msg)
		time.Sleep(1000 * time.Millisecond)
	}
}
