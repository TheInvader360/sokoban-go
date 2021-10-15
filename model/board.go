package model

type Board struct {
	Width, Height int
	Cells         []Cell
	Player        *Player
}

// NewBoard - Creates a board (map data encoding: Player "@", Box "$", Goal ".", Wall "#", Goal+Player "+", Goal+Box "*")
func NewBoard(mapData string, boardWidth, boardHeight int) *Board {
	b := Board{}

	b.Width = boardWidth
	b.Height = boardHeight

	b.Cells = make([]Cell, b.Width*b.Height)

	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			code := string(mapData[(y*b.Width)+x])
			cell := Cell{}
			switch code {
			case "@":
				b.Player = NewPlayer(x, y)
			case "$":
				cell.Box = true
			case ".":
				cell.TypeOf = CellTypeGoal
			case "#":
				cell.TypeOf = CellTypeWall
			case "+":
				cell.TypeOf = CellTypeGoal
				b.Player = NewPlayer(x, y)
			case "*":
				cell.TypeOf = CellTypeGoal
				cell.Box = true
			}
			b.Cells[(y*b.Width)+x] = cell
		}
	}

	return &b
}

// Get - Returns the cell at the given location
func (b *Board) Get(x, y int) *Cell {
	return &b.Cells[(y*b.Width)+x]
}

func (b *Board) IsComplete() bool {
	for _, cell := range b.Cells {
		if cell.TypeOf == CellTypeGoal && !cell.Box {
			return false
		}
	}
	return true
}
