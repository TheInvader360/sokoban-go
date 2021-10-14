package model

type Board struct {
	Width, Height int
	Cells         [][]Cell
	Player        *Player
}

// NewBoard - Creates a board from a slice of strings (Player "@", Box "$", Goal ".", Wall "#", Goal+Player "+", Goal+Box "*", None " ")
func NewBoard(mapData []string) *Board {
	b := Board{}

	b.Width = len(mapData[0])
	b.Height = len(mapData)
	b.Cells = make([][]Cell, b.Height)

	for y := 0; y < b.Height; y++ {
		b.Cells[y] = make([]Cell, b.Width)
		for x := 0; x < b.Width; x++ {
			code := mapData[y][x]
			cell := Cell{}
			switch code {
			case '@':
				b.Player = NewPlayer(x, y)
			case '$':
				cell.Box = true
			case '.':
				cell.TypeOf = CellTypeGoal
			case '#':
				cell.TypeOf = CellTypeWall
			case '+':
				cell.TypeOf = CellTypeGoal
				b.Player = NewPlayer(x, y)
			case '*':
				cell.TypeOf = CellTypeGoal
				cell.Box = true
			}
			b.Cells[y][x] = cell
		}
	}

	return &b
}

// Get - Returns the cell at the given location
func (b *Board) Get(x, y int) *Cell {
	return &b.Cells[y][x]
}
