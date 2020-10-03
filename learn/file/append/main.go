package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("data.txt",
		os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	now := time.Now()
	s := fmt.Sprintf("timestamp: %s\n", now.Format(time.RFC3339))
	_, err = f.WriteString(s)
	if err != nil {
		panic(err)
	}
}
