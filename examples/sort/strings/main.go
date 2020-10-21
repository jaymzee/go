package main

import (
	"fmt"
	"sort"
)

func fileString(xs []string, x string) []string {
	i := sort.SearchStrings(xs, x)
	ys := append(xs, "")
	copy(ys[i+1:], xs[i:])
	ys[i] = x
	return ys
}

func main() {
	fruit := []string{"pear", "apple", "banana"}
	fmt.Println(fruit, sort.StringsAreSorted(fruit))
	sort.Strings(fruit)
	fmt.Println(fruit, sort.StringsAreSorted(fruit))
	fruit = fileString(fruit, "peasoup")
	fruit = fileString(fruit, "bananrama")
	fruit = fileString(fruit, "aardvark")
	fmt.Println(fruit, sort.StringsAreSorted(fruit))
}
