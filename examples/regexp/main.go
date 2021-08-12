package main

import (
	"fmt"
	"os"
	"regexp"
)

var regex = regexp.MustCompile("abc|def")

func main() {
	if len(os.Args) < 2 {
		os.Exit(2)
	}
	s := os.Args[1]

	if regex.MatchString(s) {
		fmt.Printf("%q", regex.FindString(s))
		fmt.Printf(" %v\n", regex.FindStringIndex(s))
	}
}
