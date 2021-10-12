package model

type Rock struct {
	X, Y int
}

// NewRock - create a Rock at the given location
func NewRock(x, y int) *Rock {
	return &Rock{X: x, Y: y}
}
