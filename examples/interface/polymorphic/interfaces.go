package main

import "time"

type Animal interface {
	CommonName() string
	Speak() string
}

type Pet interface {
	Animal
	BirthDate() time.Time
	Name() string
	Play()
}

type Cat interface {
	Pet
	Breed() string
}

type Dog interface {
	Pet
	Breed() string
}
