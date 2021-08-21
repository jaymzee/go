package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(str string) (int, int) {
	start := 0
	end := 0
	s := strings.Split(str, ":")
	if len(s) > 1 {
		start, _ = strconv.Atoi(s[0])
	}
	if len(s) > 0 {
		end, _ = strconv.Atoi(s[len(s)-1])
	}
	return start, end
}

func foo(str string) {
	s, e := parse(str)
	fmt.Printf("%#v = %d:%d\n", str, s, e)
}

func main() {
	foo("");
	foo(":");
	foo("0");
	foo("100");
	foo("0:100");
	foo(":100");
	foo("10:100");
	foo("20:");
}
