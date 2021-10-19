package view

import (
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/TheInvader360/sokoban-go/model"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
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
	SpriteLogo
)

type View struct {
	m           *model.Model
	win         *pixelgl.Window
	scaleFactor float64
	text        *text.Text
	sprites     []*pixel.Sprite
}

// NewView - Creates a view
func NewView(m *model.Model, win *pixelgl.Window, scaleFactor float64) *View {
	fontFile, err := os.Open("assets/HackJack.ttf")
	if err != nil {
		panic(err)
	}
	defer fontFile.Close()
	fontBytes, err := ioutil.ReadAll(fontFile)
	if err != nil {
		panic(err)
	}
	font, err := truetype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size:              10 * scaleFactor,
		GlyphCacheEntries: 1,
	})
	atlas := text.NewAtlas(face, text.ASCII)
	text := text.New(pixel.V(0, 0), atlas)
	text.LineHeight = 11 * scaleFactor
	text.Color = colornames.White

	spritesheetFile, err := os.Open("assets/spritesheet.png")
	if err != nil {
		panic(err)
	}
	defer spritesheetFile.Close()
	image, _, err := image.Decode(spritesheetFile)
	if err != nil {
		panic(err)
	}
	pictureData := pixel.PictureDataFromImage(image)

	v := View{
		m:           m,
		win:         win,
		scaleFactor: scaleFactor,
		text:        text,
		sprites: []*pixel.Sprite{
			pixel.NewSprite(pictureData, pixel.R(float64(0), float64(0), float64(16), float64(16))),   // player
			pixel.NewSprite(pictureData, pixel.R(float64(16), float64(0), float64(32), float64(16))),  // box
			pixel.NewSprite(pictureData, pixel.R(float64(32), float64(0), float64(48), float64(16))),  // goal
			pixel.NewSprite(pictureData, pixel.R(float64(48), float64(0), float64(64), float64(16))),  // wall
			pixel.NewSprite(pictureData, pixel.R(float64(64), float64(0), float64(80), float64(16))),  // goal+player
			pixel.NewSprite(pictureData, pixel.R(float64(80), float64(0), float64(96), float64(16))),  // goal+box
			pixel.NewSprite(pictureData, pixel.R(float64(0), float64(16), float64(112), float64(64))), // logo
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

	v.drawLogoSprite()

	boardOffsetX := ((23 - v.m.Board.Width) / 2) + 1
	boardOffsetY := ((14 - v.m.Board.Height) / 2) + 1
	v.drawBoardSprite(SpritePlayer, float64(v.m.Board.Player.X), float64(v.m.Board.Player.Y), float64(boardOffsetX), float64(boardOffsetY))
	for y := 0; y < v.m.Board.Height; y++ {
		for x := 0; x < v.m.Board.Width; x++ {
			cell := v.m.Board.Get(x, y)
			switch cell.TypeOf {
			case model.CellTypeNone:
				if cell.HasBox {
					v.drawBoardSprite(SpriteBox, float64(x), float64(y), float64(boardOffsetX), float64(boardOffsetY))
				}
			case model.CellTypeGoal:
				if cell.HasBox {
					v.drawBoardSprite(SpriteGoalAndBox, float64(x), float64(y), float64(boardOffsetX), float64(boardOffsetY))
				} else if v.m.Board.Player.X == x && v.m.Board.Player.Y == y {
					v.drawBoardSprite(SpriteGoalAndPlayer, float64(x), float64(y), float64(boardOffsetX), float64(boardOffsetY))
				} else {
					v.drawBoardSprite(SpriteGoal, float64(x), float64(y), float64(boardOffsetX), float64(boardOffsetY))
				}
			case model.CellTypeWall:
				v.drawBoardSprite(SpriteWall, float64(x), float64(y), float64(boardOffsetX), float64(boardOffsetY))
			}
		}
	}

	//v.printString("A", 0, 0)
	//v.printString("Z", 63, 22)
	//v.printString(" !\"#$%&'()*+,-./\n0123456789\n:;<=>?@\nABCDEFGHIJKLMNOPQRSTUVWXYZ\n[\\]^_`\nabcdefghijklmnopqrstuvwxyz\n{|}~", 0, 0)
	v.printString(fmt.Sprintf("Level %d/%d", v.m.LM.GetCurrentLevelNumber(), v.m.LM.GetFinalLevelNumber()), 53, 22)

	v.win.Update()
}

func (v *View) drawLogoSprite() {
	r := pixel.R(float64(400)*v.scaleFactor, v.win.Bounds().H()-float64(48)*v.scaleFactor, float64(512)*v.scaleFactor, v.win.Bounds().H()-float64(0)*v.scaleFactor)
	v.sprites[SpriteLogo].Draw(v.win, pixel.IM.ScaledXY(pixel.ZV, pixel.V(r.W()/v.sprites[SpriteLogo].Frame().W(), r.H()/v.sprites[SpriteLogo].Frame().H())).Moved(r.Center()))
}

func (v *View) drawBoardSprite(s spriteIndex, x, y, offsetX, offsetY float64) {
	r := pixel.R((offsetX+x)*16*v.scaleFactor, v.win.Bounds().H()-(offsetY+y+1)*16*v.scaleFactor, (offsetX+x+1)*16*v.scaleFactor, v.win.Bounds().H()-(offsetY+y)*16*v.scaleFactor)
	v.sprites[s].Draw(v.win, pixel.IM.ScaledXY(pixel.ZV, pixel.V(r.W()/v.sprites[s].Frame().W(), r.H()/v.sprites[s].Frame().H())).Moved(r.Center()))
}

// printString - prints the given string at screen position x,y (i.e. 0-63,0-22)
func (v *View) printString(s string, x, y int) {
	v.text.Clear()
	v.text.WriteString(s)
	v.text.Draw(v.win, pixel.IM.Moved(pixel.V(float64(x*8)*v.scaleFactor, (245-float64(y*11))*v.scaleFactor)))
}
