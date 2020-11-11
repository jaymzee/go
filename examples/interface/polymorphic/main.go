package main

import (
	"fmt"
	"time"
)

func birthCertificate(p Pet) {
	fmt.Println("name:", p.Name())
	fmt.Println("birthdate:", p.BirthDate())
	fmt.Println("what:", p.Speak())
}

func poke(a Animal) {
	fmt.Println("poke:", a.Speak())
}

func main() {
	d1 := NewDalmation("dash", time.Now())
	d2 := NewDalmation("fido", time.Now())
	d3 := NewDog("snoopy", time.Now(), "beagle")
	s1 := NewSnake("python")
	s2 := NewSnake("corn")
	c1 := NewBurmese("socks", time.Now())
	poke(s1)
	poke(s2)
	poke(d1)
	poke(d2)
	poke(c1)
	birthCertificate(d1)
	birthCertificate(d2)
	birthCertificate(d3)
	birthCertificate(c1)
	c1.Play()
	d3.Play()
}
