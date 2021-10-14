package model

type Board struct {
	Width, Height int
	Cells         [][]Cell
	Player        *Player
	Rocks         []*Rock
}

// NewBoard - Creates a board from a slice of strings ([P]layer start position / [R]ock start position / [T]arget / [W]all)
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
			case 'P':
				b.Player = NewPlayer(x, y)
			case 'R':
				b.Rocks = append(b.Rocks, NewRock(x, y))
			case 'T':
				cell.TypeOf = CellTypeTarget
			case 'W':
				cell.TypeOf = CellTypeWall
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
