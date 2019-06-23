package main

import (
	"container/list"
	"fmt"
)

type Deque struct {
	data list.List
	size uint64
}

func (d *Deque) Size() uint64 {
	return d.size
}

func (d *Deque) Empty() bool {
	return d.size == 0
}

func (d *Deque) PushBack(value interface{}) {
	d.data.PushBack(value)
	d.size++
}

func (d *Deque) PushFront(value interface{}) {
	d.data.PushFront(value)
	d.size++
}

func (d *Deque) Front() interface{} {
	return d.data.Front().Value
}

func (d *Deque) Back() interface{} {
	return d.data.Back().Value
}

func (d *Deque) PopFront() {
	d.data.Remove(d.data.Front())
	d.size--
}

func (d *Deque) PopBack() {
	d.data.Remove(d.data.Back())
	d.size--
}

func main() {
	d := new(Deque)
	d.PushBack('a')
	d.PushFront('b')
	d.PushBack('c')
	d.PushFront('d')
	d.PopFront()
	d.PopBack()
	fmt.Printf("%c\n", d.Front())
	fmt.Printf("%c\n", d.Back())
}
