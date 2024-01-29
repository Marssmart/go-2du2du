package objects

import (
	"go-2du2du/constants"
)

type Player interface {
	Lives() int
}

type player struct {
	lives int
}

func NewPlayer() Player {
	return &player{lives: constants.DefaultLives - 3}
}

func (p *player) Lives() int {
	return p.lives
}
