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
	doors         [4]bool
	furniture     []Furniture
	traps         []Trap
	puzzles       []Puzzle
}

func NewRoom(doorLocations [4]bool) *Room {
	return &Room{
		startingPoint: true,
		doors:         doorLocations,
		furniture:     []Furniture{},
		traps:         []Trap{},
		puzzles:       []Puzzle{},
	}
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
	r.DrawRoomBorders(screen)
}
