package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"go-2du2du/constants"
	"log"
)

type CachedImage interface {
	File() *ebiten.Image
	Options() *ebiten.DrawImageOptions
	Update(input Input)
}

func NewImage(path *string) CachedImage {
	var offset degrees = -45
	var current degrees = 0
	image := &cachedImage{path: path, position: &position{rotationOffset: &offset, currentAngle: &current}}
	image.Load()
	return image
}

type cachedImage struct {
	path     *string
	file     *ebiten.Image
	options  *ebiten.DrawImageOptions
	position Position
}

func (i *cachedImage) Load() {
	file, _, err := ebitenutil.NewImageFromFile(*i.path)
	if err != nil {
		log.Fatalf("Failed to preload image %v", i.path)
	}
	bounds := file.Bounds()
	x, y := Center(&bounds)
	i.options = ScaledOptions(x, y, i.position)
	i.file = file
	var padding float64 = 5
	var width float64 = constants.ScreenWidth
	var height float64 = constants.ScreenHeight

	i.position.UpdateCoordinates(&boundaryCoordinate{
		current: &x,
		min:     &padding,
		max:     &width,
	}, &boundaryCoordinate{
		current: &y,
		min:     &padding,
		max:     &height,
	})
}

func (i *cachedImage) Update(input Input) {
	if input.HasChanged() {
		switch input.LastInput() {
		case KeyDown:
			i.position.UpdateCoordinatePositionY(5)
		case KeyLeft:
			i.position.UpdateCoordinatePositionX(-5)
		case KeyRight:
			i.position.UpdateCoordinatePositionX(5)
		case KeyUp:
			i.position.UpdateCoordinatePositionY(-5)
		default:
		}
		i.position.Update(input)
		updateGeometry(&i.Options().GeoM, i.position.X(), i.position.Y(), i.position)
	}
}

func updateGeometry(geometry *ebiten.GeoM, x *boundaryCoordinate, y *boundaryCoordinate, r Position) {
	geometry.Reset()
	geometry.Rotate(r.Offset().InRadians() + r.Current().InRadians())
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
