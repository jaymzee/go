package main

// #cgo CFLAGS:
// #cgo LDFLAGS:
// #include <unistd.h>
import "C"
import (
	"fmt"
	"regexp"
)

func main() {
	tty := ttyname()
	fmt.Printf("%s\n", tty)

	pattern := regexp.MustCompile(`tty\d`)
	if pattern.MatchString(tty) {
		fmt.Println("console")
	} else {
		fmt.Println("not a console")
	}
}

func ttyname() string {
	return C.GoString(C.ttyname(0))
}
