package main

import (
	"fmt"
	"github.com/jaymzee/go/examples/person"
	"sort"
)

func main() {
	people := []person.Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Printf("unsorted: %v\n", people)

	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Printf("by age:   %v\n", people)

	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Printf("by name:  %v\n", people)
}
