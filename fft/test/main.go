package main

import (
	"fmt"
	"github.com/jaymzee/go/fft"
)

func printArray(name string, s []complex128) {
	for i, num := range s {
		fmt.Printf("%s[%d] = %v\n", name, i, num);
	}
}

func main() {
	x := []complex128{1,2,3,4,3,2,1,0}
	X := fft.Fft(x)
	x_ := fft.Ifft(X)

	printArray("x", x)

	fmt.Println("X = fft(x)")
	printArray("X", X)

	fmt.Println("x = fft(X)")
	printArray("x", x_)

	fmt.Println(int(real(x_[1])))
}
