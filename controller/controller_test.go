package controller

import (
	"testing"

	"github.com/TheInvader360/sokoban-go/model"
	"github.com/stretchr/testify/assert"
)

func TestSkipAndRestartLevel(t *testing.T) {
	m := model.Model{}
	c := NewController(&m)
	assert.IsType(t, &Controller{}, c)

	// level 0
	assert.Equal(t, 0, c.lm.CurrentLevelNumber)
	assert.Equal(t, "", c.lm.GetCurrentLevel().MapData)

	// skip level - level 1, player at start position
	c.SkipLevel()
	assert.Equal(t, 1, c.lm.CurrentLevelNumber)
	assert.Equal(t, "########@ $ .########", c.lm.GetCurrentLevel().MapData)
	assert.True(t, m.Board.Player.X == 1)

	// move player away from start position
	c.TryMovePlayerRight()
	assert.False(t, m.Board.Player.X == 1)

	// restart level - level 1, player back at start position
	c.RestartLevel()
	assert.Equal(t, "########@ $ .########", c.lm.GetCurrentLevel().MapData)
	assert.True(t, m.Board.Player.X == 1)
}

func TestTryMovePlayer(t *testing.T) {
	mapData := "" +
		"  " +
		" @"
	b := model.NewBoard(mapData, 2, 2)
	m := model.Model{Board: b}
	c := Controller{m: &m}

	// start position
	assert.Equal(t, 1, b.Player.X)
	assert.Equal(t, 1, b.Player.Y)

	// move up
	c.TryMovePlayerUp()
	assert.Equal(t, 1, b.Player.X)
	assert.Equal(t, 0, b.Player.Y)

	// move left
	c.TryMovePlayerLeft()
	assert.Equal(t, 0, b.Player.X)
	assert.Equal(t, 0, b.Player.Y)

	// move down
	c.TryMovePlayerDown()
	assert.Equal(t, 0, b.Player.X)
	assert.Equal(t, 1, b.Player.Y)

	// move right
	c.TryMovePlayerRight()
	assert.Equal(t, 1, b.Player.X)
	assert.Equal(t, 1, b.Player.Y)
}
