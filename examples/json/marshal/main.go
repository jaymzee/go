package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func marshal(o interface{}) {
	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("marshal(%v: %T) = %s\n", o, o, b)
}

func main() {
	marshal(true)
	marshal(42)
	marshal(3.14)
	marshal("hello")
	marshal([]int{1, 2, 3})
	marshal(map[string]int{"apple": 1, "pear": 2})
	marshal(person{"george", 32})
}
