// Closures
//
// based on:
// https://ewencp.org/blog/golang-iterators/index.html
//
// author claims it's an ugly pattern but it doesn't even work properly :/
// fixed example with semantics similar to javascript and python
package main

import "fmt"

func iterate(xs []int) func() (int, bool) {
	i := 0
	return func() (int, bool) {
		if i < len(xs) {
			val := xs[i]
			i++
			return val, true
		}
		return 0, false
	}
}

func main() {
	sum := 0
	nums := []int{1, 2, 3, 4, 5}

	for next := iterate(nums); ; {
		if val, ok := next(); ok {
			sum += val
		} else {
			break
		}
	}
	fmt.Println("sum", sum)
}
