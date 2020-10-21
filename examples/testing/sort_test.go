package testing

import (
	"fmt"
	. "github.com/jaymzee/go/examples/person"
	"sort"
)

type PeopleByAge []Person

func (a PeopleByAge) Len() int           { return len(a) }
func (a PeopleByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PeopleByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type PeopleByName []Person

func (a PeopleByName) Len() int           { return len(a) }
func (a PeopleByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PeopleByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func Example() {
	people := []Person{
		{Name: "Bob", Age: 31},
		{Name: "John", Age: 42},
		{Name: "Michael", Age: 17},
		{Name: "Jenny", Age: 26},
	}
	sort.Sort(PeopleByAge(people))
	fmt.Printf("by age:   %v\n", people)
	sort.Sort(PeopleByName(people))
	fmt.Printf("by name:  %v\n", people)
	// Output:
	// by age:   [{Michael 17} {Jenny 26} {Bob 31} {John 42}]
	// by name:  [{Bob 31} {Jenny 26} {John 42} {Michael 17}]
}
