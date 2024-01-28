package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/constants"
)

type Player interface {
	Icon() CachedImage
	Update(input Input)
	Draw(screen *ebiten.Image)
	Lives() int
}

type player struct {
	icon     CachedImage
	lives    int
	position PlayerPosition
}

func NewPlayer() Player {
	path := constants.PlayerIconPath
	image := NewImage(&path)
	image.PreLoadImage()
	bounds := image.File().Bounds()
	paddingWidth := float64(constants.PaddingWidth)
	paddingHeight := float64(constants.PaddingHeight)
	var widthBoundary = float64(constants.PlayerWidthBoundary)
	var heightBoundary = float64(constants.PlayerHeightBoundary)
	x, y := Center(&bounds)

	xCoordinate := boundaryCoordinate{
		current: &x,
		min:     &paddingWidth,
		max:     &widthBoundary,
	}
	yCoordinate := boundaryCoordinate{
		current: &y,
		min:     &paddingHeight,
		max:     &heightBoundary,
	}
	player := player{
		icon:     image,
		lives:    constants.DefaultLives - 3,
		position: &playerPosition{x: &xCoordinate, y: &yCoordinate},
	}
	player.Icon().UpdateOptionsCoordinates(x, y)

	return &player
}

func (p *player) Lives() int {
	return p.lives
}

func (p *player) Update(i Input) {
	p.Icon().Update(i, p.position)
}

func (p *player) Draw(screen *ebiten.Image) {
	screen.DrawImage(p.Icon().File(), p.Icon().Options())
}

func (p *player) Icon() CachedImage {
	return p.icon
}
