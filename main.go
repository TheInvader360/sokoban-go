package main

import (
	"time"

	"github.com/TheInvader360/sokoban-go/controller"
	"github.com/TheInvader360/sokoban-go/model"
	"github.com/TheInvader360/sokoban-go/view"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	width       = 512
	height      = 256
	scaleFactor = 3
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Sokoban",
		Bounds: pixel.R(0, 0, width*scaleFactor, height*scaleFactor),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	m := model.NewModel()
	v := view.NewView(m, win, scaleFactor)
	c := controller.NewController(m)
	lastKey := pixelgl.KeyUnknown
	c.StartNewGame()

	// Main game loop
	for !win.Closed() {
		if win.Pressed(pixelgl.KeyEscape) {
			return
		}

		// Fire an event once per key press (no repeats if the key is held down)
		// Note: JustPressed() is a cleaner way to achieve this, but Pressed() more closely matches the Jack OS API
		if win.Pressed(pixelgl.KeyUp) {
			if lastKey != pixelgl.KeyUp {
				c.HandleInput(pixelgl.KeyUp)
			}
			lastKey = pixelgl.KeyUp
		} else if win.Pressed(pixelgl.KeyDown) {
			if lastKey != pixelgl.KeyDown {
				c.HandleInput(pixelgl.KeyDown)
			}
			lastKey = pixelgl.KeyDown
		} else if win.Pressed(pixelgl.KeyLeft) {
			if lastKey != pixelgl.KeyLeft {
				c.HandleInput(pixelgl.KeyLeft)
			}
			lastKey = pixelgl.KeyLeft
		} else if win.Pressed(pixelgl.KeyRight) {
			if lastKey != pixelgl.KeyRight {
				c.HandleInput(pixelgl.KeyRight)
			}
			lastKey = pixelgl.KeyRight
		} else if win.Pressed(pixelgl.KeyR) {
			if lastKey != pixelgl.KeyR {
				c.HandleInput(pixelgl.KeyR)
			}
			lastKey = pixelgl.KeyR
		} else if win.Pressed(pixelgl.KeySpace) {
			if lastKey != pixelgl.KeySpace {
				c.HandleInput(pixelgl.KeySpace)
			}
			lastKey = pixelgl.KeySpace
		} else {
			lastKey = pixelgl.KeyUnknown
		}

		v.Draw()

		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	pixelgl.Run(run)
}
