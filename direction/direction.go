package direction

type Direction int

const (
	U Direction = iota
	D
	L
	R
)

func (d Direction) String() string {
	return [...]string{"U", "D", "L", "R"}[d]
}
