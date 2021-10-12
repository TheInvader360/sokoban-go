package model

type cellType int

const (
	CellTypeNone cellType = iota
	CellTypeTarget
	CellTypeWall
)

type Cell struct {
	TypeOf cellType
}
