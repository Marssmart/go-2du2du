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

type boundaryCoordinate struct {
	current *float64
	min     *float64
	max     *float64
}

type cachedImage struct {
	path    *string
	file    *ebiten.Image
	options *ebiten.DrawImageOptions
	x       *boundaryCoordinate
	y       *boundaryCoordinate
}

func (b *boundaryCoordinate) Update(value float64) bool {
	newValue := *b.current + value
	if newValue >= *b.min && newValue <= *b.max {
		b.current = &newValue
		return true
	}
	return false
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
	var padding float64 = 5
	var width float64 = constants.ScreenWidth
	var height float64 = constants.ScreenHeight
	i.x = &boundaryCoordinate{
		current: &x,
		min:     &padding,
		max:     &width,
	}
	i.y = &boundaryCoordinate{
		current: &y,
		min:     &padding,
		max:     &height,
	}
}

func (i *cachedImage) Update(input Input) {
	if input.HasChanged() {
		//log.Printf("Player icon coordnates before (%v,%v)", i.x, i.y)
		switch input.LastInput() {
		case KeyDown:
			i.y.Update(5)
		case KeyLeft:
			i.x.Update(-5)
		case KeyRight:
			i.x.Update(5)
		case KeyUp:
			i.y.Update(-5)
		default:
		}
		//log.Printf("Player icon coordnates before (%v,%v)", i.x, i.y)
		updateGeometry(&i.Options().GeoM, i.x, i.y)
	}
}

func updateGeometry(geometry *ebiten.GeoM, x *boundaryCoordinate, y *boundaryCoordinate) {
	geometry.Reset()
	geometry.Scale(constants.PlayerScale, constants.PlayerScale)
	geometry.Translate(*x.current, *y.current)
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
