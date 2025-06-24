package main

import (
	"escape-horror/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	screenWidth  = 800
	screenHeight = 800
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Escape Horror Family!")

	g := logic.NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
