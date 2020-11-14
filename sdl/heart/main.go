package main

import (
	"flag"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
)

// sFlag enables software rendering
var sFlag bool

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "options:\n")
		flag.PrintDefaults()
	}
	flag.BoolVar(&sFlag, "s", false, "use software rendering")
}

func main() {
	flag.Parse()

	// initialize SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// initialize font library
	if err := ttf.Init(); err != nil {
		panic(err)
	}

	// create window
	window, err := sdl.CreateWindow("initializing SDL...",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		ScreenWidth, ScreenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// create renderer
	var renderer *sdl.Renderer
	if sFlag {
		fmt.Println("using software rendering")
		renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	} else {
		fmt.Println("using hardware rendering")
		renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	}
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	scene := new(Scene)
	scene.Loop(window, renderer)
}
