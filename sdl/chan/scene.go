package main

import (
	"fmt"
	"github.com/jaymzee/go/sdl/sevensegment"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"time"
)

const (
	// ScreenWidth of window
	ScreenWidth = 400
	// ScreenHeight of window
	ScreenHeight = 400
)

// Colors
var (
	Red    = sdl.Color{R: 255, G: 0, B: 0, A: 255}
	Green  = sdl.Color{R: 0, G: 255, B: 0, A: 255}
	Yellow = sdl.Color{R: 255, G: 255, B: 0, A: 255}
)

// Scene contains the state for the scene
type Scene struct {
	ch1    chan int
	ch2    chan int
	sans18 *ttf.Font
	ssd    *sevensegment.Display
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

func foo(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
		time.Sleep(1000 * time.Millisecond)
	}
}

func bar(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
		time.Sleep(16666 * time.Microsecond)
	}
}

// NewScene initializes the scene
func (scene *Scene) init(window *sdl.Window, renderer *sdl.Renderer) {
	if font, err := ttf.OpenFont("DejaVuSans.ttf", 18); err != nil {
		panic(err)
	} else {
		scene.sans18 = font
	}
	if ssd, err := sevensegment.Open("seg7.json"); err != nil {
		panic(err)
	} else {
		scene.ssd = ssd
	}

	window.SetTitle("chan")

	scene.ch1 = make(chan int)
	go foo(scene.ch1)
	scene.ch2 = make(chan int)
	go bar(scene.ch2)
}

// Draw draws a single frame of the scene
func (scene *Scene) draw(window *sdl.Window, renderer *sdl.Renderer) {
	select {
	case number := <-scene.ch1:
		for i := 0; i < 4; i++ {
			b := scene.ssd.Encode(number>>(4*i)&0xF, false)
			scene.ssd.Draw(renderer, int32(250-50*i), 100, b, Green)
		}
	case number := <-scene.ch2:
		for i := 0; i < 4; i++ {
			b := scene.ssd.Encode(number>>(4*i)&0xF, false)
			scene.ssd.Draw(renderer, int32(250-50*i), 200, b, Red)
		}
	}
}
