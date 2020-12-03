package main

import (
	"fmt"
	"github.com/jaymzee/go/sdl/ssd"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

const (
	// ScreenWidth of window
	ScreenWidth = 260
	// ScreenHeight of window
	ScreenHeight = 115
	// Hour24 enables 24 hour display mode
	Hour24 = true
)

// Scene contains the state for the scene
type Scene struct {
	digit  *ssd.Display
	colon  *ssd.Colon
	toggle bool
	ticker *time.Ticker
}

// Loop is the event loop for the scene
func (scene *Scene) Loop(window *sdl.Window, renderer *sdl.Renderer) {
	scene.init(window, renderer)
	for running := true; running; {
		// respond to events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit")
				running = false
			default:
				//fmt.Printf("%T %#v\n", event, event)
			}
		}

		// draw a single frame of the scene
		scene.draw(window, renderer)
		renderer.Present()
	}
}

// init initializes the scene
func (scene *Scene) init(window *sdl.Window, renderer *sdl.Renderer) {
	var err error
	scene.digit, err = ssd.OpenSSD("ssd10d75.json")
	if err != nil {
		panic(err)
	}
	scene.digit.Encoder[16] = 0 // want spc to encode as off
	scene.colon, err = ssd.OpenColon("ssdc10d75.json")
	if err != nil {
		panic(err)
	}
	window.SetTitle("clock")
	scene.ticker = time.NewTicker(500 * time.Millisecond)
}

// draw draws a single frame of the scene
func (scene *Scene) draw(window *sdl.Window, renderer *sdl.Renderer) {
	time := <-scene.ticker.C
	window.SetTitle(time.Format("01/02/06"))
	h := time.Hour()
	if !Hour24 && h > 12 {
		h -= 12
	}
	hour := fmt.Sprintf("%2d", h)
	min := fmt.Sprintf("%02d", time.Minute())

	digit := scene.digit
	digit.Draw(renderer, 190, 20, digit.Encode(int(min[1])-'0', false))
	digit.Draw(renderer, 140, 20, digit.Encode(int(min[0])-'0', false))
	digit.Draw(renderer, 70, 20, digit.Encode(int(hour[1])-'0', false))
	digit.Draw(renderer, 20, 20, digit.Encode(int(hour[0])-'0', false))

	scene.toggle = !scene.toggle
	scene.colon.Draw(renderer, 120, 20, scene.toggle)
}
