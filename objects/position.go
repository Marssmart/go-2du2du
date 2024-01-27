package objects

import "go-2du2du/constants"

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
