package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/constants"
)

type StatusBar interface {
	Update(i Input)
	Draw(screen *ebiten.Image)
}

type statusBar struct {
	player        Player
	iconLifeFull  CachedImage
	iconLifeEmpty CachedImage
}

func NewStatusBar(player Player) StatusBar {
	iconLifeFullPath := constants.HeartFullIconPath
	iconLifeFull := NewImage(&iconLifeFullPath)
	iconLifeFull.PreLoadImage()
	iconLifeEmptyPath := constants.HeartEmptyIconPath
	iconLifeEmpty := NewImage(&iconLifeEmptyPath)
	iconLifeEmpty.PreLoadImage()
	return &statusBar{player: player, iconLifeFull: iconLifeFull, iconLifeEmpty: iconLifeEmpty}
}

func (s *statusBar) Update(i Input) {

}

func (s *statusBar) Draw(screen *ebiten.Image) {
	perIconSpace := float64(constants.DefaultIconWidth + constants.StatusBarIconGap)
	x := (constants.ScreenWidth / 2) - ((constants.DefaultLives) / 2 * perIconSpace)
	y := float64(constants.ScreenHeight - constants.ReservedRowsStatusBar)
	for i := 0; i < s.player.Lives(); i++ {
		s.iconLifeFull.UpdateOptionsCoordinates(x, y)
		screen.DrawImage(s.iconLifeFull.File(), s.iconLifeFull.Options())
		x = x + perIconSpace
	}

	for i := 0; i < constants.DefaultLives-s.player.Lives(); i++ {
		s.iconLifeEmpty.UpdateOptionsCoordinates(x, y)
		screen.DrawImage(s.iconLifeEmpty.File(), s.iconLifeEmpty.Options())
		x = x + perIconSpace
	}
}
