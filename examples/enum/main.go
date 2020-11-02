package main

import "fmt"

type Color int

const (
	Red Color = iota
	Blue
	Green
)

func (c Color) String() string {
	return [...]string{Red: "Red", Blue: "Blue", Green: "Green"}[c]
}

func main() {
	fmt.Println(Red)
	fmt.Println(Blue)
	fmt.Println(Green)
	fmt.Printf("%T\n", Green)
	fmt.Println("fav color is " + Green.String())
}
