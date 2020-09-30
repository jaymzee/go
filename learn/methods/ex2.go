package main

import "fmt"

type TZ int

const (
	HOUR TZ = 60 * 60
	UTC  TZ = 0 * HOUR
	EST  TZ = -5 * HOUR
)

var timeZones = map[string]TZ{
	"UTC": UTC,
	"EST": EST,
}

func (tz TZ) String() string {
	// scan through timeZones map to see if value = tz
	for name, zone := range timeZones {
		if tz == zone {
			return name
		}
	}
	// construct a string representation of minutes/seconds offset
	return fmt.Sprintf("%+d:%02d", tz/3600, (tz%3600)/60)
}

func example2() {
	fmt.Println("methods on int")
	// Print knows about method String()
	fmt.Println(EST)
	fmt.Println(5*HOUR/2)
}
