package stack

import (
	"container/list"
	"fmt"
)

type Stack struct {
	data list.List
	size uint64
}

func (s *Stack) Size() uint64 {
	return s.size
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Push(value interface{}) {
	s.data.PushBack(value)
	s.size++
}

func (s *Stack) Top() interface{} {
	return s.data.Back().Value
}

func (s *Stack) Pop() {
	s.data.Remove(s.data.Back())
	s.size--
}

func main() {
	s := new(Stack)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Pop()
	s.Pop()
	fmt.Printf("%d\n", s.Top())
}
