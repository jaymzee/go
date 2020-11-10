package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

const (
	Width, Height = 1000, 1000
	CenterX       = Width / 2
	CenterY       = Height / 2
	Radius        = 450
	FPS           = 500
	Points        = 500
)

func draw(r *sdl.Renderer, t, factor float64) {
	r.SetDrawColor(0, 0, 0, 255)
	r.Clear()
	r.SetDrawColor(0, 255, 0, 255)
	var x1, y1, x2, y2 float64
	for n := 0; n < Points; n++ {
		n1 := float64(n)
		n2 := factor * n1
		theta := 2 * math.Pi * n1 / Points
		phi := 2 * math.Pi * n2 / Points
		x1 = Radius * math.Cos(theta)
		y1 = Radius * math.Sin(theta)
		x2 = Radius * math.Cos(phi)
		y2 = Radius * math.Sin(phi)
		r.DrawLine(
			int32(CenterX-x1), int32(CenterY-y1),
			int32(CenterX-x2), int32(CenterY-y2),
		)
	}
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(Width, Height, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	window.SetTitle("Heart")

	t := 0.0
	factor := 1.0
	fps := float64(FPS)
	for running := true; running; {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
		draw(renderer, t, factor)
		renderer.Present()
		sdl.Delay(uint32(1000.0 / fps))
		t += 1.0 / fps
		factor += 0.001
	}
}
