package main

import (
	"container/list"
	"fmt"
)

type Steque struct {
	data list.List
	size uint64
}

func (steq *Steque) Size() uint64 {
	return steq.size
}

func (steq *Steque) Empty() bool {
	return steq.size == 0
}

func (steq *Steque) Push(value interface{}) {
	steq.data.PushBack(value)
	steq.size++
}

func (steq *Steque) Enqueue(value interface{}) {
	steq.data.PushFront(value)
	steq.size++
}

func (steq *Steque) Front() interface{} {
	return steq.data.Front().Value
}

func (steq *Steque) Dequeue() {
	steq.data.Remove(steq.data.Front())
	steq.size--
}

func main() {
	steq := new(Steque)
	steq.Push('a')
	steq.Enqueue('b')
	steq.Enqueue('c')
	fmt.Printf("%c\n", steq.Front())
}
