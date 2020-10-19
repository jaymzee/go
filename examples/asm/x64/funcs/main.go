package main

import "fmt"

func main() {
	var a, b int64
	fmt.Println("2 + 3 =", Addi64(2, 3))

	a, b = Muli64(2, 3)
	fmt.Println("2 * 3 =", a)
	a, b = Muli64(0x200000003, 0x300000004)
	fmt.Printf("0x200000003 * 0x300000004 = x%016x%016x\n", b, a)
	fmt.Printf("%v\n", Mulf64(3.14, 2.7))
}

