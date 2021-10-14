package model

type Model struct {
	Board *Board
}

// NewModel - Creates a model
func NewModel() *Model {
	m := Model{}

	mapData := []string{
		"WWWWWWWWW",
		"W    TTTW",
		"W W WTWTW",
		"W W  TTTW",
		"W RRR W W",
		"W RPR   W",
		"W RRRWW W",
		"W       W",
		"WWWWWWWWW",
	}

	m.Board = NewBoard(mapData)

	return &m
}
