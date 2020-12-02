// Closures
//
// based on:
// https://ewencp.org/blog/golang-iterators/index.html
//
// author claims it's an ugly patter but it doesn't even work properly :/
package main

import "fmt"

func iterate(xs []int) (func() (int, bool), bool) {
	i := 0
	return func() (int, bool) {
		iPrev := i
		i++
		return xs[iPrev], (i < len(xs))
	}, (i < len(xs))
}

func main() {
	sum, val := 0, 0
	nums := []int{1, 2, 3, 4, 5}

	for it, more := iterate(nums); more; val, more = it() {
		sum += val
	}
	fmt.Println("sum", sum)
}
