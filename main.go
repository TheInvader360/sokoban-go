package main

import (
	"github.com/TheInvader360/sokoban-go/model"
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

	// initialise model
	cellData := []string{
		"WWWWWWWWW",
		"W    TTTW",
		"W W WTWTW",
		"W W  TTTW",
		"W RRR W W",
		"W RPR   W",
		"W RRRWW W",
		"W       W",
		"WWWWWWWWW",
	}
	board := model.NewBoard(cellData)

	// initialise imdraw
	imd := imdraw.New(nil)

	win.Clear(colornames.Black)

	for !win.Closed() {
		// draw the board
		imd.Clear()
		for y := 0; y < board.Height; y++ {
			for x := 0; x < board.Width; x++ {
				switch board.Get(x, y).TypeOf {
				case model.CellTypeNone:
					imd.Color = colornames.Grey
				case model.CellTypeTarget:
					imd.Color = colornames.Limegreen
				case model.CellTypeWall:
					imd.Color = colornames.Purple
				}
				imd.Push(
					pixel.V(float64(x*16), float64(height-y*16-16)),
					pixel.V(float64(x*16+16), float64(height-y*16)),
				)
				imd.Rectangle(0)
			}
		}
		for _, rock := range board.Rocks {
			imd.Color = colornames.Red
			imd.Push(
				pixel.V(float64(rock.X*16), float64(height-rock.Y*16-16)),
				pixel.V(float64(rock.X*16+16), float64(height-rock.Y*16)),
			)
			imd.Rectangle(0)
		}
		imd.Color = colornames.Blue
		imd.Push(
			pixel.V(float64(board.Player.X*16), float64(height-board.Player.Y*16-16)),
			pixel.V(float64(board.Player.X*16+16), float64(height-board.Player.Y*16)),
		)
		imd.Rectangle(0)
		imd.Draw(win)

		// update the window
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
