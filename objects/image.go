package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"go-2du2du/constants"
	"go-2du2du/utils"
	"log"
)

type CachedImage interface {
	File() *ebiten.Image
	Options() *ebiten.DrawImageOptions
	Update(input Input)
}

func NewImage(path *string) CachedImage {
	image := &cachedImage{path: path}
	image.Load()
	return image
}

type cachedImage struct {
	path    *string
	file    *ebiten.Image
	options *ebiten.DrawImageOptions
	x       float64
	y       float64
}

func (i *cachedImage) Load() {
	file, _, err := ebitenutil.NewImageFromFile(*i.path)
	if err != nil {
		log.Fatalf("Failed to preload image %v", i.path)
	}
	bounds := file.Bounds()
	x, y := utils.Center(&bounds)
	i.options = utils.ScaledOptions(x, y)
	i.file = file
	i.x = x
	i.y = y
}

func (i *cachedImage) Update(input Input) {
	if input.HasChanged() {
		log.Printf("Player icon coordnates before (%v,%v)", i.x, i.y)
		switch input.LastInput() {
		case KeyDown:
			i.y = i.y + 5
		case KeyLeft:
			i.x = i.x - 5
		case KeyRight:
			i.x = i.x + 5
		case KeyUp:
			i.y = i.y - 5
		default:
		}
		log.Printf("Player icon coordnates before (%v,%v)", i.x, i.y)
		updateGeometry(&i.Options().GeoM, &i.x, &i.y)
	}
}

func updateGeometry(geometry *ebiten.GeoM, x *float64, y *float64) {
	geometry.Reset()
	geometry.Scale(constants.PlayerScale, constants.PlayerScale)
	geometry.Translate(*x, *y)
}

func (i *cachedImage) File() *ebiten.Image {
	if i.file == nil {
		log.Fatalf("CachedImage %v not preloaded", i.path)
	}
	return i.file
}

func (i *cachedImage) Options() *ebiten.DrawImageOptions {
	if i.options == nil {
		log.Fatalf("CachedImage %v not preloaded", i.path)
	}
	return i.options
}
