package movement

func NewEmptyMovementBehavior() Behavior {
	return &emptyMovementBehavior{}
}

type emptyMovementBehavior struct {
}

func (d *emptyMovementBehavior) Next(lastFailedDirection Direction) (error, Direction) {
	return nil, None
}

func (d *emptyMovementBehavior) Current() Direction {
	return None
}
