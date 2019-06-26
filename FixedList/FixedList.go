package main

import "fmt"

type Element interface{}

type FixedList struct {
	data [6]interface{}
	size uint64
}

func (l *FixedList) Size() uint64 {
	return l.size
}

func (l *FixedList) Empty() bool {
	return l.size == 0
}

func (l *FixedList) Insert(value Element, index uint64) {
	if l.size >= 6 {
		panic("List full")
	} else if index > l.size {
		panic("Invalid index")
	} else {
		for index < l.size {
			l.data[index] = l.data[index-1]
			index++
		}
		l.data[index] = value
	}
	l.size++
}

func (l *FixedList) Remove(index uint64) {
	if l.Empty() {
		panic("List empty")
	}
	if index > l.size {
		panic("Invalid index")
	} else {
		for index < l.size {
			l.data[index] = l.data[index+1]
			index++
		}
	}
	l.size--
}

func (l *FixedList) Get(index uint64) Element {
	if index >= l.size {
		panic("Invalid index")
	} else {
		return l.data[index]
	}
}

func main() {
	list := new(FixedList)
	list.Insert(0, 0)
	list.Insert(1, 1)
	list.Remove(0)
	fmt.Printf("Hello World %d\n", list.Get(0))
}
