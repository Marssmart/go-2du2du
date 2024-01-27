package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/objects"
	"log"
)

func main() {
	ebiten.SetWindowTitle("Hello, 2du2du!")
	if err := ebiten.RunGame(objects.NewGame()); err != nil {
		log.Fatal(err)
	}
}
