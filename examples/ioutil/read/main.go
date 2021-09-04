// as of Go 1.16 ioutil functionality is now provided by io or os and those
// implementation should be preferred

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(bytes))
}
