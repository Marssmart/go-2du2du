package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/constants"
	"go-2du2du/services"
)

type StatusBar interface {
	Draw(screen *ebiten.Image)
}

type statusBar struct {
	player           Player
	serviceContainer services.ServiceContainer
}

func NewStatusBar(player Player, serviceContainer services.ServiceContainer) StatusBar {
	return &statusBar{player: player, serviceContainer: serviceContainer}
}

func (s *statusBar) Draw(screen *ebiten.Image) {
	perIconSpace := float64(constants.DefaultIconWidth + constants.StatusBarIconGap)
	x := (constants.ScreenWidth / 2) - ((constants.DefaultLives) / 2 * perIconSpace)
	y := float64(constants.ScreenHeight - constants.ReservedRowsSpaceForStatusBar)
	for i := 0; i < s.player.Lives(); i++ {
		s.serviceContainer.ImageDrawingService().Draw(screen, x, y, services.ImageHeartFull)
		x = x + perIconSpace
	}

	for i := 0; i < constants.DefaultLives-s.player.Lives(); i++ {
		s.serviceContainer.ImageDrawingService().Draw(screen, x, y, services.ImageHeartEmpty)
		x = x + perIconSpace
	}
}
