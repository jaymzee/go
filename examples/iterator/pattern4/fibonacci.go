package main

// Fibonacci is a statefule iterator for the fibonacci sequence
type Fibonacci struct {
	a, b, n int
}

// NewFibonacci returns a new stateful iterator for the fibonacci sequence
func NewFibonacci(n int) *Fibonacci {
	return &Fibonacci{a: 0, b: 1, n: n}
}

// Value returns the current fibonacci number
func (it *Fibonacci) Value() int {
	return it.a
}

// Next advances to the next fibonacci number and returns false if the sequence is complete
func (it *Fibonacci) Next() bool {
	it.n--
	it.a, it.b = it.b, it.a+it.b
	return it.n >= 0
}
