package controller

import (
	"testing"

	"github.com/TheInvader360/sokoban-go/direction"
	"github.com/TheInvader360/sokoban-go/model"
	"github.com/stretchr/testify/assert"
)

func TestStartNextLevelAndRestartLevel(t *testing.T) {
	m := model.Model{}
	c := NewController(&m)
	assert.IsType(t, &Controller{}, c)

	// level 0
	assert.Equal(t, 0, c.lm.CurrentLevelNumber)
	assert.Equal(t, "", c.lm.GetCurrentLevel().MapData)

	// start next level - level 1, player at start position
	c.StartNextLevel()
	assert.Equal(t, 1, c.lm.CurrentLevelNumber)
	assert.Equal(t, "########@ $ .########", c.lm.GetCurrentLevel().MapData)
	assert.True(t, m.Board.Player.X == 1)

	// move player away from start position
	c.TryMovePlayer(direction.R)
	assert.False(t, m.Board.Player.X == 1)

	// restart level - level 1, player back at start position
	c.RestartLevel()
	assert.Equal(t, "########@ $ .########", c.lm.GetCurrentLevel().MapData)
	assert.True(t, m.Board.Player.X == 1)
}

func TestTryMovePlayer(t *testing.T) {
	mapData := "" +
		"####" +
		"#  #" +
		"# @#" +
		"####"
	b := model.NewBoard(mapData, 4, 4)
	m := model.Model{Board: b}
	c := Controller{m: &m}

	// start position
	assert.Equal(t, 2, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)

	// move up (first attempt succeeds, second attempt fails)
	c.TryMovePlayer(direction.U)
	assert.Equal(t, 2, b.Player.X)
	assert.Equal(t, 1, b.Player.Y)
	c.TryMovePlayer(direction.U)
	assert.Equal(t, 2, b.Player.X)
	assert.Equal(t, 1, b.Player.Y)

	// move left (first attempt succeeds, second attempt fails)
	c.TryMovePlayer(direction.L)
	assert.Equal(t, 1, b.Player.X)
	assert.Equal(t, 1, b.Player.Y)
	c.TryMovePlayer(direction.L)
	assert.Equal(t, 1, b.Player.X)
	assert.Equal(t, 1, b.Player.Y)

	// move down (first attempt succeeds, second attempt fails)
	c.TryMovePlayer(direction.D)
	assert.Equal(t, 1, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)
	c.TryMovePlayer(direction.D)
	assert.Equal(t, 1, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)

	// move right (first attempt succeeds, second attempt fails)
	c.TryMovePlayer(direction.R)
	assert.Equal(t, 2, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)
	c.TryMovePlayer(direction.R)
	assert.Equal(t, 2, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)

	mapData = "" +
		"#######" +
		"#.  ..#" +
		"#$@$ $#" +
		"#     #" +
		"#######"
	b = model.NewBoard(mapData, 7, 5)
	m = model.Model{Board: b}
	c = Controller{m: &m}

	// start position
	assert.Equal(t, 2, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)

	// try move left (fail: can't push box into wall)
	c.TryMovePlayer(direction.L)
	assert.Equal(t, 2, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)

	// try move right (success: box pushed to the right)
	c.TryMovePlayer(direction.R)
	assert.Equal(t, 3, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)
	assert.False(t, b.Get(3, 2).Box)
	assert.True(t, b.Get(4, 2).Box)

	// try move right (fail: can't push box into other box)
	c.TryMovePlayer(direction.R)
	assert.Equal(t, 3, b.Player.X)
	assert.Equal(t, 2, b.Player.Y)

	/*
		// solve
		c.TryMovePlayer(direction.D)
		c.TryMovePlayer(direction.R)
		c.TryMovePlayer(direction.U)
		c.TryMovePlayer(direction.D)
		c.TryMovePlayer(direction.R)
		c.TryMovePlayer(direction.U)
		c.TryMovePlayer(direction.D)
		c.TryMovePlayer(direction.L)
		c.TryMovePlayer(direction.L)
		c.TryMovePlayer(direction.L)
		c.TryMovePlayer(direction.L)
		c.TryMovePlayer(direction.U)
	*/
}
