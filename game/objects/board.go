package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/constants"
	"go-2du2du/game"
	"go-2du2du/services"
	"time"
)

type Board interface {
	Coordinates(x int, y int) (float64, float64)
	Draw(screen *ebiten.Image)
	Update(input game.Input)
	MaxX() int
	MaxY() int
}

func NewBoard(x int, y int, width float64, height float64, serviceContainer services.ServiceContainer, player Object) Board {
	var cells = make([][]*cell, x)
	emptyObject := NewEmptyObject()
	for i := range cells {
		cells[i] = make([]*cell, y)
		for j := 0; j < len(cells[i]); j++ {
			cells[i][j] = &cell{x: i, y: j, object: emptyObject}
		}
	}
	b := board{cells: cells, width: width, height: height, serviceContainer: serviceContainer, lastMoveTime: time.Now()}
	b.playerPosition = b.cells[(x/2)-1][(y/2)-1]
	b.playerPosition.object = player

	//set ghost mobs
	mobCell := b.cells[0][0]
	mobCell.object = NewGhost("first")

	return &b
}

type board struct {
	width          float64
	height         float64
	cells          [][]*cell
	playerPosition *cell
	lastMoveTime   time.Time

	serviceContainer services.ServiceContainer
}

func (b *board) Swap(sourceCell *cell, targetCell *cell) {
	/*if _, ok := sourceCell.object.(Player); ok {
		return
	}*/

	if _, ok := targetCell.object.(Player); ok {
		return
	}

	oldTargetCellObject := targetCell.object
	targetCell.object = sourceCell.object
	sourceCell.object = oldTargetCellObject
}

func (b *board) Update(input game.Input) {
	if input.HasChanged() {
		x := b.playerPosition.x
		y := b.playerPosition.y
		switch input.LastInput() {
		case game.KeyDown:
			y++
		case game.KeyLeft:
			x--
		case game.KeyRight:
			x++
		case game.KeyUp:
			y--
		default:
			return
		}
		b.Swap(b.playerPosition, b.cells[x][y])
		b.playerPosition = b.cells[x][y]
	}

	now := time.Now()

	//if b.lastMoveTime.Second() != now.Second() {
	if now.UnixMilli()-b.lastMoveTime.UnixMilli() > 200 {
		//log.Printf("Moving on time %s", now)

		movableCells := make([]*cell, 0)

		//TODO fix the movement skip
		for x := range b.cells {
			for _, c := range b.cells[x] {
				_, ok := c.object.(MovableObject)
				if ok {
					movableCells = append(movableCells, c)
				}
			}
		}

		for _, c := range movableCells {
			movableObject, ok := c.object.(MovableObject)
			if ok {
				x, y := movableObject.Move(c.x, c.y, b.MaxX(), b.MaxY())
				b.Swap(c, b.cells[x][y])
			}
		}

		b.lastMoveTime = now
	}
}

func (b *board) Draw(screen *ebiten.Image) {
	for x := 0; x < len(b.cells); x++ {
		for y := 0; y < len(b.cells[x]); y++ {
			pX, pY := b.Coordinates(x, y)
			b.cells[x][y].object.Draw(screen, b.serviceContainer, pX, pY)
		}
	}
}

func (b *board) Coordinates(x int, y int) (float64, float64) {
	sX := float64(constants.MarginX + (constants.ScreenHeight / constants.BoardItemIconHeight * x) + (x * (constants.BoardItemIconHeight / 2)))
	sY := float64(constants.MarginY + (constants.ScreenWidth / constants.BoardItemIconWidth * y) + (y * (constants.BoardItemIconWidth / 2)))
	return sX, sY
}

func (b *board) MaxX() int {
	return len(b.cells) - 1
}

func (b *board) MaxY() int {
	return len(b.cells[0]) - 1
}
