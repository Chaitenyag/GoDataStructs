package main

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type LinkedList struct {
	head *Node
	size uint64
}

func (l *LinkedList) Size() uint64 {
	return l.size
}

func (l *LinkedList) Empty() bool {
	return l.size == 0
}

func (l *LinkedList) GetNode(index uint64) *Node {
	n := l.head
	for index > 0 {
		n = n.next
		index--
	}
	return n
}

func (l *LinkedList) Insert(value interface{}, index uint64) {
	n := &Node{value: value}
	if index == 0 {
		(*n).next = l.head
		l.head = n
	} else {
		prev := l.GetNode(index - 1)
		(*n).next = prev.next
		prev.next = n
	}
	l.size++
}

func (l *LinkedList) Remove(index uint64) {
	if index == 0 {
		n := l.head
		l.head = n.next
	} else {
		prev := l.GetNode(index - 1)
		n := prev.next
		prev.next = n.next
	}
	l.size--
}

func (l *LinkedList) Find(value interface{}) int {
	n := l.head
	index := 0
	for n != nil {
		if n.value == value {
			return index
		}
		n = n.next
		index++
	}
	return -1
}

func (l *LinkedList) Contains(value interface{}) bool {
	return l.Find(value) != -1
}

func (l *LinkedList) Get(index uint64) interface{} {
	n := l.head
	for index > 0 {
		n = n.next
		index--
	}
	return n.value
}

func main() {
	b := new(LinkedList)
	b.Insert(301.89, 0)
	b.Insert(1.0, 1)
	b.Insert(2.0, 2)
	b.Insert(5.0, 3)
	b.Insert(10.0, 4)
	b.Insert(20.0, 5)
	// b.Remove(1.0)
	// b.Remove(2.0)
	b.Remove(3)

	fmt.Printf("Binary Tree Size is empty: %f\n", b.Get(2))
}
