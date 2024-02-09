package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/game/objects"
	"go-2du2du/services"
	"log"
)

func main() {
	container := services.NewServiceContainer()
	container.RegisterImageDrawingService(services.NewImageDrawingService())

	ebiten.SetWindowTitle("Hello, 2du2du!")
	if err := ebiten.RunGame(objects.NewGame(container)); err != nil {
		log.Fatal(err)
	}
}
