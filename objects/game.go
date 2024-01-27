package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/constants"
	"image/color"
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
	screen.Clear()
	screen.Fill(color.White)
	screen.DrawImage(g.player.Icon().File(), g.player.Icon().Options())
}

func (g *Game) Layout(int, int) (screenWidth int, screenHeight int) {
	return constants.ScreenWidth, constants.ScreenHeight
}
