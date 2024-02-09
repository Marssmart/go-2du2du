package movement

import "fmt"

type Direction int

type Movement interface {
	Update(x int, y int, maxX int, maxY int) (bool, int, int)
	ToString() string
}

type Behavior interface {
	Next() (error, Direction)
	Current() Direction
}

const (
	Up Direction = iota
	Down
	Left
	Right
	LeftUpDiagonal
	RightUpDiagonal
	LeftDownDiagonal
	RightDownDiagonal
	None
)

func (d Direction) ToString() string {
	switch d {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	case LeftUpDiagonal:
		return "LeftUpDiagonal"
	case RightUpDiagonal:
		return "RightUpDiagonal"
	case LeftDownDiagonal:
		return "LeftDownDiagonal"
	case RightDownDiagonal:
		return "RightDownDiagonal"
	case None:
		return "None"
	default:
		panic(fmt.Sprintf("Unknown Direction %v", d))
	}
}

func (d Direction) Update(x int, y int, maxX int, maxY int) (bool, int, int) {
	switch d {
	case Up:
		return MoveUp(x, y)
	case Down:
		return MoveDown(x, y, maxY)
	case Left:
		return MoveLeft(x, y)
	case Right:
		return MoveRight(x, y, maxX)
	case LeftUpDiagonal:
		possible, newX, newY := MoveUp(x, y)
		if possible {
			return MoveLeft(newX, newY)
		}
		return false, -1, -1
	case RightUpDiagonal:
		possible, newX, newY := MoveUp(x, y)
		if possible {
			return MoveRight(newX, newY, maxX)
		}
		return false, -1, -1
	case LeftDownDiagonal:
		possible, newX, newY := MoveDown(x, y, maxY)
		if possible {
			return MoveLeft(newX, newY)
		}
		return false, -1, -1
	case RightDownDiagonal:
		possible, newX, newY := MoveDown(x, y, maxY)
		if possible {
			return MoveRight(newX, newY, maxX)
		}
		return false, -1, -1
	case None:
		return MoveNone(x, y)
	default:
		return false, -1, -1
	}
}

func MoveRight(x int, y int, maxX int) (bool, int, int) {
	if x+1 > maxX {
		return false, -1, -1
	}
	return true, x + 1, y
}

func MoveLeft(x int, y int) (bool, int, int) {
	if x-1 < 0 {
		return false, -1, -1
	}
	return true, x - 1, y
}

func MoveDown(x int, y int, maxY int) (bool, int, int) {
	if y+1 > maxY {
		return false, -1, -1
	}
	return true, x, y + 1
}

func MoveUp(x int, y int) (bool, int, int) {
	if y-1 < 0 {
		return false, -1, -1
	}
	return true, x, y - 1
}

func MoveNone(x int, y int) (bool, int, int) {
	return true, x, y
}
