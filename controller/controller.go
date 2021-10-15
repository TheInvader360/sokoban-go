package controller

import (
	"fmt"

	"github.com/TheInvader360/sokoban-go/direction"
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

func (c *Controller) StartNextLevel() {
	c.lm.NextLevel()
	l := c.lm.GetCurrentLevel()
	c.m.Board = model.NewBoard(l.MapData, l.Width, l.Height)
	fmt.Printf("Start level %d\n", c.lm.CurrentLevelNumber)
}

func (c *Controller) RestartLevel() {
	l := c.lm.GetCurrentLevel()
	c.m.Board = model.NewBoard(l.MapData, l.Width, l.Height)
	fmt.Printf("Restart level %d\n", c.lm.CurrentLevelNumber)
}

func (c *Controller) TryMovePlayer(dir direction.Direction) {
	targetX := c.m.Board.Player.X
	targetY := c.m.Board.Player.Y
	nextX := targetX
	nextY := targetY

	switch dir {
	case direction.U:
		targetY--
		nextY -= 2
	case direction.D:
		targetY++
		nextY += 2
	case direction.L:
		targetX--
		nextX -= 2
	case direction.R:
		targetX++
		nextX += 2
	}

	targetCell := c.m.Board.Get(targetX, targetY)

	if targetCell.TypeOf == model.CellTypeWall {
		fmt.Printf("%v: Player blocked (wall)\n", dir)
	} else {
		if targetCell.Box {
			nextCell := c.m.Board.Get(nextX, nextY)
			if nextCell.TypeOf == model.CellTypeWall {
				fmt.Printf("%v: Box blocked (wall)\n", dir)
			} else if nextCell.Box {
				fmt.Printf("%v: Box blocked (box)\n", dir)
			} else {
				targetCell.Box = false
				nextCell.Box = true
				c.m.Board.Player.X = targetX
				c.m.Board.Player.Y = targetY
				fmt.Printf("%v: Player moved (push)\n", dir)
				if c.m.Board.IsComplete() {
					c.StartNextLevel()
				}
			}
		} else {
			c.m.Board.Player.X = targetX
			c.m.Board.Player.Y = targetY
			fmt.Printf("%v: Player moved (clear)\n", dir)
		}
	}
}
