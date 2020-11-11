package main

import (
	"fmt"
	"time"
)

type MyDog struct {
	birthdate time.Time
	breed string
	name string
}

func NewDog(name string, birthdate time.Time, breed string) Dog {
	return &MyDog{
		name: name,
		birthdate: birthdate,
		breed: breed,
	}
}

func (d *MyDog) CommonName() string {
	return "dog"
}

func (d *MyDog) Speak() string {
	return fmt.Sprintf("I am a %s %s, call me %s",
		d.Breed(), d.CommonName(), d.Name())
}

func (d *MyDog) BirthDate() time.Time {
	return d.birthdate
}

func (d *MyDog) Name() string {
	return d.name
}

func (d *MyDog) Play() {
	fmt.Println("play:", d.Speak())
}

func (d *MyDog) Breed() string {
	return d.breed
}

