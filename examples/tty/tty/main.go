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
	fmt.Printf("%s\n", ttyname())
	fmt.Printf("console=%v\n", isaconsole())
}

func isaconsole() bool {
	pattern := regexp.MustCompile(`/dev/tty\d`)
	return pattern.MatchString(ttyname())
}

func ttyname() string {
	return C.GoString(C.ttyname(C.STDIN_FILENO))
}
