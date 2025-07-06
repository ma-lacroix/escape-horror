package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

const (
	puzzleSizeX = 15
	puzzleSizeY = 15
)

type Puzzle struct {
	position PairFloat
}

func NewPuzzle() *Puzzle {
	return &Puzzle{PairFloat{screenWidth / 2, screenHeight / 2}}
}

func (p *Puzzle) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.position.x, p.position.y, puzzleSizeX, puzzleSizeY,
		color.RGBA{100, 50, 100, 255}, true)
}
