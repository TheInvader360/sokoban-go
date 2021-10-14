package model

type Model struct {
	Board *Board
}

// NewModel - Creates a model
func NewModel() *Model {
	m := Model{}

	// Player "@", Box "$", Goal ".", Wall "#", Goal+Player "+", Goal+Box "*", None " ")
	mapData := "" +
		"#########" +
		"#    ..*#" +
		"# # #.#.#" +
		"# #  ...#" +
		"# $$$ # #" +
		"# $@$   #" +
		"#  $$## #" +
		"#       #" +
		"#########"

	m.Board = NewBoard(mapData, 9, 9)

	return &m
}
