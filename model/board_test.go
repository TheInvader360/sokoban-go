package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	cellData := []string{
		" WWW",
		"PRTR",
		"WWW?",
	}

	// board size
	board := NewBoard(cellData)
	assert.Equal(t, 4, board.Width)
	assert.Equal(t, 3, board.Height)

	// player start location
	assert.Equal(t, 0, board.Player.X)
	assert.Equal(t, 1, board.Player.Y)

	// rock start locations
	assert.Equal(t, 2, len(board.Rocks))
	assert.Equal(t, 1, board.Rocks[0].X)
	assert.Equal(t, 1, board.Rocks[0].Y)
	assert.Equal(t, 3, board.Rocks[1].X)
	assert.Equal(t, 1, board.Rocks[1].Y)

	// top left
	cell := board.get(0, 0)
	assert.Equal(t, none, cell.typeOf)

	// top right
	cell = board.get(3, 0)
	assert.Equal(t, wall, cell.typeOf)

	// middle row
	cell = board.get(0, 1)
	assert.Equal(t, none, cell.typeOf)
	cell = board.get(1, 1)
	assert.Equal(t, none, cell.typeOf)
	cell = board.get(2, 1)
	assert.Equal(t, target, cell.typeOf)
	cell = board.get(3, 1)
	assert.Equal(t, none, cell.typeOf)

	// bottom left
	cell = board.get(0, 2)
	assert.Equal(t, wall, cell.typeOf)

	// bottom right
	cell = board.get(3, 2)
	assert.Equal(t, none, cell.typeOf)
}
