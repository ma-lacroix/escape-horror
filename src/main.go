package main

import (
	"escape-horror/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Escape Horror Family!")

	g := logic.NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
