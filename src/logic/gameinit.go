package logic

import (
	"errors"
	"math/rand"
)

func getDoorLocations(i int, j int, layout *[c][r]bool) [4]bool {
	doors := [4]bool{}
	// top
	if i > 0 && layout[i-1][j] {
		doors[0] = true
	}
	// left
	if j > 0 && layout[i][j-1] {
		doors[1] = true
	}
	// bottom
	if i < c-1 && layout[i+1][j] {
		doors[2] = true
	}
	// right
	if j < r-1 && layout[i][j+1] {
		doors[3] = true
	}
	return doors
}

func checkStartingPoint(i int, j int) bool {
	return i == 0 && j == 0
}

func checkEndingPoint(i int, j int) bool {
	return i == c-1 && j == r-1
}

func generateFurniture() []*Furniture {
	var furniture []*Furniture
	num := 0
	for num < 3 {
		val1 := rand.Float32()
		val2 := rand.Float32()
		var furnitureType FurnitureType
		var furnitureDirection DirectionType
		if 0.75 < val2 {
			furnitureDirection = furnitureUp
		} else if 0.5 < val2 {
			furnitureDirection = furnitureLeft
		} else if 0.25 < val2 {
			furnitureDirection = furnitureDown
		} else {
			furnitureDirection = furnitureRight
		}
		if 0.75 < val1 {
			furnitureType = couch
		} else if 0.5 < val1 {
			furnitureType = chair
		}
		if 0.5 < val1 {
			furniture = append(furniture, NewFurniture(furnitureType, furnitureDirection))
		}
		num++
	}
	return furniture
}

func generatePuzzles() *Puzzle {
	// Add more instructions to define which puzzle this is
	return NewPuzzle()
}

func generateRooms(layout *[c][r]bool) map[Pair]*Room {
	// this function will initialize the rooms, doors, connections, traps, puzzles and displayed furniture
	rooms := make(map[Pair]*Room)
	puzzleCounter := 3
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			if !layout[i][j] {
				continue
			}

			doors := getDoorLocations(i, j, layout)
			furniture := generateFurniture()
			startingPoint := checkStartingPoint(i, j)
			endingPoint := checkEndingPoint(i, j)
			if !(i == 0 && j == 0) && i%2 == 0 && j%2 == 0 && puzzleCounter > 0 {
				// temp logic while I figure out a better way to distribute puzzles and challenges
				puzzle := generatePuzzles()
				puzzleCounter--
				rooms[Pair{i, j}] = &Room{
					doors:         doors,
					startingPoint: startingPoint,
					endingPoint:   endingPoint,
					puzzle:        puzzle,
					furniture:     furniture}
			} else {
				rooms[Pair{i, j}] = &Room{
					doors:         doors,
					startingPoint: startingPoint,
					endingPoint:   endingPoint,
					furniture:     furniture}
			}
		}
	}
	return rooms
}

func validateHouseLayout(layout [c][r]bool) bool {
	// Top left to bottom right BFS
	startPos := &Pair{0, 0}
	q := Queue[*Pair]{}
	q.Offer(startPos)
	for !q.IsEmpty() {
		currentPos, _ := q.Poll()
		if currentPos.x == r-1 && currentPos.y == c-1 {
			return true
		}
		if currentPos.y < c-1 && layout[currentPos.y+1][currentPos.x] {
			q.Offer(&Pair{currentPos.x, currentPos.y + 1})
		}
		if currentPos.x < r-1 && layout[currentPos.y][currentPos.x+1] {
			q.Offer(&Pair{currentPos.x + 1, currentPos.y})
		}
	}
	return false
}

func randomHouseLayout() [c][r]bool {
	layout := [c][r]bool{}
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			val := rand.Float32()
			if val < 0.70 {
				layout[i][j] = true
			} else {
				layout[i][j] = false
			}
		}
	}
	// Start and end rooms
	layout[0][0] = true
	layout[c-1][r-1] = true
	return layout
}

func generateHouseLayout() ([c][r]bool, error) {
	p := 0
	check := false
	layout := [c][r]bool{}
	for !check && p < 6 {
		layout = randomHouseLayout()
		check = validateHouseLayout(layout)
		p++
	}
	if !check {
		return layout, errors.New("failed to generate a proper house layout")
	}
	return layout, nil
}
