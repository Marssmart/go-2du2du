package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/constants"
)

type Game struct {
	player Player
	input  Input
}

func NewGame() *Game {
	newPlayer := NewPlayer()
	newInput := NewInput()
	return &Game{
		player: newPlayer,
		input:  newInput,
	}
}

func (g *Game) Update() error {
	g.input.Update()
	g.player.Update(g.input)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(int, int) (screenWidth int, screenHeight int) {
	return constants.ScreenWidth, constants.ScreenHeight
}
