package services

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"go-2du2du/constants"
	"log"
)

type ImageKey int

const (
	ImagePlayer ImageKey = iota
	ImageGhost
	ImageDevil
	ImageHeartFull
	ImageHeartEmpty
)

type ImageDrawingService interface {
	Draw(screen *ebiten.Image, x float64, y float64, image ImageKey)
}

func NewImageDrawingService() ImageDrawingService {
	service := imageDrawingService{imageIndex: make(map[ImageKey]*imageContainer)}
	options := ebiten.DrawImageOptions{}
	service.imageIndex[ImagePlayer] = &imageContainer{image: LoadImage(constants.PlayerIconPath), options: &options}
	service.imageIndex[ImageGhost] = &imageContainer{image: LoadImage(constants.GhostIconPath), options: &options}
	service.imageIndex[ImageDevil] = &imageContainer{image: LoadImage(constants.DevilIconPath), options: &options}
	service.imageIndex[ImageHeartFull] = &imageContainer{image: LoadImage(constants.HeartFullIconPath), options: &options}
	service.imageIndex[ImageHeartEmpty] = &imageContainer{image: LoadImage(constants.HeartEmptyIconPath), options: &options}
	return &service
}

func LoadImage(path string) *ebiten.Image {
	file, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatalf("Failed to load image %v", constants.PlayerIconPath)
	}
	return file
}

type imageContainer struct {
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
}

type imageDrawingService struct {
	imageIndex map[ImageKey]*imageContainer
}

func (s *imageDrawingService) Draw(screen *ebiten.Image, x float64, y float64, image ImageKey) {
	container := s.imageIndex[image]
	if container == nil {
		log.Fatalf("Failed to find image for key %v", image)
	}
	container.options.GeoM.Reset()
	container.options.GeoM.Translate(x, y)
	screen.DrawImage(container.image, container.options)
}
