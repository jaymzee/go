// goodbye is a program
package main

import (
	"fmt"
)

// TZ is the offset in seconds for the timezone
type TZ int

const (
	// UTC is Universal Time
	UTC TZ = 0 * 60 * 60
	// EST is Eastern Standard Time
	EST TZ = -5 * 60 * 60
)

var weekend = []string{"Saturday", "Sunday"}
var timeZone = map[string]TZ{"UTC": UTC, "EST": EST}

// so uhm
func main() {
	fmt.Println("timeZone:", timeZone)
	fmt.Println("weekend:", weekend)
	fmt.Println("typeZone[\"EST\"]:", timeZone["EST"])
}
