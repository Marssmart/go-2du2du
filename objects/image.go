package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type CachedImage interface {
	File() *ebiten.Image
	Options() *ebiten.DrawImageOptions
	Update(input Input, position PlayerPosition)
	PreLoadImage()
	UpdateOptionsCoordinates(x float64, y float64)
}

func NewImage(path *string) CachedImage {
	image := &cachedImage{path: path}
	return image
}

type cachedImage struct {
	path    *string
	file    *ebiten.Image
	options *ebiten.DrawImageOptions
}

func (i *cachedImage) PreLoadImage() {
	file, _, err := ebitenutil.NewImageFromFile(*i.path)
	if err != nil {
		log.Fatalf("Failed to preload image %v", i.path)
	}
	i.file = file
}

func (i *cachedImage) UpdateOptionsCoordinates(x float64, y float64) {
	if i.options == nil {
		i.options = NewImageOptions(x, y)
	} else {
		updateGeometry(&i.options.GeoM, x, y)
	}
}

func (i *cachedImage) Update(input Input, position PlayerPosition) {
	if input.HasChanged() {
		position.Update(input)
		updateGeometry(&i.Options().GeoM, *position.X().current, *position.Y().current)
	}
}

func updateGeometry(geometry *ebiten.GeoM, x float64, y float64) {
	geometry.Reset()
	geometry.Translate(x, y)
}

func (i *cachedImage) File() *ebiten.Image {
	if i.file == nil {
		log.Fatalf("CachedImage %v file not preloaded", *i.path)
	}
	return i.file
}

func (i *cachedImage) Options() *ebiten.DrawImageOptions {
	if i.options == nil {
		log.Fatalf("CachedImage %v options not preloaded", *i.path)
	}
	return i.options
}
