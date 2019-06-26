package main

import "fmt"

type Element interface{}

type DynamicList struct {
	data []interface{}
	size uint64
}

func (l *DynamicList) Size() uint64 {
	return l.size
}

func (l *DynamicList) Empty() bool {
	return l.size == 0
}

func (l *DynamicList) Insert(value Element, index uint64) {
	if index == l.size {
		l.data = append(l.data, value)
	}
	if index > l.size {
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

func (l *DynamicList) Remove(index uint64) {
	if l.Empty() {
		panic("List empty")
	}
	if index > l.size {
		panic("Invalid index")
	} else {
		for index < (l.size - 1) {
			l.data[index] = l.data[index+1]
			index++
		}
	}
	l.size--
}

func (l *DynamicList) Get(index uint64) Element {
	if index >= l.size {
		panic("Invalid index")
	} else {
		return l.data[index]
	}
}

func main() {
	list := new(DynamicList)
	list.Insert(0, 0)
	list.Insert(1, 1)
	list.Insert(2, 2)
	list.Remove(1)
	// list.Remove(0)
	fmt.Printf("Hello World %d\n", list.Get(1))
}
