package ssd

import (
	"errors"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

// Colon holds the configuration for the colon display
type Colon struct {
	Size        sdl.Rect
	P1          sdl.Point `json:"point1"`
	P2          sdl.Point `json:"point2"`
	R           int32     `json:"pointRadius"`
	FaceColor   sdl.Color
	BorderColor sdl.Color
	OnColor     sdl.Color
	OffColor    sdl.Color
}

// Draw draws a seven-segment display on the screen
//  renderer is the renderer to draw on
//  x is the x coordinate of the upper left corner
//  y is the y coordinate of the upper left corner
func (d *Colon) Draw(rend *sdl.Renderer, x, y int32, on bool) error {
	// calculate new viewport
	vp := sdl.Rect{X: x, Y: y, W: d.Size.X, H: d.Size.Y}

	// fill shape with background color (erase it)
	fill := d.FaceColor
	if err := rend.SetDrawColor(fill.R, fill.G, fill.B, fill.A); err != nil {
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
	if !gfx.RectangleColor(rend, 0, 0, vp.W, vp.H, d.BorderColor) {
		return errors.New("draw seven-segment border failed")
	}

	// draw leds
	led := d.OffColor
	if on {
		led = d.OnColor
	}
	if !gfx.FilledCircleColor(rend, d.P1.X, d.P1.Y, d.R, led) {
		return errors.New("draw colon led failed")
	}
	if !gfx.FilledCircleColor(rend, d.P2.X, d.P2.Y, d.R, led) {
		return errors.New("draw colon led failed")
	}

	return nil
}
