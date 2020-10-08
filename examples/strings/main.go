package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world"
	s2 := strings.ToUpper(s)
	s3 := "lo"
	slc := strings.Split(s, " ")
	fmt.Printf("Upper(%q) = %q\n", s, s2)
	fmt.Printf("Split(%q) = %q\n", s, slc)
	fmt.Printf("Contains(%q, %q) = %v\n", s, s3, strings.Contains(s, s3))
}
