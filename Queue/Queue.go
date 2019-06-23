package queue

import (
	"container/list"
	"fmt"
)

type Queue struct {
	data list.List
	size uint64
}

func (q *Queue) Size() uint64 {
	return q.size
}

func (q *Queue) Empty() bool {
	return q.size == 0
}

func (q *Queue) Push(value interface{}) {
	q.data.PushBack(value)
	q.size++
}

func (q *Queue) Top() interface{} {
	return q.data.Front().Value
}

func (q *Queue) Pop() {
	q.data.Remove(q.data.Front())
	q.size--
}

func main() {
	q := new(Queue)
	q.Push('a')
	q.Push('b')
	q.Push('c')
	q.Pop()
	q.Pop()
	fmt.Printf("%c\n", q.Top())
}
