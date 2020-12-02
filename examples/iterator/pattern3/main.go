// Channels
// based on:
// https://ewencp.org/blog/golang-iterators/index.html
//
package main

import "fmt"

func iterate(xs []int) <-chan int {
	ch := make(chan int, 10)
	go func() {
		for _, x := range xs {
			ch <- x
		}
		close(ch) // Remember to close or the loop never ends!
	}()
	return ch
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0

	for v := range iterate(nums) {
		sum += v
	}
	fmt.Println("sum:", sum)

	prod := 1
	for v := range iterate(nums) {
		prod *= v
	}
	fmt.Println("product:", prod)
}
