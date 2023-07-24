package queue

import "errors"

type queue struct {
	front    int
	contents []any
}

func (q *queue) Insert(elem any) {
	q.contents = append(q.contents, elem)
}

func (q queue) Peek() (any, error) {
	if len(q.contents) == 0 {
		return nil, errors.New("No items in Queue")
	}

	return q.contents[q.front], nil
}

func (q *queue) Pop() any {
	elem := q.contents[q.front]
	q.front++

	return elem
}
