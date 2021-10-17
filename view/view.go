package view

import (
	"image"
	_ "image/png"
	"os"

	"github.com/TheInvader360/sokoban-go/model"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type spriteIndex int

const (
	SpritePlayer spriteIndex = iota
	SpriteBox
	SpriteGoal
	SpriteWall
	SpriteGoalAndPlayer
	SpriteGoalAndBox
)

type View struct {
	m       *model.Model
	win     *pixelgl.Window
	sprites []*pixel.Sprite
}

// NewView - Creates a view
func NewView(m *model.Model, win *pixelgl.Window) *View {
	file, err := os.Open("assets/spritesheet.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	image, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	pictureData := pixel.PictureDataFromImage(image)

	v := View{
		m:   m,
		win: win,
		sprites: []*pixel.Sprite{
			pixel.NewSprite(pictureData, pixel.R(float64(0), float64(0), float64(16), float64(16))),  // player
			pixel.NewSprite(pictureData, pixel.R(float64(16), float64(0), float64(32), float64(16))), // box
			pixel.NewSprite(pictureData, pixel.R(float64(32), float64(0), float64(48), float64(16))), // goal
			pixel.NewSprite(pictureData, pixel.R(float64(48), float64(0), float64(64), float64(16))), // wall
			pixel.NewSprite(pictureData, pixel.R(float64(64), float64(0), float64(80), float64(16))), // goal+player
			pixel.NewSprite(pictureData, pixel.R(float64(80), float64(0), float64(96), float64(16))), // goal+box
		},
	}

	return &v
}

// Draw - Draws a graphical representation of the model's current state (called once per main game loop iteration)
func (v *View) Draw() {
	switch v.m.State {
	case model.StatePlaying:
		v.win.Clear(colornames.Black)
	case model.StateLevelComplete:
		v.win.Clear(colornames.Blue)
	case model.StateGameComplete:
		v.win.Clear(colornames.Red)
	}

	v.drawSprite(SpritePlayer, float64(v.m.Board.Player.X), float64(v.m.Board.Player.Y))
	for y := 0; y < v.m.Board.Height; y++ {
		for x := 0; x < v.m.Board.Width; x++ {
			cell := v.m.Board.Get(x, y)
			switch cell.TypeOf {
			case model.CellTypeNone:
				if cell.HasBox {
					v.drawSprite(SpriteBox, float64(x), float64(y))
				}
			case model.CellTypeGoal:
				if cell.HasBox {
					v.drawSprite(SpriteGoalAndBox, float64(x), float64(y))
				} else if v.m.Board.Player.X == x && v.m.Board.Player.Y == y {
					v.drawSprite(SpriteGoalAndPlayer, float64(x), float64(y))
				} else {
					v.drawSprite(SpriteGoal, float64(x), float64(y))
				}
			case model.CellTypeWall:
				v.drawSprite(SpriteWall, float64(x), float64(y))
			}
		}
	}

	v.win.Update()
}

func (v *View) drawSprite(s spriteIndex, x, y float64) {
	r := pixel.R(x*16, v.win.Bounds().H()-(y+1)*16, (x+1)*16, v.win.Bounds().H()-(y)*16)
	v.sprites[s].Draw(v.win, pixel.IM.ScaledXY(pixel.ZV, pixel.V(r.W()/v.sprites[s].Frame().W(), r.H()/v.sprites[s].Frame().H())).Moved(r.Center()))
}
