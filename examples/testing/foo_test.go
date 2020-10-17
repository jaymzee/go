package main

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

func BenchmarkRandInt(b *testing.B) {
	// initialization code here
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Rand()
	}
}

func ExampleHello() {
	fmt.Println("Hello")
	// Output: Hello
}
