package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
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

func (p *Puzzle) DrawEnlargedPuzzleScreen(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, screenWidth*0.2, screenHeight*0.2, screenWidth*0.6, screenHeight*0.6,
		color.RGBA{150, 200, 100, 255}, true)
	text.Draw(screen, "PUZZLE", basicfont.Face7x13, screenWidth/3, screenHeight/3, color.Black)
}
