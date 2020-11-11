package main

import (
	"fmt"
	"time"
)

type Burmese struct {
	birthdate time.Time
	name      string
}

func NewBurmese(name string, birthdate time.Time) Cat {
	return &Burmese{name: name, birthdate: birthdate}
}

func (d *Burmese) CommonName() string {
	return "cat"
}

func (d *Burmese) Speak() string {
	return fmt.Sprintf("I am a %s %s, call me %s",
		d.Breed(), d.CommonName(), d.Name())
}

func (d *Burmese) BirthDate() time.Time {
	return d.birthdate
}

func (d *Burmese) Name() string {
	return d.name
}

func (d *Burmese) Play() {
	fmt.Println("play:", d.Speak())
}

func (d *Burmese) Breed() string {
	return "Burmese"
}

