// as of Go 1.16 ioutil functionality is now provided by io or os and those
// implementation should be preferred

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
