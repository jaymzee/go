package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(keys)
}
