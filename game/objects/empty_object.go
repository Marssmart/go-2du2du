package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/game/movement"
	"go-2du2du/services"
)

func NewEmptyObject() Object {
	return &emptyObject{}
}

type emptyObject struct {
}

func (g *emptyObject) Identity() string {
	return "Empty"
}

func (g *emptyObject) Draw(screen *ebiten.Image, serviceContainer services.ServiceContainer, x float64, y float64) {
	//nothing
}

func (g *emptyObject) NextMove() (error, movement.Movement) {
	return nil, movement.None
}
