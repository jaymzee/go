package fft

import (
	"math"
	"math/cmplx"
)

func Twiddle(N int) complex128 {
	angle := 2 * math.Pi / float64(N)
	return cmplx.Exp(complex(0, angle))
}

func ReverseBits(x uint, w uint) uint {
	x = (x&0xaaaaaaaa)>>1 | (x&0x55555555)<<1
	x = (x&0xcccccccc)>>2 | (x&0x33333333)<<2
	x = (x&0xf0f0f0f0)>>4 | (x&0x0f0f0f0f)<<4
	x = (x&0xff00ff00)>>8 | (x&0x00ff00ff)<<8
	x = x>>16 | x<<16

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

func Fft(x []complex128) {
	N := len(x)
	log2N := uint(math.Log2(float64(N)))

	Shuffle(x, x)

	for s := uint(1); s <= log2N; s++ {
		m := 1 << s
		m_2 := m >> 1
		W_m := Twiddle(-m)
		for k := 0; k < N; k += m {
			W := complex(1, 0)
			for j := 0; j < m_2; j++ {
				t := x[k+j]
				u := W * x[k+j+m_2]
				x[k+j] = t + u
				x[k+j+m_2] = t - u
				W *= W_m
			}
		}
	}
}

func Ifft(x []complex128) {
	N := len(x)
	log2N := uint(math.Log2(float64(N)))

	Shuffle(x, x)

	for s := uint(1); s <= log2N; s++ {
		m := 1 << s
		m_2 := m >> 1
		W_m := Twiddle(m)
		for k := 0; k < N; k += m {
			W := complex(1, 0)
			for j := 0; j < m_2; j++ {
				t := x[k+j]
				u := W * x[k+j+m_2]
				x[k+j] = t + u
				x[k+j+m_2] = t - u
				W *= W_m
			}
		}
	}
	NN := complex(float64(N), 0)
	for n := 0; n < N; n++ {
		x[n] /= NN
	}
}
