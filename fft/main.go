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

func Shuffle(out []complex128, in []complex128) {
	if len(in) != len(out) {
		panic("Shuffle: out and in must be the same length!")
	}
	N := uint(len(out))
	w := uint(math.Log2(float64(N)))
	if &out[0] != &in[0] {
		for k := uint(0); k < N; k++ {
			out[ReverseBits(k, w)] = in[k]
		}
	} else {
		for a := uint(0); a < N; a++ {
			b := ReverseBits(a, w)
			if a < b {
				out[a], out[b] = out[b], out[a]
			}
		}
	}
}


func main() {
	x := []complex128{1,2,3,4}
	y := make([]complex128, 4)

	Shuffle(y, x)
	fmt.Println(y)
}
