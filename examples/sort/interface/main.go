package main

import (
	"fmt"
	. "github.com/jaymzee/go/examples/person"
	"sort"
)

type PeopleByAge []Person

func (p PeopleByAge) Len() int           { return len(p) }
func (p PeopleByAge) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PeopleByAge) Less(i, j int) bool { return p[i].Age < p[j].Age }

type PeopleByName []Person

func (p PeopleByName) Len() int           { return len(p) }
func (p PeopleByName) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PeopleByName) Less(i, j int) bool { return p[i].Name < p[j].Name }

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Printf("unsorted: %v\n", people)
	sort.Sort(PeopleByAge(people))
	fmt.Printf("by age:   %v\n", people)
	sort.Sort(PeopleByName(people))
	fmt.Printf("by name:  %v\n", people)
}
