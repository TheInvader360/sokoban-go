package controller

import (
	"github.com/TheInvader360/sokoban-go/model"
)

type Controller struct {
	m *model.Model
}

// NewController - Creates a controller
func NewController(m *model.Model) *Controller {
	c := Controller{
		m: m,
	}
	return &c
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
