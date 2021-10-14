package model

type Model struct {
	Board *Board
}

// NewModel - Creates a model
func NewModel() *Model {
	m := Model{}

	// Player "@", Box "$", Goal ".", Wall "#", Goal+Player "+", Goal+Box "*", None " ")
	mapData := []string{
		"#########",
		"#    ..*#",
		"# # #.#.#",
		"# #  ...#",
		"# $$$ # #",
		"# $@$   #",
		"#  $$## #",
		"#       #",
		"#########",
	}

	m.Board = NewBoard(mapData)

	return &m
}
