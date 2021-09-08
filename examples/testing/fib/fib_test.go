package main

import "testing"

var fibTests = map[int]int {
	0: 0,
	1: 1,
	2: 1,
	3: 2,
	4: 3,
	5: 5,
	6: 8,
	7: 13,
	8: 21,
	9: 34,
	10: 55,
	12: 144,
	20: 6765,
}

func TestFib(t *testing.T) {
	for i, expected := range fibTests {
		got := Fib(i)
		if got != expected {
			t.Errorf("Fib(%d): got %d, expected %d", i, got, expected)
		}
	}
}

func BenchmarkFib10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}

func BenchmarkFib20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(20)
	}
}
