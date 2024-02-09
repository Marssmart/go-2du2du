package movement

import (
	"fmt"
)

func NewDiagonalMovementBehavior() Behavior {
	return &diagonalMovementBehavior{LeftUpDiagonal}
}

type diagonalMovementBehavior struct {
	direction Direction
}

func (d *diagonalMovementBehavior) Next() (error, Direction) {
	var newDirection Direction

	switch d.direction {
	case LeftUpDiagonal:
		newDirection = LeftDownDiagonal
	case LeftDownDiagonal:
		newDirection = RightDownDiagonal
	case RightUpDiagonal:
		newDirection = LeftUpDiagonal
	case RightDownDiagonal:
		newDirection = RightUpDiagonal
	default:
		return fmt.Errorf("unsupported direction %v for diagonal movement", d.direction), Up
	}
	d.direction = newDirection
	return nil, newDirection
}

func (d *diagonalMovementBehavior) Current() Direction {
	return d.direction
}
