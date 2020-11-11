package main

import (
	"fmt"
	"time"
)

type Dalmation struct {
	birthdate time.Time
	name string
}

func NewDalmation(name string, birthdate time.Time) Dog {
	return &Dalmation{name: name, birthdate: birthdate}
}

func (d *Dalmation) CommonName() string {
	return "dog"
}

func (d *Dalmation) Speak() string {
	return fmt.Sprintf("I am a %s %s, call me %s",
		d.Breed(), d.CommonName(), d.Name())
}

func (d *Dalmation) BirthDate() time.Time {
	return d.birthdate
}

func (d *Dalmation) Name() string {
	return d.name
}

func (d *Dalmation) Play() {
	fmt.Println("play:", d.Speak())
}

func (d *Dalmation) Breed() string {
	return "dalmation"
}

