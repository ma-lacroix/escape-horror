package logic

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

type GameStatus int

const (
	Menu GameStatus = iota
	MapNavigationTime
	RoamingTime
	PuzzleTime
)

const (
	screenWidth  = 400
	screenHeight = 400
	c            = 3
	r            = 4
	playerSpeed  = 5.0
)

//go:embed media/images/*
var imageFS embed.FS

type Game struct {
	Status                    GameStatus
	ScreenWidth, ScreenHeight int
	HouseLayout               [c][r]bool
	Rooms                     map[Pair]*Room
	CurrentRoom               Pair
	moveCooldown              int
	moveCooldownMax           int
	Player                    *Player
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func NewGame(screenWidth, screenHeight int) *Game {
	houseLayout, e := generateHouseLayout()
	if e != nil {
		panic(e)
	}
	rooms := generateRooms(&houseLayout)
	return &Game{
		RoamingTime,
		screenWidth,
		screenHeight,
		houseLayout,
		rooms,
		Pair{0, 0},
		8,
		8,
		NewPlayer(Robbie),
	}
}

func (g *Game) HandleMapNavigation() {
	// debug mode -> to be deleted once the game is complete
	if g.moveCooldown > 0 {
		g.moveCooldown--
		return
	}
	newPair := g.CurrentRoom
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		newPair = Pair{g.CurrentRoom.x - 1, g.CurrentRoom.y}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		newPair = Pair{g.CurrentRoom.x + 1, g.CurrentRoom.y}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		newPair = Pair{g.CurrentRoom.x, g.CurrentRoom.y - 1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		newPair = Pair{g.CurrentRoom.x, g.CurrentRoom.y + 1}
	}
	if newPair != g.CurrentRoom {
		if _, ok := g.Rooms[newPair]; ok {
			g.CurrentRoom = newPair
		}
	}
	g.moveCooldown = g.moveCooldownMax
}

func (g *Game) HandleRoaming() {
	newMove := PairFloat{0.0, 0.0}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		newMove = PairFloat{0.0, -playerSpeed}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		newMove = PairFloat{0.0, playerSpeed}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		newMove = PairFloat{-playerSpeed, 0.0}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		newMove = PairFloat{playerSpeed, 0.0}
	}
	offset := g.Player.checkWithinRoomTransfer(newMove, g.Rooms[g.CurrentRoom].doors)
	if offset != (Pair{0, 0}) {
		newPair := Pair{g.CurrentRoom.x + offset.x, g.CurrentRoom.y + offset.y}
		if _, ok := g.Rooms[newPair]; ok {
			// screen wrapping
			if offset.x != 0 {
				if offset.x > 0 {
					g.Player.position.y = float32(screenWidth * 0.15)
				} else {
					g.Player.position.y = float32(screenWidth * 0.85)
				}
			}
			if offset.y != 0 {
				if offset.y > 0 {
					g.Player.position.x = float32(screenHeight * 0.15)
				} else {
					g.Player.position.x = float32(screenHeight * 0.85)
				}
			}
			g.CurrentRoom = newPair
			g.Player.currentRoom = newPair
		}
	}

	if (newMove.x != 0.0 || newMove.y != 0.0) && g.Player.checkWithinBoundaries(newMove) &&
		g.Player.checkCollisionWithinFurniture(newMove, g.Rooms[g.CurrentRoom].furniture) {
		g.Player.Update(newMove)
	}
	if g.Rooms[g.CurrentRoom].puzzle != nil && g.Player.checkCollisionWithPuzzle(g.Rooms[g.CurrentRoom].puzzle) {
		g.Status = PuzzleTime
	}
}

func (g *Game) HandlePuzzleTime() {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.Player.Update(PairFloat{0.0, -playerSizeY})
		g.Status = RoamingTime
	}
}

func (g *Game) ResetGame() {
	houseLayout, e := generateHouseLayout()
	if e != nil {
		panic(e)
	}
	rooms := generateRooms(&houseLayout)
	g.Rooms = rooms
	g.HouseLayout = houseLayout
	g.CurrentRoom = Pair{0, 0}
	g.Player.position = PairFloat{float32(screenWidth / 2), float32(screenHeight / 2)}
	g.Player.currentRoom = g.CurrentRoom
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.ResetGame()
	}
	if ebiten.IsKeyPressed(ebiten.Key1) {
		g.Status = RoamingTime
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		g.Status = MapNavigationTime
	}
	switch g.Status {
	case Menu:
		panic("Not implemented")
	case MapNavigationTime:
		g.HandleMapNavigation()
	case RoamingTime:
		g.HandleRoaming()
	case PuzzleTime:
		g.HandlePuzzleTime()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 250, G: 250, B: 255, A: 255})
	g.Rooms[g.CurrentRoom].Draw(screen)
	if g.Player.currentRoom == g.CurrentRoom {
		g.Player.Draw(screen)
	}
	if g.Status == PuzzleTime {
		g.Rooms[g.CurrentRoom].puzzle.DrawEnlargedPuzzleScreen(screen)
	}
	if g.Status == MapNavigationTime {
		text.Draw(screen, "DEBUG MODE", basicfont.Face7x13, 10, 10, color.Black)
	}
}
