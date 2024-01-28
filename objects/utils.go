package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/constants"
	"image"
)

func Center(bounds *image.Rectangle) (float64, float64) {
	midX := float64((constants.ScreenWidth - constants.PaddingWidth) / 2)
	adjustmentX := float64(bounds.Dx() / 2)
	midY := float64((constants.ScreenHeight - constants.PaddingHeight) / 2)
	adjustmentY := float64(bounds.Dy() / 2)
	return midX - adjustmentX, midY - adjustmentY
}

func NewImageOptions(x float64, y float64) *ebiten.DrawImageOptions {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(x, y)
	return options
}
