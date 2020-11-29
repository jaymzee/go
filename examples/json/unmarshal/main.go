package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data := `{"name":"george","age":42}`
	var p person
	json.Unmarshal([]byte(data), &p)
	fmt.Printf("%s -> %#v\n", data, p)
}
