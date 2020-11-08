package main

import "fmt"

type Person struct {
	Name string
	Age int
}

type Child Person

type Pet struct {
	Name string
}

func main() {
	bob := Person{Name: "bob", Age: 15}
	babyBob := Child(bob)

	fmt.Printf("%#v\n", bob)
	fmt.Printf("%#v\n", babyBob)
}
