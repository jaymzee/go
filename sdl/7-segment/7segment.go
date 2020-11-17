package main

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

func Draw7segment(rend *sdl.Renderer, x, y int32, b uint8, color sdl.Color) {
	vpsave := rend.GetViewport()
	defer rend.SetViewport(&vpsave)
	rend.SetViewport(&sdl.Rect{x, y, 50, 75})
	dot := segment7dot
	bg := segment7bg
	vx := segment7vx
	vy := segment7vy
	outline := sdl.Rect{0, 0, 50, 75}

	// fill with background color (erase it)
	rend.SetDrawColor(bg.R, bg.G, bg.B, bg.A)
	rend.FillRect(&outline)

	// draw border
	gfx.RectangleColor(rend, 0, 0, outline.W, outline.H, segment7border)

	for i := 0; i < 8; i++ {
		led := segment7ledOff
		if (b>>i)&1 == 1 {
			led = color
		}
		if i < 7 {
			gfx.FilledPolygonColor(rend, vx[i][:], vy[i][:], led)
		} else {
			gfx.FilledCircleColor(rend, dot.X, dot.Y, segment7dotR, led)
		}
	}
}

// encode a decimal or hexadecimal digit for a seven-segment display
func Encode7segment(digit int, point bool) uint8 {
	var code uint8
	if digit >= 0 && digit < 16 {
		code = encode7segment[digit]
	}
	if point {
		code |= 0x80
	}
	return code
}

var (
	segment7outline = sdl.Rect{0, 0, 50, 75}
	segment7vx = [7][6]int16{
		[6]int16{18, 37, 40, 36, 19, 16}, // A x
		[6]int16{41, 39, 37, 34, 36, 39}, // B x
		[6]int16{38, 36, 34, 31, 33, 36}, // C x
		[6]int16{14, 30, 33, 30, 12, 10}, // D x
		[6]int16{8, 10, 12, 14, 12, 9},   // E x
		[6]int16{11, 13, 15, 18, 15, 12}, // F x
		[6]int16{17, 33, 37, 33, 15, 12}, // G x
	}
	segment7vy = [7][6]int16{
		[6]int16{9, 9, 12, 14, 14, 12},   // A y
		[6]int16{13, 33, 35, 32, 15, 12}, // B y
		[6]int16{38, 57, 59, 56, 40, 37}, // C y
		[6]int16{57, 57, 60, 62, 62, 60}, // D y
		[6]int16{58, 39, 37, 40, 56, 59}, // E y
		[6]int16{34, 14, 12, 15, 32, 35}, // F y
		[6]int16{33, 33, 36, 39, 39, 36}, // G y
	}
	segment7dot = sdl.Point{41, 60}
	segment7dotR = int32(3)
	segment7ledOff = sdl.Color{48, 48, 48, 255}
	segment7border = sdl.Color{255, 255, 255, 255}
	segment7bg = sdl.Color{32, 32, 32, 255}
)

var encode7segment = [16]uint8{
    0077,   // 0
    0006,   // 1
    0133,   // 2
    0117,   // 3
    0146,   // 4
    0155,   // 5
    0175,   // 6
    0007,   // 7
    0177,   // 8
    0157,   // 9
    0167,   // A
    0174,   // b
    0071,   // C
    0136,   // d
    0171,   // E
    0161,   // F
}
