package main

import "fmt"

type Element interface{}

type Node struct {
	prev  *Node
	next  *Node
	value Element
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size uint64
}

func (l *DoublyLinkedList) Size() uint64 {
	return l.size
}

func (l *DoublyLinkedList) Empty() bool {
	return l.size == 0
}

func (l *DoublyLinkedList) GetNode(index uint64) *Node {
	if index >= l.size {
		return nil
	}
	n := l.head
	for index > 0 {
		n = n.next
		index--
	}
	return n
}

func (l *DoublyLinkedList) Insert(value Element, index uint64) {
	n := &Node{value: value}
	if l.head == nil && l.tail == nil {
		l.head = n
		l.tail = n
	} else {
		prev := l.GetNode(index - 1)
		(*n).next = prev.next
		(*n).prev = prev
		prev.next = n
	}
	l.size++
}

func (l *DoublyLinkedList) Remove(index uint64) {
	if l.size == 1 {
		l.head = nil
		l.tail = nil
	} else {
		prev := l.GetNode(index - 1)
		next := l.GetNode(index + 1)
		prev.next = next
		next.prev = prev
	}
	l.size--
}

func (l *DoublyLinkedList) Find(value Element) int {
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

func (l *DoublyLinkedList) Contains(value Element) bool {
	return l.Find(value) != -1
}

func (l *DoublyLinkedList) Get(index uint64) Element {
	n := l.head
	if index >= l.size {
		panic("Invalid index")
	}
	for index > 0 {
		n = n.next
		index--
	}
	return n.value
}

func main() {
	b := new(DoublyLinkedList)
	b.Insert(301.89, 0)
	b.Insert(1.0, 1)
	b.Insert(2.0, 2)
	b.Insert(5.0, 3)
	b.Insert(10.0, 4)
	b.Insert(20.0, 5)
	b.Remove(1)
	b.Remove(2)
	// b.Remove(3)

	fmt.Printf("Binary Tree Size is empty: %t\n", b.Contains(5.0))
}
