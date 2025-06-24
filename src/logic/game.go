package logic

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	ScreenWidth, ScreenHeight int
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func NewGame(screenWidth, screenHeight int) *Game {
	return &Game{
		screenWidth,
		screenHeight,
	}
}

func (game *Game) Update() error {
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {}
