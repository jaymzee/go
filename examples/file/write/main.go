package main

import (
	"os"
	"fmt"
	"time"
)

func main() {
	f, err := os.Create("data.txt")
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
