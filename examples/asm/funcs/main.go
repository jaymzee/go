package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("2 + 3 =", AddInt(2, 3))

	a, b = MultInt(2, 3)
	fmt.Println("2 * 3 =", a)
	a, b = MultInt(0x3, 0x4)
	fmt.Printf("0x3 * 0x4 = 0x%016x%016x\n", b, a)
	fmt.Printf("%v\n", MultFloat32(3.14, 2.7))
	fmt.Printf("%v\n", MultFloat64(3.14, 2.7))
}

