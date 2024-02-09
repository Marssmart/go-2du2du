package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/constants"
	"go-2du2du/game/movement"
	"go-2du2du/services"
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

func (p *player) Identity() string {
	return "Player"
}

func (p *player) Draw(screen *ebiten.Image, serviceContainer services.ServiceContainer, x float64, y float64) {
	serviceContainer.ImageDrawingService().Draw(screen, x, y, services.ImagePlayer)
}

func (p *player) NextMove() (error, movement.Movement) {
	return nil, movement.None
}
