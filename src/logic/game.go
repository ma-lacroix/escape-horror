package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

const (
	screenWidth  = 800
	screenHeight = 800
)

type Game struct {
	ScreenWidth, ScreenHeight int
	HouseLayout               [3][4]bool
	Rooms                     [6]*Room
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func generateRooms(layout [3][4]bool) [6]*Room {
	// Graph traversal to generate doors & rooms
	return [6]*Room{}
}

func generateHouseLayout() [3][4]bool {
	// TODO: randomise this
	return [3][4]bool{{true, true, false, false},
		{false, true, true, false},
		{false, false, true, true}}
}

func NewGame(screenWidth, screenHeight int) *Game {
	doors := [4]bool{true, false, true, false}
	rooms := [6]*Room{NewRoom(doors)}
	houseLayout := generateHouseLayout()
	return &Game{
		screenWidth,
		screenHeight,
		houseLayout,
		rooms,
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 250, G: 250, B: 255, A: 255})
	g.Rooms[0].Draw(screen)
}
