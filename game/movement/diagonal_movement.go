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

func (d *diagonalMovementBehavior) Next(lastFailedDirection Direction) (error, Direction) {
	var newDirection Direction

	//chatty but fast
	switch d.direction {
	case LeftUpDiagonal:
		{
			switch lastFailedDirection {
			case Left:
				newDirection = RightUpDiagonal
			case Up:
				newDirection = LeftDownDiagonal
			default:
				panic("Not supposed to happen")
			}
		}
	case LeftDownDiagonal:
		{
			switch lastFailedDirection {
			case Left:
				newDirection = RightDownDiagonal
			case Down:
				newDirection = LeftUpDiagonal
			default:
				panic("Not supposed to happen")
			}
		}
	case RightUpDiagonal:
		{
			switch lastFailedDirection {
			case Right:
				newDirection = LeftUpDiagonal
			case Up:
				newDirection = RightDownDiagonal
			default:
				panic("Not supposed to happen")
			}
		}
	case RightDownDiagonal:
		{
			switch lastFailedDirection {
			case Right:
				newDirection = LeftDownDiagonal
			case Down:
				newDirection = RightUpDiagonal
			default:
				panic("Not supposed to happen")
			}
		}
	default:
		return fmt.Errorf("unsupported direction %v for diagonal movement", d.direction), Up
	}
	d.direction = newDirection
	return nil, newDirection
}

func (d *diagonalMovementBehavior) Current() Direction {
	return d.direction
}
