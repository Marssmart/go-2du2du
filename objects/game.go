package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/constants"
	"image/color"
)

type Game struct {
	player    Player
	input     Input
	statusBar StatusBar
}

func NewGame() *Game {
	newPlayer := NewPlayer()
	newInput := NewInput()
	newStatusBar := NewStatusBar(newPlayer)
	return &Game{
		player:    newPlayer,
		input:     newInput,
		statusBar: newStatusBar,
	}
}

func (g *Game) Update() error {
	g.input.Update()
	g.player.Update(g.input)
	g.statusBar.Update(g.input)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(color.White)
	g.player.Draw(screen)
	g.statusBar.Draw(screen)
}

func (g *Game) Layout(int, int) (screenWidth int, screenHeight int) {
	return constants.ScreenWidth, constants.ScreenHeight
}
