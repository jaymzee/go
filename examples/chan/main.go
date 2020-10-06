package main

import (
	"os"
	"log"
	"fmt"
	"time"
)

var lg = log.New(os.Stdout, "", log.Ldate | log.Ltime | log.Lmicroseconds)

func main() {
	ch := make(chan string)

	go func() {
		for i := 1; true; i++ {
			var msg string
			if i % 2 == 1 {
				msg = fmt.Sprintf("ping %d", i)
			} else {
				msg = fmt.Sprintf("pong %d", i)
			}
			lg.Printf("send %s", msg)
			ch <- msg
		}
	}()

	for i := 0; i < 8; i++ {
		msg := <-ch
		lg.Printf("recv %s", msg)
		time.Sleep(1000 * time.Millisecond)
	}
}
