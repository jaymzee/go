package main

import "fmt"

func iterate(data []int, cb func(int)) {
	for _, val := range data {
		cb(val)
	}
}

func fold(xs []int, f func(int, int) int, initial int) int {
	acc := initial
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}

// fold should be preferred because it avoids a potential runtime error
func reduce(xs []int, f func(int, int) int) int {
	if len(xs) < 1 {
		panic("reduce of empty slice")
	}
	var acc int
	for i, x := range xs {
		if i > 0 {
			acc = f(acc, x)
		} else {
			acc = x
		}
	}
	return acc
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	iterate(nums, func(v int) {
		sum += v
	})
	fmt.Println("sum with iterator", sum)

	fmt.Println("sum with fold", fold(nums, func(a, b int) int { return a + b }, 0))
	fmt.Println("product with fold", fold(nums, func(a, b int) int { return a * b }, 1))
	fmt.Println("product with reduce", reduce(nums, func(a, b int) int { return a * b }))
}
