package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		i   int
		i32 int32
		i64 int64
		y   []int // no data pointer is allocated
	)

	x := []int{}                     // data pointer = zero length array
	array := [...]int{1, 2, 3, 4, 5} // array
	slice := []int{1, 2, 3, 4, 5}    // slice
	lu := map[string]int{"abc": 7, "nbc": 4}
	s := "hello"

	fmt.Printf("    i: %T, %d bytes\n", i, unsafe.Sizeof(i))
	fmt.Printf("  i32: %T, %d bytes\n", i32, unsafe.Sizeof(i32))
	fmt.Printf("  i64: %T, %d bytes\n", i64, unsafe.Sizeof(i64))
	fmt.Printf("array: %#v, %d bytes\n", array, unsafe.Sizeof(array))
	fmt.Printf("slice: %#v, %d bytes\n", slice, unsafe.Sizeof(slice))
	fmt.Printf("    x: %#v, %d bytes\n", x, unsafe.Sizeof(y))
	fmt.Printf("    y: %#v, %d bytes\n", y, unsafe.Sizeof(y))
	fmt.Printf("   lu: %#v, %d bytes\n", lu, unsafe.Sizeof(lu))
	fmt.Printf("    s: %#v, %d bytes\n", s, unsafe.Sizeof(s))
}
