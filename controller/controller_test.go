package controller

import (
	"testing"

	"github.com/TheInvader360/sokoban-go/model"
	"github.com/stretchr/testify/assert"
)

func TestNewController(t *testing.T) {
	m := model.Model{}
	c := NewController(&m)
	assert.IsType(t, &Controller{}, c)
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
