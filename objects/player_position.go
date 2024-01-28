package objects

import "go-2du2du/constants"

type PlayerPosition interface {
	Update(input Input)
	AddToCoordinatePositionX(x float64)
	AddToCoordinatePositionY(y float64)
	X() *boundaryCoordinate
	Y() *boundaryCoordinate
}

type playerPosition struct {
	x *boundaryCoordinate
	y *boundaryCoordinate
}

type boundaryCoordinate struct {
	current *float64
	min     *float64
	max     *float64
}

func (r *playerPosition) X() *boundaryCoordinate {
	return r.x
}

func (r *playerPosition) Y() *boundaryCoordinate {
	return r.y
}

func (r *playerPosition) Update(input Input) {
	if input.HasChanged() {
		switch input.LastInput() {
		case KeyDown:
			r.AddToCoordinatePositionY(constants.PlayerIconHeight)
		case KeyLeft:
			r.AddToCoordinatePositionX(-constants.PlayerIconWidth)
		case KeyRight:
			r.AddToCoordinatePositionX(constants.PlayerIconWidth)
		case KeyUp:
			r.AddToCoordinatePositionY(-constants.PlayerIconHeight)
		default:
		}
	}
}

func (r *playerPosition) AddToCoordinatePositionX(adjustment float64) {
	newValue := *r.x.current + adjustment
	if !r.checkValidX(newValue) {
		return
	}
	r.x.current = &newValue
}

func (r *playerPosition) AddToCoordinatePositionY(adjustment float64) {
	newValue := *r.y.current + adjustment
	if !r.checkValidY(newValue) {
		return
	}
	r.y.current = &newValue
}

func (r *playerPosition) checkValidX(value float64) bool {
	return value <= *r.x.max && value >= *r.x.min
}

func (r *playerPosition) checkValidY(value float64) bool {
	return value <= *r.y.max && value >= *r.y.min
}
