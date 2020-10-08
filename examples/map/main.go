package main

import "fmt"

var lookup = map[string]int{
	"abc": 7,
	"fox": 5,
	"cbs": 2,
	"nbc": 4,
}

func main() {
	fmt.Println("nbc =", lookup["nbc"])
	// if key doesn't exist, returns zero value of type
	fmt.Println("pbs =", lookup["pbs"])

	if v, ok := lookup["abc"]; ok {
		// key exists
		fmt.Println("abc =", v)
	}

	if _, ok := lookup["pbs"]; !ok {
		// key does not exist
		fmt.Println("could not find pbs")
	}
}
