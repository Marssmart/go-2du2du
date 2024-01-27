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

type Position interface {
	Update(input Input)
	UpdateCoordinates(x *boundaryCoordinate, y *boundaryCoordinate)
	UpdateCoordinatePositionX(x float64)
	UpdateCoordinatePositionY(y float64)
	Offset() *degrees
	Current() *degrees
	X() *boundaryCoordinate
	Y() *boundaryCoordinate
}

func NewImage(path *string) CachedImage {
	var offset degrees = -45
	var current degrees = 0
	image := &cachedImage{path: path, position: &position{rotationOffset: &offset, currentAngle: &current}}
	image.Load()
	return image
}

type degrees float64

type position struct {
	//how many degrees the image has to be rotated to face up
	rotationOffset *degrees
	// degrees of rotation
	currentAngle *degrees
	x            *boundaryCoordinate
	y            *boundaryCoordinate
}

type boundaryCoordinate struct {
	current *float64
	min     *float64
	max     *float64
}

type cachedImage struct {
	path     *string
	file     *ebiten.Image
	options  *ebiten.DrawImageOptions
	position Position
}

func (d *degrees) InRadians() float64 {
	return float64(*d) * (constants.PI / 180)
}

func (r *position) Update(i Input) {
	//TODO fix rotation counterclockwise
	if i.HasChanged() {
		switch i.LastInput() {
		case KeyUp:
			r.currentAngle.update(0)
		case KeyDown:
			r.currentAngle.update(180)
		case KeyLeft:
			r.currentAngle.update(270)
		case KeyRight:
			r.currentAngle.update(90)
		default:
			//noop
		}
	}
}

func (r *position) Offset() *degrees {
	return r.rotationOffset
}

func (r *position) Current() *degrees {
	return r.currentAngle
}

func (r *position) X() *boundaryCoordinate {
	return r.x
}

func (r *position) Y() *boundaryCoordinate {
	return r.y
}

func (r *position) UpdateCoordinates(x *boundaryCoordinate, y *boundaryCoordinate) {
	r.x = x
	r.y = y
}

func (r *position) UpdateCoordinatePositionX(adjustment float64) {
	newValue := *r.x.current + adjustment
	r.x.current = &newValue
}

func (r *position) UpdateCoordinatePositionY(adjustment float64) {
	newValue := *r.y.current + adjustment
	r.y.current = &newValue
}

func (d *degrees) update(newValue float64) {
	newRotation := degrees(newValue)
	*d = newRotation
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
