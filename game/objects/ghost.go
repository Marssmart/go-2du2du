package objects

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/game/movement"
	"go-2du2du/services"
	"log"
)

type Ghost interface {
}

type ghost struct {
	id       string
	behavior movement.Behavior
}

func NewGhost(id string) Object {
	return &ghost{id, movement.NewDiagonalMovementBehavior()}
}

func (g *ghost) Identity() string {
	return fmt.Sprintf("Ghost(%s)", g.id)
}

func (g *ghost) Draw(screen *ebiten.Image, serviceContainer services.ServiceContainer, x float64, y float64) {
	serviceContainer.ImageDrawingService().Draw(screen, x, y, services.ImageGhost)
}

func (g *ghost) Move(x int, y int, maxX int, maxY int) (int, int) {
	current := g.behavior.Current()
	for i := 1; i <= 4; i++ {
		ok, rX, rY := current.Update(x, y, maxX, maxY)
		if ok {
			log.Printf("Moving %v from %v:%v to %v:%v", current.ToString(), x, y, rX, rY)
			return rX, rY
		}
		err, newDir := g.behavior.Next()
		if err != nil {
			panic("failed to get next direction")
		}
		log.Printf("Could not move %v, switching to %v", current.ToString(), newDir.ToString())
		current = newDir
	}
	panic("failed to find move")
}
