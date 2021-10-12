package model

type Board struct {
	Width, Height int
	Cells         [][]Cell
	Player        *Player
	Rocks         []*Rock
}

// NewBoard - create a board from a slice of strings ([P]layer start position / [R]ock start position / [T]arget / [W]all)
func NewBoard(cellData []string) *Board {
	board := Board{}

	board.Width = len(cellData[0])
	board.Height = len(cellData)
	board.Cells = make([][]Cell, board.Height)

	for y := 0; y < board.Height; y++ {
		board.Cells[y] = make([]Cell, board.Width)
		for x := 0; x < board.Width; x++ {
			code := cellData[y][x]
			cell := Cell{}
			switch code {
			case 'P':
				board.Player = NewPlayer(x, y)
			case 'R':
				board.Rocks = append(board.Rocks, NewRock(x, y))
			case 'T':
				cell.typeOf = target
			case 'W':
				cell.typeOf = wall
			}
			board.Cells[y][x] = cell
		}
	}

	return &board
}

func (b *Board) get(x, y int) Cell {
	return b.Cells[y][x]
}
