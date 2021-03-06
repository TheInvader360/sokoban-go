package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard(t *testing.T) {
	mapData := "" +
		" ###" +
		"@$.$" +
		"###?"
	b := NewBoard(mapData, 4, 3)

	// board size
	assert.Equal(t, 4, b.Width)
	assert.Equal(t, 3, b.Height)

	// player location
	assert.Equal(t, 0, b.Player.X)
	assert.Equal(t, 1, b.Player.Y)

	// box locations
	assert.True(t, b.Get(1, 1).HasBox)
	assert.True(t, b.Get(3, 1).HasBox)

	// top left
	assert.Equal(t, CellTypeNone, b.Get(0, 0).TypeOf)

	// top right
	assert.Equal(t, CellTypeWall, b.Get(3, 0).TypeOf)

	// middle row
	assert.Equal(t, CellTypeNone, b.Get(0, 1).TypeOf)
	assert.Equal(t, CellTypeNone, b.Get(1, 1).TypeOf)
	assert.Equal(t, CellTypeGoal, b.Get(2, 1).TypeOf)
	assert.Equal(t, CellTypeNone, b.Get(3, 1).TypeOf)

	// bottom left
	assert.Equal(t, CellTypeWall, b.Get(0, 2).TypeOf)

	// bottom right
	assert.Equal(t, CellTypeNone, b.Get(3, 2).TypeOf)

	// goal and player
	b = NewBoard("+", 1, 1)
	assert.Equal(t, 0, b.Player.X)
	assert.Equal(t, 0, b.Player.Y)
	assert.Equal(t, CellTypeGoal, b.Get(0, 0).TypeOf)

	// goals and boxes
	b = NewBoard(".$*", 3, 1)
	assert.Equal(t, CellTypeGoal, b.Get(0, 0).TypeOf)
	assert.Equal(t, CellTypeNone, b.Get(1, 0).TypeOf)
	assert.Equal(t, CellTypeGoal, b.Get(2, 0).TypeOf)
	assert.False(t, b.Get(0, 0).HasBox)
	assert.True(t, b.Get(1, 0).HasBox)
	assert.True(t, b.Get(2, 0).HasBox)
}

func TestIsComplete(t *testing.T) {
	mapData := "" +
		"#####" +
		"#@$.#" +
		"#####"
	b := NewBoard(mapData, 5, 3)
	assert.False(t, b.IsComplete())

	mapData = "" +
		"#####" +
		"# @*#" +
		"#####"
	b = NewBoard(mapData, 5, 3)
	assert.True(t, b.IsComplete())
}
