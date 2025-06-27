package logic

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
