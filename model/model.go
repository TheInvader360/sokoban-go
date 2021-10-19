package model

type state int

const (
	StatePlaying state = iota
	StateLevelComplete
	StateGameComplete
)

type Model struct {
	LM    *LevelManager
	Board *Board
	State state
}

// NewModel - Creates a model
func NewModel() *Model {
	m := Model{
		LM: NewLevelManager(false),
	}

	return &m
}
