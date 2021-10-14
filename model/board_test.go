package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	mapData := []string{
		" ###",
		"@$.$",
		"###?",
	}
	b := NewBoard(mapData)

	// board size
	assert.Equal(t, 4, b.Width)
	assert.Equal(t, 3, b.Height)

	// player location
	assert.Equal(t, 0, b.Player.X)
	assert.Equal(t, 1, b.Player.Y)

	// box locations
	assert.True(t, b.Get(1, 1).Box)
	assert.True(t, b.Get(3, 1).Box)

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
	mapData = []string{
		"+",
	}
	b = NewBoard(mapData)
	assert.Equal(t, 0, b.Player.X)
	assert.Equal(t, 0, b.Player.Y)
	assert.Equal(t, CellTypeGoal, b.Get(0, 0).TypeOf)

	// goals and boxes
	mapData = []string{
		".$*",
	}
	b = NewBoard(mapData)
	assert.Equal(t, CellTypeGoal, b.Get(0, 0).TypeOf)
	assert.Equal(t, CellTypeNone, b.Get(1, 0).TypeOf)
	assert.Equal(t, CellTypeGoal, b.Get(2, 0).TypeOf)
	assert.False(t, b.Get(0, 0).Box)
	assert.True(t, b.Get(1, 0).Box)
	assert.True(t, b.Get(2, 0).Box)
}
