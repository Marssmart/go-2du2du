package objects

import (
	"go-2du2du/constants"
)

type Player interface {
	Icon() CachedImage
	Update(input Input)
}

type player struct {
	icon CachedImage
}

func NewPlayer() Player {
	path := constants.PlayerIconPath
	return &player{
		icon: NewImage(&path),
	}
}

func (p *player) Update(i Input) {
	p.icon.Update(i)
}

func (p *player) Icon() CachedImage {
	return p.icon
}
