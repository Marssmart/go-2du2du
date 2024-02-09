package movement

import "fmt"

type Direction int

type Movement interface {
	Update(x int, y int, maxX int, maxY int) (bool, int, int, Direction)
	ToString() string
}

type Behavior interface {
	Next(lastAttemptedDir Direction) (error, Direction)
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

func (d Direction) Update(x int, y int, maxX int, maxY int) (bool, int, int, Direction) {
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
		possible, newX, newY, dir := MoveUp(x, y)
		if possible {
			return MoveLeft(newX, newY)
		}
		return false, -1, -1, dir
	case RightUpDiagonal:
		possible, newX, newY, dir := MoveUp(x, y)
		if possible {
			return MoveRight(newX, newY, maxX)
		}
		return false, -1, -1, dir
	case LeftDownDiagonal:
		possible, newX, newY, dir := MoveDown(x, y, maxY)
		if possible {
			return MoveLeft(newX, newY)
		}
		return false, -1, -1, dir
	case RightDownDiagonal:
		possible, newX, newY, dir := MoveDown(x, y, maxY)
		if possible {
			return MoveRight(newX, newY, maxX)
		}
		return false, -1, -1, dir
	case None:
		return MoveNone(x, y)
	default:
		panic(fmt.Sprintf("Unknown direction %v", d))
	}
}

func MoveRight(x int, y int, maxX int) (bool, int, int, Direction) {
	if x+1 > maxX {
		return false, -1, -1, Right
	}
	return true, x + 1, y, Right
}

func MoveLeft(x int, y int) (bool, int, int, Direction) {
	if x-1 < 0 {
		return false, -1, -1, Left
	}
	return true, x - 1, y, Left
}

func MoveDown(x int, y int, maxY int) (bool, int, int, Direction) {
	if y+1 > maxY {
		return false, -1, -1, Down
	}
	return true, x, y + 1, Down
}

func MoveUp(x int, y int) (bool, int, int, Direction) {
	if y-1 < 0 {
		return false, -1, -1, Up
	}
	return true, x, y - 1, Up
}

func MoveNone(x int, y int) (bool, int, int, Direction) {
	return true, x, y, None
}
