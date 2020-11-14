package main

import "github.com/veandco/go-sdl2/sdl"

// Scene is the interface for initialization and then drawing each frame
type Scene interface {
	Init(window *sdl.Window, renderer *sdl.Renderer)
	Draw(window *sdl.Window, renderer *sdl.Renderer)
}
