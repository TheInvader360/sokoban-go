package model

type state int

const (
	StatePlaying state = iota
	StateLevelComplete
	StateGameComplete
)

type Model struct {
	LM              *LevelManager
	Board           *Board
	State           state
	TickAccumulator int
}

// NewModel - Creates a model
func NewModel() *Model {
	m := Model{
		LM: NewLevelManager(false),
	}

	return &m
}

// Update - Updates the model's current state (called once per main game loop iteration)
func (m *Model) Update() {
	m.TickAccumulator++
	if m.TickAccumulator > 20 {
		m.TickAccumulator = 0
	}
}
