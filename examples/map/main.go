package main

import "fmt"

var channel = map[string]int{
	"abc": 7,
	"fox": 5,
	"cbs": 2,
	"nbc": 4,
}

func printChannel(network string) {
	if ch, ok := channel[network]; ok {
		fmt.Printf("channel[%q] = %d\n", network, ch)
	} else {
		fmt.Printf("could not find a channel for %q\n", network)
	}
}

func main() {
	// if key isn't found in the map, a zero value of the type is returned
	fmt.Println("nbc =", channel["nbc"])
	fmt.Println("pbs =", channel["pbs"])

	// for handling this better (e.g. the zero value is used by valid entries)
	// there is a second boolean value returned by the map
	// that may be queried to see if the key exists
	printChannel("abc")
	printChannel("pbs")
}
