package model

type Player struct {
	X, Y int
}

// NewPlayer - Creates a Player at the given location
func NewPlayer(x, y int) *Player {
	return &Player{X: x, Y: y}
}
