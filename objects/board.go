package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-2du2du/constants"
	"go-2du2du/services"
)

type CellType int

const (
	Empty CellType = iota
	Life
	Ghost
	Self
)

type Board interface {
	UpdateType(t CellType, x int, y int)
	Type(x int, y int) CellType
	Coordinates(x int, y int) (float64, float64)
	Draw(screen *ebiten.Image)
	Update(input Input)
}

func NewBoard(x int, y int, width float64, height float64, serviceContainer services.ServiceContainer) Board {
	var cells = make([][]*cell, x)
	for i := range cells {
		cells[i] = make([]*cell, y)
		for j := 0; j < len(cells[i]); j++ {
			cells[i][j] = &cell{t: Empty, x: i, y: j}
		}
	}
	b := board{cells: cells, width: width, height: height, serviceContainer: serviceContainer}
	b.playerPosition = b.cells[(x/2)-1][(y/2)-1]
	b.playerPosition.UpdateType(Self)

	return &b
}

type Cell interface {
	UpdateType(t CellType)
	Type() CellType
	Swap(r *board, c *cell, x int, y int) *cell
}

type board struct {
	width          float64
	height         float64
	cells          [][]*cell
	playerPosition *cell

	serviceContainer services.ServiceContainer
}

// having coordinates also here makes it easier to make position adjustments without storing last position
type cell struct {
	t CellType
	x int
	y int
}

func (b *board) Swap(sourceCell *cell, x int, y int) *cell {
	maxX := len(b.cells)
	maxY := len(b.cells[0])

	if maxX <= x || maxY <= y || x < 0 || y < 0 {
		return sourceCell
	}

	sourceCellType := sourceCell.Type()
	targetCell := b.cells[x][y]
	targetCellType := targetCell.Type()
	targetCell.UpdateType(sourceCellType)
	sourceCell.UpdateType(targetCellType)
	return targetCell
}

func (b *board) Update(input Input) {
	if input.HasChanged() {
		x := b.playerPosition.x
		y := b.playerPosition.y
		switch input.LastInput() {
		case KeyDown:
			y++
		case KeyLeft:
			x--
		case KeyRight:
			x++
		case KeyUp:
			y--
		default:
			return
		}
		b.playerPosition = b.Swap(b.playerPosition, x, y)
	}
}

func (b *board) Draw(screen *ebiten.Image) {
	for x := 0; x < len(b.cells); x++ {
		for y := 0; y < len(b.cells[x]); y++ {
			switch b.cells[x][y].Type() {
			case Life:
			case Self:
				pX, pY := b.Coordinates(x, y)
				b.serviceContainer.ImageDrawingService().Draw(screen, pX, pY, services.ImagePlayer)
			case Ghost:
				pX, pY := b.Coordinates(x, y)
				b.serviceContainer.ImageDrawingService().Draw(screen, pX, pY, services.ImageGhost)
			case Empty:
			default:

			}
		}
	}
}

func (b *board) Coordinates(x int, y int) (float64, float64) {
	sX := float64(constants.MarginX + (constants.ScreenHeight / constants.BoardItemIconHeight * x) + (x * (constants.BoardItemIconHeight / 2)))
	sY := float64(constants.MarginY + (constants.ScreenWidth / constants.BoardItemIconWidth * y) + (y * (constants.BoardItemIconWidth / 2)))
	return sX, sY
}

func (b *board) UpdateType(t CellType, x int, y int) {
	b.cells[x][y].UpdateType(t)
}

func (b *board) Type(x int, y int) CellType {
	return b.cells[x][y].Type()
}

func (c *cell) UpdateType(t CellType) {
	c.t = t
}

func (c *cell) Type() CellType {
	return c.t
}
