package controller

import (
	"fmt"

	"github.com/TheInvader360/sokoban-go/direction"
	"github.com/TheInvader360/sokoban-go/model"
	"github.com/faiface/pixel/pixelgl"
)

type Controller struct {
	m  *model.Model
	lm *LevelManager
}

// NewController - Creates a controller
func NewController(m *model.Model) *Controller {
	c := Controller{
		m:  m,
		lm: NewLevelManager(false),
	}

	return &c
}

// StartNewGame - Starts a new game at level 1
func (c *Controller) StartNewGame() {
	c.lm.Reset()
	c.tryStartNextLevel()
}

// HandleInput - Handles user input as appropriate (game state dependent behaviour)
func (c *Controller) HandleInput(key pixelgl.Button) {
	switch c.m.State {
	case model.StatePlaying:
		switch key {
		case pixelgl.KeyUp:
			c.tryMovePlayer(direction.U)
		case pixelgl.KeyDown:
			c.tryMovePlayer(direction.D)
		case pixelgl.KeyLeft:
			c.tryMovePlayer(direction.L)
		case pixelgl.KeyRight:
			c.tryMovePlayer(direction.R)
		case pixelgl.KeyR:
			c.restartLevel()
		}
	case model.StateLevelComplete:
		if key == pixelgl.KeySpace {
			c.tryStartNextLevel()
		}
	case model.StateGameComplete:
		if key == pixelgl.KeySpace {
			c.StartNewGame()
		}
	}
}

// tryMovePlayer - Move player (and an adjacent box where appropriate) in the specified direction if possible. Check for board completion (and handle appropriately) if a box is moved
func (c *Controller) tryMovePlayer(dir direction.Direction) {
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
		if targetCell.HasBox {
			nextCell := c.m.Board.Get(nextX, nextY)
			if nextCell.TypeOf == model.CellTypeWall {
				fmt.Printf("%v: Box blocked (wall)\n", dir)
			} else if nextCell.HasBox {
				fmt.Printf("%v: Box blocked (box)\n", dir)
			} else {
				targetCell.HasBox = false
				nextCell.HasBox = true
				c.m.Board.Player.X = targetX
				c.m.Board.Player.Y = targetY
				fmt.Printf("%v: Player moved (push)\n", dir)
				if c.m.Board.IsComplete() {
					c.m.State = model.StateLevelComplete
					fmt.Print("*** Level complete! ***\n(space key to continue)\n")
				}
			}
		} else {
			c.m.Board.Player.X = targetX
			c.m.Board.Player.Y = targetY
			fmt.Printf("%v: Player moved (clear)\n", dir)
		}
	}
}

// tryStartNextLevel - Starts the next level if the current one isn't the last, else sets game state to game complete
func (c *Controller) tryStartNextLevel() {
	if c.lm.HasNextLevel() {
		c.lm.ProgressToNextLevel()
		l := c.lm.GetCurrentLevel()
		c.m.Board = model.NewBoard(l.MapData, l.Width, l.Height)
		c.m.State = model.StatePlaying
		fmt.Printf("Start level %d\n", c.lm.GetCurrentLevelNumber())
	} else {
		c.m.State = model.StateGameComplete
		fmt.Print("*** GAME COMPLETE! ***\n(space key to restart)\n")
	}
}

// restartLevel - Resets the game board to the current level's starting state
func (c *Controller) restartLevel() {
	l := c.lm.GetCurrentLevel()
	c.m.Board = model.NewBoard(l.MapData, l.Width, l.Height)
	c.m.State = model.StatePlaying
	fmt.Printf("Restart level %d\n", c.lm.GetCurrentLevelNumber())
}
