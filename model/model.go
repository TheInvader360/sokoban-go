package model

type Model struct {
	Board *Board
}

// NewModel - Creates a model
func NewModel() *Model {
	m := Model{}

	return &m
}
