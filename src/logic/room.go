package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

const (
	begin  = 0.05
	midOne = 0.4
	midTwo = 0.6
	end    = 1 - begin
)

type Room struct {
	startingPoint bool
	endingPoint   bool
	doors         [4]bool
	furniture     []*Furniture
	traps         []*Trap
	puzzles       []*Puzzle
}

func (r *Room) DrawRoomBorders(screen *ebiten.Image) {
	strokeWidth := float32(6.0)
	borderColor := color.RGBA{100, 100, 10, 0xFF}
	// Top
	if r.doors[0] {
		vector.StrokeLine(screen, screenWidth*begin, screenHeight*begin, screenWidth*midOne, screenHeight*begin,
			strokeWidth, borderColor, true)
		vector.StrokeLine(screen, screenWidth*midTwo, screenHeight*begin, screenWidth*end, screenHeight*begin,
			strokeWidth, borderColor, true)
	} else {
		vector.StrokeLine(screen, screenWidth*begin, screenHeight*begin, screenWidth*end, screenHeight*begin,
			strokeWidth, borderColor, true)
	}
	// Left
	if r.doors[1] {
		vector.StrokeLine(screen, screenWidth*begin, screenHeight*begin, screenWidth*begin, screenHeight*midOne,
			strokeWidth, borderColor, true)
		vector.StrokeLine(screen, screenWidth*begin, screenHeight*midTwo, screenWidth*begin, screenHeight*end,
			strokeWidth, borderColor, true)
	} else {
		vector.StrokeLine(screen, screenWidth*begin, screenHeight*begin, screenWidth*begin, screenHeight*end,
			strokeWidth, borderColor, true)
	}
	// Bottom
	if r.doors[2] {
		vector.StrokeLine(screen, screenWidth*begin, screenHeight*end, screenWidth*midOne, screenHeight*end,
			strokeWidth, borderColor, true)
		vector.StrokeLine(screen, screenWidth*midTwo, screenHeight*end, screenWidth*end, screenHeight*end,
			strokeWidth, borderColor, true)
	} else {
		vector.StrokeLine(screen, screenWidth*begin, screenHeight*end, screenWidth*end, screenHeight*end,
			strokeWidth, borderColor, true)
	}
	// Right
	if r.doors[3] {
		vector.StrokeLine(screen, screenWidth*end, screenHeight*end, screenWidth*end, screenHeight*midTwo,
			strokeWidth, borderColor, true)
		vector.StrokeLine(screen, screenWidth*end, screenHeight*midOne, screenWidth*end, screenHeight*begin,
			strokeWidth, borderColor, true)
	} else {
		vector.StrokeLine(screen, screenWidth*end, screenHeight*end, screenWidth*end, screenHeight*begin,
			strokeWidth, borderColor, true)
	}
}

func (r *Room) Draw(screen *ebiten.Image) {
	bColor := color.RGBA{90, 90, 90, 50}
	if r.startingPoint {
		bColor = color.RGBA{150, 10, 10, 50}
	} else if r.endingPoint {
		bColor = color.RGBA{10, 150, 10, 50}
	}
	vector.DrawFilledRect(screen, screenWidth*begin, screenHeight*begin,
		screenWidth*(end-begin), screenHeight*(end-begin), bColor, true)
	r.DrawRoomBorders(screen)
	for _, f := range r.furniture {
		f.Draw(screen)
	}
}
