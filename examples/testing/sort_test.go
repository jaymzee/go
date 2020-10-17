package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func Example() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	sort.Sort(ByAge(people))
	fmt.Printf("by age:   %v\n", people)
	sort.Sort(ByName(people))
	fmt.Printf("by name:  %v\n", people)
	// Output:
	// by age:   [{Michael 17} {Jenny 26} {Bob 31} {John 42}]
	// by name:  [{Bob 31} {Jenny 26} {John 42} {Michael 17}]
}
