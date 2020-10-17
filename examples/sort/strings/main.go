package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []string{"pear", "apple", "banana"}
	sort.Strings(x)
	fmt.Println(x)
}
