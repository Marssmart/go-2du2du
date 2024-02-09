package movement

import (
	"fmt"
	"math/rand"
)

func NewCompassMovementBehavior() Behavior {
	return &compassMovementBehavior{Up}
}

type compassMovementBehavior struct {
	direction Direction
}

func (c *compassMovementBehavior) Next(lastAttempted Direction) (error, Direction) {
	newDirection := None

	switch c.Current() {
	case Up:
		newDirection = Left
	case Down:
		newDirection = Right
	case Left:
		newDirection = Down
	case Right:
		newDirection = Up
	default:
		return fmt.Errorf("unsupported move %v", c.Current()), None
	}
	c.direction = newDirection
	return nil, newDirection
}

func (c *compassMovementBehavior) Current() Direction {
	//20% chance for random movement other than last attempted if last move was successful
	if rand.Int31n(5) == 1 {
		//first 4 movements are iota 0,1,2,3, so we can just do this
		newDirection := Direction(rand.Int31n(4))

		for newDirection == c.Current() {
			newDirection = Direction(rand.Int31n(4))
		}

		c.direction = newDirection
		return newDirection
	}

	return c.direction
}
