package objects

type CellType int

const (
	EmptyCell CellType = iota
	LifeCell
	GhostCell
	SelfCell
)

// having coordinates also here makes it easier to make position adjustments without storing last position
type cell struct {
	x      int
	y      int
	object Object
}
