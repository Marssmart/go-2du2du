package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/services"
)

type Object interface {
	Identity() string
	Draw(screen *ebiten.Image, serviceContainer services.ServiceContainer, x float64, y float64)
}

type MovableObject interface {
	Move(x int, y int, maxX int, maxY int) (int, int)
}
