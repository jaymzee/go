package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	now := time.Now()
	s := fmt.Sprintf("timestamp: %s\n", now.Format(time.RFC3339))
	err := ioutil.WriteFile("data.txt", []byte(s), 0644)
	if err != nil {
		panic(err)
	}
}
