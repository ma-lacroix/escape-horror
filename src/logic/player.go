package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
)

type Character int

const (
	Robbie Character = iota
	Amanda
)

const (
	playerSizeX = 20
	playerSizeY = 20
)

type Player struct {
	character   Character
	position    PairFloat
	currentRoom Pair
	// Not sure if the player items should be in this struct
	playerKeys int
	keysImage  *ebiten.Image
}

func NewPlayer(character Character) *Player {
	keysImage := loadImage("media/images/lock-and-key.png")
	return &Player{character: character,
		position:    PairFloat{float32(screenWidth / 2), float32(screenHeight / 2)},
		currentRoom: Pair{0, 0},
		playerKeys:  3,
		keysImage:   keysImage}
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

func (p *Player) checkWithinRoomTransfer(newMove PairFloat, doors [4]bool) Pair {
	newPair := Pair{0, 0}
	newX := p.position.x + newMove.x
	newY := p.position.y + newMove.y
	left := newX - playerSizeX/2
	right := newX + playerSizeX/2
	top := newY - playerSizeY/2
	bottom := newY + playerSizeY/2
	if top <= screenHeight*begin && left >= screenWidth*midOne && right <= screenWidth*midTwo && doors[0] {
		newPair = Pair{-1, 0}
	}
	if left <= screenWidth*begin && top >= screenHeight*midOne && bottom <= screenHeight*midTwo && doors[1] {
		newPair = Pair{0, -1}
	}
	if bottom >= screenHeight*end && left >= screenWidth*midOne && right <= screenWidth*midTwo && doors[2] {
		newPair = Pair{1, 0}
	}
	if right >= screenWidth*end && top >= screenHeight*midOne && bottom <= screenHeight*midTwo && doors[3] {
		newPair = Pair{0, 1}
	}
	return newPair
}

func (p *Player) checkCollisionWithinFurniture(newMove PairFloat, furniture []*Furniture) bool {
	newX := p.position.x + newMove.x
	newY := p.position.y + newMove.y
	left := newX - playerSizeX/2
	right := newX + playerSizeX/2
	top := newY - playerSizeY/2
	bottom := newY + playerSizeY/2

	for _, f := range furniture {
		fLeft := f.x1
		fRight := fLeft + f.x2
		fTop := f.y1
		fBottom := fTop + f.y2
		if left < fRight &&
			right > fLeft &&
			top < fBottom &&
			bottom > fTop {
			return false
		}
	}
	return true
}

func (p *Player) Update(newMove PairFloat) {
	p.position.x += newMove.x
	p.position.y += newMove.y
}

func (p *Player) drawKeys(screen *ebiten.Image) {

	w, h := p.keysImage.Size()
	spriteW := w / 4
	spriteH := h / 2

	srcRect := image.Rect(
		spriteW,
		0,
		(p.playerKeys+1)*spriteW,
		spriteH,
	)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.4, 0.4)
	op.GeoM.Translate(screenWidth*(1-(float64(spriteW)/float64(screenHeight)*0.5*float64(p.playerKeys))), 20)

	screen.DrawImage(p.keysImage.SubImage(srcRect).(*ebiten.Image), op)
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.position.x-playerSizeX/2, p.position.y-playerSizeY/2, playerSizeX, playerSizeY,
		color.RGBA{250, 50, 200, 255}, true)
	p.drawKeys(screen)
}
