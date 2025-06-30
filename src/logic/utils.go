package logic

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Offer(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Poll() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

type Pair struct {
	x int
	y int
}

type PairFloat struct {
	x float32
	y float32
}

func loadImage(path string) *ebiten.Image {
	data, err := imageFS.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read embedded image %s: %v", path, err)
	}
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatalf("Failed to decode image %s: %v", path, err)
	}
	return ebiten.NewImageFromImage(img)
}
