package view

import (
	"github.com/TheInvader360/sokoban-go/model"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type View struct {
	m      *model.Model
	imd    *imdraw.IMDraw
	width  int
	height int
}

// NewView - Creates a view
func NewView(m *model.Model, imd *imdraw.IMDraw, width, height int) *View {
	v := View{
		m:      m,
		imd:    imd,
		width:  width,
		height: height,
	}
	return &v
}

// Draw - Draws a graphical representation of the model's current state to IMDraw (called once per main game loop iteration)
func (v *View) Draw() {
	switch v.m.State {
	case model.StatePlaying:
		for y := 0; y < v.m.Board.Height; y++ {
			for x := 0; x < v.m.Board.Width; x++ {
				cell := v.m.Board.Get(x, y)
				switch cell.TypeOf {
				case model.CellTypeNone:
					v.imd.Color = colornames.Grey
				case model.CellTypeGoal:
					v.imd.Color = colornames.Limegreen
				case model.CellTypeWall:
					v.imd.Color = colornames.Purple
				}
				v.imd.Push(
					pixel.V(float64(x*16), float64(v.height-y*16-16)),
					pixel.V(float64(x*16+16), float64(v.height-y*16)),
				)
				v.imd.Rectangle(0)
				if cell.HasBox {
					v.imd.Color = colornames.Red
					v.imd.Push(
						pixel.V(float64(x*16+2), float64(v.height-y*16-16+2)),
						pixel.V(float64(x*16+16-2), float64(v.height-y*16-2)),
					)
					v.imd.Rectangle(0)
				}
			}
		}
		v.imd.Color = colornames.Blue
		v.imd.Push(
			pixel.V(float64(v.m.Board.Player.X*16+2), float64(v.height-v.m.Board.Player.Y*16-16+2)),
			pixel.V(float64(v.m.Board.Player.X*16+16-2), float64(v.height-v.m.Board.Player.Y*16-2)),
		)
		v.imd.Rectangle(0)
	case model.StateLevelComplete:
		v.imd.Color = colornames.Yellow
		v.imd.Push(
			pixel.V(float64(0), float64(v.height)),
			pixel.V(float64(v.m.Board.Width*16), float64(v.height-v.m.Board.Height*16)),
		)
		v.imd.Rectangle(0)
	case model.StateGameComplete:
		v.imd.Color = colornames.Cyan
		v.imd.Push(
			pixel.V(float64(0), float64(v.height)),
			pixel.V(float64(v.m.Board.Width*16), float64(v.height-v.m.Board.Height*16)),
		)
		v.imd.Rectangle(0)
	}
}
