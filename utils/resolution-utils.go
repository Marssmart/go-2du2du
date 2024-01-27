package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/constants"
	"image"
)

func Center(bounds *image.Rectangle) (float64, float64) {
	midX := float64(constants.ScreenWidth / 2)
	adjustmentX := (float64(bounds.Dx()) * constants.PlayerScale) / 2
	midY := float64(constants.ScreenHeight / 2)
	adjustmentY := (float64(bounds.Dy()) * constants.PlayerScale) / 2
	return midX - adjustmentX, midY - adjustmentY
}

func ScaledOptions(x float64, y float64) *ebiten.DrawImageOptions {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(constants.PlayerScale, constants.PlayerScale)
	options.GeoM.Translate(x, y)
	return options
}
