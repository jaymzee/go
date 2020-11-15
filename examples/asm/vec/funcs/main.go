package main

import "fmt"

func main() {
	var lo, hi int
	fmt.Println("2 + 3 =", AddInt(2, 3))

	lo, _ = MultInt(2, 3)
	fmt.Println("2 * 3 =", lo)
	lo, hi = MultInt(0x3, 0x4)
	fmt.Printf("0x3 * 0x4 = 0x%016x%016x\n", hi, lo)
	fmt.Printf("%v\n", MultFloat32(3.14, 2.7))
	fmt.Printf("%v\n", MultFloat64(3.14, 2.7))
}
