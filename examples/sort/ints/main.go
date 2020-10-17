package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{1,5,2,3,7}
	sort.Ints(x)
	fmt.Println(x)
}
