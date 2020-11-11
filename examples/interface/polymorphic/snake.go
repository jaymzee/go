package main

import "fmt"

type Snake struct {
	species string
}

func NewSnake(species string) Animal {
	return &Snake{species: species}
}

func (d *Snake) CommonName() string {
	return "snake"
}

func (d *Snake) Species() string {
	return d.species
}

func (d *Snake) Speak() string {
	return fmt.Sprintf("I am a %s %s", d.Species(), d.CommonName())
}
