package objects

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/game/movement"
	"go-2du2du/services"
	"log"
)

type devil struct {
	id       string
	behavior movement.Behavior
}

func NewDevil(id string) Object {
	return &devil{id, movement.NewCompassMovementBehavior()}
}

func (g *devil) Identity() string {
	return fmt.Sprintf("Devil(%s)", g.id)
}

func (g *devil) Draw(screen *ebiten.Image, serviceContainer services.ServiceContainer, x float64, y float64) {
	serviceContainer.ImageDrawingService().Draw(screen, x, y, services.ImageDevil)
}

// TODO - reuse this logic
func (g *devil) Move(x int, y int, maxX int, maxY int) (int, int) {
	current := g.behavior.Current()
	for i := 1; i <= 4; i++ {
		ok, rX, rY, lastAttemptedDir := current.Update(x, y, maxX, maxY)
		if ok {
			log.Printf("Moving %v from %v:%v to %v:%v", current.ToString(), x, y, rX, rY)
			return rX, rY
		}
		err, newDir := g.behavior.Next(lastAttemptedDir)
		if err != nil {
			panic("failed to get next direction")
		}
		log.Printf("Could not move %v, switching to %v", current.ToString(), newDir.ToString())
		current = newDir
	}
	panic("failed to find move")
}
