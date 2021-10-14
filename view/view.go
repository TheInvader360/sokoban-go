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

// Draw - Draws a graphical representation of the model's current state to IMDraw
func (v *View) Draw() {
	for y := 0; y < v.m.Board.Height; y++ {
		for x := 0; x < v.m.Board.Width; x++ {
			switch v.m.Board.Get(x, y).TypeOf {
			case model.CellTypeNone:
				v.imd.Color = colornames.Grey
			case model.CellTypeTarget:
				v.imd.Color = colornames.Limegreen
			case model.CellTypeWall:
				v.imd.Color = colornames.Purple
			}
			v.imd.Push(
				pixel.V(float64(x*16), float64(v.height-y*16-16)),
				pixel.V(float64(x*16+16), float64(v.height-y*16)),
			)
			v.imd.Rectangle(0)
		}
	}
	for _, rock := range v.m.Board.Rocks {
		v.imd.Color = colornames.Red
		v.imd.Push(
			pixel.V(float64(rock.X*16), float64(v.height-rock.Y*16-16)),
			pixel.V(float64(rock.X*16+16), float64(v.height-rock.Y*16)),
		)
		v.imd.Rectangle(0)
	}
	v.imd.Color = colornames.Blue
	v.imd.Push(
		pixel.V(float64(v.m.Board.Player.X*16), float64(v.height-v.m.Board.Player.Y*16-16)),
		pixel.V(float64(v.m.Board.Player.X*16+16), float64(v.height-v.m.Board.Player.Y*16)),
	)
	v.imd.Rectangle(0)
}
