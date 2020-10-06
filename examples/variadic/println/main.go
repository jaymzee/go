package main

import "log"

func p(args ...interface{}) {
	log.Println(args...)
}

func main() {
	x := []string{"banana", "pear", "orange"}
	log.SetFlags(0)
	p("hello world!")
	p("apple", 5, 3.14, x)
}
