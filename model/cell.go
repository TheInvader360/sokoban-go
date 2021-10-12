package model

type typeOf int

const (
	none typeOf = iota
	target
	wall
)

type Cell struct {
	typeOf typeOf
}
