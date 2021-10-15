package controller

import (
	"github.com/TheInvader360/sokoban-go/model"
)

type Controller struct {
	m  *model.Model
	lm *LevelManager
}

// NewController - Creates a controller
func NewController(m *model.Model) *Controller {
	c := Controller{
		m:  m,
		lm: NewLevelManager(),
	}

	return &c
}

func (c *Controller) SkipLevel() {
	c.lm.NextLevel()
	l := c.lm.GetCurrentLevel()
	c.m.Board = model.NewBoard(l.MapData, l.Width, l.Height)
}

func (c *Controller) RestartLevel() {
	l := c.lm.GetCurrentLevel()
	c.m.Board = model.NewBoard(l.MapData, l.Width, l.Height)
}

func (c *Controller) TryMovePlayerUp() {
	c.m.Board.Player.Y--
}

func (c *Controller) TryMovePlayerDown() {
	c.m.Board.Player.Y++
}

func (c *Controller) TryMovePlayerLeft() {
	c.m.Board.Player.X--
}

func (c *Controller) TryMovePlayerRight() {
	c.m.Board.Player.X++
}
