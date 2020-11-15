package main

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS:
// #include <stdlib.h>
// #include "greeter.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// GreeterNumber demonstrates import of C #define constants
const GreeterNumber = C.GREETERNUMBER

// OtherNumber demonstrates import of C int
var OtherNumber = C.OtherNumber

// const int can also be imported like this but care must be taken to not
// modify the value in Go because it will result in a segmentation fault

func main() {
	name := C.CString("Gopher")
	defer C.free(unsafe.Pointer(name))

	year := C.int(2018)

	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))

	size := C.greet(name, year, (*C.char)(ptr))

	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))

	fmt.Printf("%v %T\n", GreeterNumber, GreeterNumber)
	C.OtherNumber++
	fmt.Printf("%v %T\n", C.OtherNumber, C.OtherNumber)
}
