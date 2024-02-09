package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputKey int

const (
	KeyUp    InputKey = iota
	KeyDown  InputKey = iota
	KeyLeft  InputKey = iota
	KeyRight InputKey = iota
	KeyNone  InputKey = iota
)

type Input interface {
	LastInput() InputKey
	Update()
	HasChanged() bool
}

func NewInput() Input {
	return &input{
		lastInput: KeyNone,
		changed:   false,
	}
}

type input struct {
	lastInput InputKey
	changed   bool
}

func (i *input) LastInput() InputKey {
	return i.lastInput
}

func (i *input) Update() {
	newLast := keyPressedWithoutRelease()
	i.changed = i.lastInput != newLast
	i.lastInput = newLast
}

func (i *input) HasChanged() bool {
	return i.changed
}

func keyPressedWithoutRelease() InputKey {
	if isKeyPressedWithoutRelease(ebiten.KeyArrowUp) {
		return KeyUp
	} else if isKeyPressedWithoutRelease(ebiten.KeyArrowDown) {
		return KeyDown
	} else if isKeyPressedWithoutRelease(ebiten.KeyArrowLeft) {
		return KeyLeft
	} else if isKeyPressedWithoutRelease(ebiten.KeyArrowRight) {
		return KeyRight
	} else {
		return KeyNone
	}
}

func isKeyPressedWithoutRelease(key ebiten.Key) bool {
	return inpututil.IsKeyJustPressed(key) && !inpututil.IsKeyJustReleased(key)
}
