package main

import "fmt"

func main() {
	var a, b int64
	fmt.Println(addi64(2, 3))
	fmt.Println(muli64(2, 3))
	a, b = muli64(0x200000003, 0x300000004)
	fmt.Printf("0x%016x%016x\n", b, a)
}

