package objects

type CellType int

const (
	Enemy CellType = iota
	Life  CellType = iota
)

type board struct {
	cells [][]cell
}

type cell struct {
}
