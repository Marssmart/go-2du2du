package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/constants"
	"go-2du2du/game"
	"go-2du2du/services"
	"image/color"
)

type Game struct {
	player    Player
	input     game.Input
	statusBar StatusBar
	board     Board

	serviceContainer services.ServiceContainer
}

func NewGame(serviceContainer services.ServiceContainer) *Game {
	newPlayer := NewPlayer()
	newInput := game.NewInput()
	newStatusBar := NewStatusBar(newPlayer, serviceContainer)
	newBoard := NewBoard(constants.Columns, constants.Rows, constants.BoardItemWidthBoundary, constants.BoardItemHeightBoundary, serviceContainer, newPlayer.(Object))

	return &Game{
		player:           newPlayer,
		input:            newInput,
		statusBar:        newStatusBar,
		board:            newBoard,
		serviceContainer: serviceContainer,
	}
}

func (g *Game) Update() error {
	g.input.Update()
	g.board.Update(g.input)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(color.White)
	g.board.Draw(screen)
	g.statusBar.Draw(screen)
}

func (g *Game) Layout(int, int) (screenWidth int, screenHeight int) {
	return constants.ScreenWidth, constants.ScreenHeight
}
