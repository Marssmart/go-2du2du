package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/constants"
	"image/color"
)

type Player interface {
	Icon() CachedImage
	Update(input Input)
	Draw(screen *ebiten.Image)
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
	p.Icon().Update(i)
	p.icon.Update(i)
}

func (p *player) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(color.White)
	screen.DrawImage(p.Icon().File(), p.Icon().Options())
}

func (p *player) Icon() CachedImage {
	return p.icon
}
