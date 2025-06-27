package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type FurnitureType int

const (
	couch FurnitureType = iota
	chair
)

type DirectionType int

const (
	furnitureUp DirectionType = iota
	furnitureLeft
	furnitureDown
	furnitureRight
)

type Furniture struct {
	x1, x2, y1, y2 float32
}

func NewFurniture(furnitureType FurnitureType, directionType DirectionType) *Furniture {
	var x1, x2, y1, y2 float32
	switch furnitureType {
	case couch:
		switch directionType {
		case furnitureUp:
			x1 = screenWidth * 0.7
			x2 = screenWidth * 0.2
			y1 = screenHeight * 0.06
			y2 = screenHeight * 0.1
		case furnitureLeft:
			x1 = screenWidth * 0.06
			x2 = screenWidth * 0.1
			y1 = screenHeight * 0.06
			y2 = screenHeight * 0.2
		case furnitureDown:
			x1 = screenWidth * 0.07
			x2 = screenWidth * 0.2
			y1 = screenHeight * 0.84
			y2 = screenHeight * 0.1
		case furnitureRight:
			x1 = screenWidth * 0.84
			x2 = screenWidth * 0.1
			y1 = screenHeight * 0.7
			y2 = screenHeight * 0.2
		default:
			panic("unhandled default case")
		}
	case chair:
		switch directionType {
		case furnitureUp:
			x1 = screenWidth * 0.5
			x2 = screenWidth * 0.1
			y1 = screenHeight * 0.6
			y2 = screenHeight * 0.1
		case furnitureLeft:
			x1 = screenWidth * 0.2
			x2 = screenWidth * 0.1
			y1 = screenHeight * 0.2
			y2 = screenHeight * 0.1
		case furnitureDown:
			x1 = screenWidth * 0.12
			x2 = screenWidth * 0.1
			y1 = screenHeight * 0.70
			y2 = screenHeight * 0.1
		case furnitureRight:
			x1 = screenWidth * 0.7
			x2 = screenWidth * 0.1
			y1 = screenHeight * 0.6
			y2 = screenHeight * 0.1
		default:
			panic("unhandled default case")
		}
	}
	return &Furniture{x1, x2, y1, y2}
}

func (f *Furniture) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, f.x1, f.y1,
		f.x2, f.y2, color.RGBA{50, 50, 50, 255}, true)
}
