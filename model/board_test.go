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
	cell := board.Get(0, 0)
	assert.Equal(t, CellTypeNone, cell.TypeOf)

	// top right
	cell = board.Get(3, 0)
	assert.Equal(t, CellTypeWall, cell.TypeOf)

	// middle row
	cell = board.Get(0, 1)
	assert.Equal(t, CellTypeNone, cell.TypeOf)
	cell = board.Get(1, 1)
	assert.Equal(t, CellTypeNone, cell.TypeOf)
	cell = board.Get(2, 1)
	assert.Equal(t, CellTypeTarget, cell.TypeOf)
	cell = board.Get(3, 1)
	assert.Equal(t, CellTypeNone, cell.TypeOf)

	// bottom left
	cell = board.Get(0, 2)
	assert.Equal(t, CellTypeWall, cell.TypeOf)

	// bottom right
	cell = board.Get(3, 2)
	assert.Equal(t, CellTypeNone, cell.TypeOf)
}
