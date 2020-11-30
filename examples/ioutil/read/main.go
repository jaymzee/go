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
