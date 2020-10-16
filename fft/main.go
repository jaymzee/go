package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func Twiddle(N int) complex128 {
	angle := 2 * math.Pi / float64(N)
	return cmplx.Exp(complex(0, angle))
}

func main() {
	fmt.Println(Twiddle(8))
}
