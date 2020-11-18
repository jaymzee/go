package seg7

import (
	"errors"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

//Draw draws a seven-segment display on the screen
//  renderer is the renderer to draw on
//  x is the x coordinate of the upper left corner
//  y is the y coordinate of the upper left corner
//  b drives the A,B,C,D,E,F,G leds (bit-0..bit-6 -> A..G, bit-7 -> dot)
//  c is the color of the on state of the leds
func Draw(rend *sdl.Renderer, x, y int32, b uint8, c sdl.Color) error {
	// calculate new viewport
	vp := Default.Outline
	vp.X = x
	vp.Y = y

	// fill shape with background color (erase it)
	bg := Default.Background
	if err := rend.SetDrawColor(bg.R, bg.G, bg.B, bg.A); err != nil {
		return err
	}
	if err := rend.FillRect(&vp); err != nil {
		return err
	}

	// save the current viewport and set to the new one
	vpsave := rend.GetViewport()
	defer rend.SetViewport(&vpsave)
	if err := rend.SetViewport(&vp); err != nil {
		return err
	}

	// draw border
	if !gfx.RectangleColor(rend, 0, 0, vp.W, vp.H, Default.Border) {
		return errors.New("draw seven-segment border failed")
	}

	// draw leds
	d := &Default
	for i := 0; i < 8; i++ {
		led := d.LedOff
		if (b>>i)&1 == 1 {
			led = c
		}
		if i < 7 {
			if !gfx.FilledPolygonColor(rend, d.X[i][:], d.Y[i][:], led) {
				return errors.New("draw seven-segment led failed")
			}
		} else {
			if !gfx.FilledCircleColor(rend, d.Dot.X, d.Dot.Y, d.DotR, led) {
				return errors.New("draw seven-segment led failed")
			}
		}
	}

	return nil
}

// Encode encodes a digit for a seven-segment display
// the digit may be hexadecimal, decimal, or octal
func Encode(digit int, point bool) uint8 {
	code := Default.Encode[16]
	if digit >= 0 && digit < 16 {
		code = Default.Encode[digit]
	}
	if point {
		code |= 0x80
	}
	return code
}

// Config is the type for configuring seven-segment display settings
type Config = struct {
	Encode     [17]uint8
	Outline    sdl.Rect
	X          [7][6]int16
	Y          [7][6]int16
	Dot        sdl.Point
	DotR       int32
	Background sdl.Color
	Border     sdl.Color
	LedOff     sdl.Color
}

// Default holds the default settings for drawing a seven-segment display
var Default = Config{
	Encode: [17]uint8{
		0077, 0006, 0133, 0117, // 0, 1, 2, 3
		0146, 0155, 0175, 0007, // 4, 5, 6, 7
		0177, 0157, 0167, 0174, // 8, 9, A, b
		0071, 0136, 0171, 0161, // C, d, E, F
		0120, // r (error)
	},
	Outline: sdl.Rect{X: 0, Y: 0, W: 50, H: 75},
	X: [7][6]int16{
		[6]int16{18, 37, 40, 36, 19, 16}, // A x
		[6]int16{41, 39, 37, 34, 36, 39}, // B x
		[6]int16{38, 36, 34, 31, 33, 36}, // C x
		[6]int16{14, 30, 33, 30, 12, 10}, // D x
		[6]int16{8, 10, 12, 14, 12, 9},   // E x
		[6]int16{11, 13, 15, 18, 15, 12}, // F x
		[6]int16{17, 33, 37, 33, 15, 12}, // G x
	},
	Y: [7][6]int16{
		[6]int16{9, 9, 12, 14, 14, 12},   // A y
		[6]int16{13, 33, 35, 32, 15, 12}, // B y
		[6]int16{38, 57, 59, 56, 40, 37}, // C y
		[6]int16{57, 57, 60, 62, 62, 60}, // D y
		[6]int16{58, 39, 37, 40, 56, 59}, // E y
		[6]int16{34, 14, 12, 15, 32, 35}, // F y
		[6]int16{33, 33, 36, 39, 39, 36}, // G y
	},
	Dot:        sdl.Point{X: 41, Y: 60},
	DotR:       3,
	Background: sdl.Color{R: 0x20, G: 0x20, B: 0x20, A: 0xFF},
	Border:     sdl.Color{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
	LedOff:     sdl.Color{R: 0x28, G: 0x28, B: 0x28, A: 0xFF},
}
