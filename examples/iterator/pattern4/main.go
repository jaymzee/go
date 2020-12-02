// Stateful Iterators
// based on:
// https://ewencp.org/blog/golang-iterators/index.html

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	printSequence("nums", NewIntsIterator(nums))
	fmt.Println("sum:", foldInt(NewIntsIterator(nums), func(a, v int) int { return a + v }, 0))
	fmt.Println("product:", foldInt(NewIntsIterator(nums), func(a, v int) int { return a * v }, 1))
	printSequence("fibonacci", NewFibonacci(10))
}

type IntIterator interface {
	Value() int
	Next() bool
}

func printSequence(name string, it IntIterator) {
	first := true
	fmt.Printf("%s: [", name)
	for it.Next() {
		if first {
			first = false
		} else {
			fmt.Print(" ")
		}
		fmt.Print(it.Value())
	}
	fmt.Println("]")
}

func foldInt(it IntIterator, f func(int, int) int, initial int) int {
	acc := initial
	for it.Next() {
		acc = f(acc, it.Value())
	}
	return acc
}
