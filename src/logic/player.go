package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Character int

const (
	Robbie Character = iota
	Amanda
)

const (
	playerSizeX = 50
	playerSizeY = 20
)

type Player struct {
	character   Character
	position    PairFloat
	currentRoom Pair
}

func (p *Player) checkWithinBoundaries(newMove PairFloat) bool {
	newX := p.position.x + newMove.x
	newY := p.position.y + newMove.y
	left := newX - playerSizeX/2
	right := newX + playerSizeX/2
	top := newY - playerSizeY/2
	bottom := newY + playerSizeY/2
	if left < screenWidth*begin || right > screenWidth*end ||
		top < screenHeight*begin || bottom > screenHeight*end {
		return false
	}
	return true
}

func (p *Player) checkWithinRoomTransfer(newMove PairFloat) bool {
	newX := p.position.x + newMove.x
	newY := p.position.y + newMove.y
	left := newX - playerSizeX/2
	right := newX + playerSizeX/2
	top := newY - playerSizeY/2
	bottom := newY + playerSizeY/2
	if (left <= screenWidth*begin && top >= screenHeight*midOne && bottom <= screenHeight*midTwo) ||
		(right >= screenWidth*end && top >= screenHeight*midOne && bottom <= screenHeight*midTwo) ||
		(top <= screenHeight*begin && left >= screenWidth*midOne && right <= screenWidth*midTwo) ||
		(bottom >= screenHeight*end && left >= screenWidth*midOne && right <= screenWidth*midTwo) {
		return true
	}
	return false
}

func (p *Player) Update(newMove PairFloat) {
	p.position.x += newMove.x
	p.position.y += newMove.y
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.position.x-playerSizeX/2, p.position.y-playerSizeY/2, playerSizeX, playerSizeY,
		color.RGBA{250, 50, 200, 255}, true)
}
