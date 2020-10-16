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

func ReverseBits(x uint, w uint) uint {
	x = (x & 0xaaaaaaaa) >> 1 | (x & 0x55555555) << 1
	x = (x & 0xcccccccc) >> 2 | (x & 0x33333333) << 2
	x = (x & 0xf0f0f0f0) >> 4 | (x & 0x0f0f0f0f) << 4
	x = (x & 0xff00ff00) >> 8 | (x & 0x00ff00ff) << 8
	x = x >> 16 | x << 16

	return x >> (32 - w)
}


func main() {
	for i := uint(0); i < 16; i++ {
		fmt.Printf("%02x %02x\n", i, ReverseBits(i, 4))
	}
}
