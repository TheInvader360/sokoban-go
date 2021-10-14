package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	mapData := []string{
		" WWW",
		"PRTR",
		"WWW?",
	}
	b := NewBoard(mapData)

	// board size
	assert.Equal(t, 4, b.Width)
	assert.Equal(t, 3, b.Height)

	// player start location
	assert.Equal(t, 0, b.Player.X)
	assert.Equal(t, 1, b.Player.Y)

	// rock start locations
	assert.Equal(t, 2, len(b.Rocks))
	assert.Equal(t, 1, b.Rocks[0].X)
	assert.Equal(t, 1, b.Rocks[0].Y)
	assert.Equal(t, 3, b.Rocks[1].X)
	assert.Equal(t, 1, b.Rocks[1].Y)

	// top left
	assert.Equal(t, CellTypeNone, b.Get(0, 0).TypeOf)

	// top right
	assert.Equal(t, CellTypeWall, b.Get(3, 0).TypeOf)

	// middle row
	assert.Equal(t, CellTypeNone, b.Get(0, 1).TypeOf)
	assert.Equal(t, CellTypeNone, b.Get(1, 1).TypeOf)
	assert.Equal(t, CellTypeTarget, b.Get(2, 1).TypeOf)
	assert.Equal(t, CellTypeNone, b.Get(3, 1).TypeOf)

	// bottom left
	assert.Equal(t, CellTypeWall, b.Get(0, 2).TypeOf)

	// bottom right
	assert.Equal(t, CellTypeNone, b.Get(3, 2).TypeOf)
}
