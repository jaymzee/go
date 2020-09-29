//complex numbers are easy in go
package main

import "fmt"

func main() {
	var z complex64 = 1.1 + 2.2i
	fmt.Println("z   =", z)
	fmt.Println("z^2 =", z * z)
}
