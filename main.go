package main

import (
	"time"

	"github.com/TheInvader360/sokoban-go/controller"
	"github.com/TheInvader360/sokoban-go/direction"
	"github.com/TheInvader360/sokoban-go/model"
	"github.com/TheInvader360/sokoban-go/view"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	width  = 512
	height = 256
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Sokoban",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	m := model.NewModel()
	v := view.NewView(m, imd, width, height)
	c := controller.NewController(m)
	lastKey := pixelgl.KeyUnknown
	c.StartNextLevel()

	for !win.Closed() {
		if win.Pressed(pixelgl.KeyEscape) {
			return
		}

		// Fire an event once per key press (no repeats if the key is held down)
		// Note: JustPressed() is a cleaner way to achieve this, but Pressed() more closely matches the Jack OS API
		if win.Pressed(pixelgl.KeyUp) {
			if lastKey != pixelgl.KeyUp {
				c.TryMovePlayer(direction.U)
			}
			lastKey = pixelgl.KeyUp
		} else if win.Pressed(pixelgl.KeyDown) {
			if lastKey != pixelgl.KeyDown {
				c.TryMovePlayer(direction.D)
			}
			lastKey = pixelgl.KeyDown
		} else if win.Pressed(pixelgl.KeyLeft) {
			if lastKey != pixelgl.KeyLeft {
				c.TryMovePlayer(direction.L)
			}
			lastKey = pixelgl.KeyLeft
		} else if win.Pressed(pixelgl.KeyRight) {
			if lastKey != pixelgl.KeyRight {
				c.TryMovePlayer(direction.R)
			}
			lastKey = pixelgl.KeyRight
		} else if win.Pressed(pixelgl.KeyR) {
			if lastKey != pixelgl.KeyR {
				c.RestartLevel()
			}
			lastKey = pixelgl.KeyR
		} else if win.Pressed(pixelgl.KeyS) {
			if lastKey != pixelgl.KeyS {
				c.StartNextLevel()
			}
			lastKey = pixelgl.KeyS
		} else {
			lastKey = pixelgl.KeyUnknown
		}

		win.Clear(colornames.Grey)
		imd.Clear()
		v.Draw()
		imd.Draw(win)
		win.Update()

		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	pixelgl.Run(run)
}
