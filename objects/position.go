package objects

type Position interface {
	Update(input Input)
	UpdateCoordinates(x *boundaryCoordinate, y *boundaryCoordinate)
	AddToCoordinatePositionX(x float64)
	AddToCoordinatePositionY(y float64)
	X() *boundaryCoordinate
	Y() *boundaryCoordinate
}

type position struct {
	x *boundaryCoordinate
	y *boundaryCoordinate
}

type boundaryCoordinate struct {
	current *float64
	min     *float64
	max     *float64
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

func (r *position) Update(input Input) {
	if input.HasChanged() {
		switch input.LastInput() {
		case KeyDown:
			r.AddToCoordinatePositionY(5)
		case KeyLeft:
			r.AddToCoordinatePositionX(-5)
		case KeyRight:
			r.AddToCoordinatePositionX(5)
		case KeyUp:
			r.AddToCoordinatePositionY(-5)
		default:
		}
	}
}

func (r *position) AddToCoordinatePositionX(adjustment float64) {
	newValue := *r.x.current + adjustment
	r.x.current = &newValue
}

func (r *position) AddToCoordinatePositionY(adjustment float64) {
	newValue := *r.y.current + adjustment
	r.y.current = &newValue
}
