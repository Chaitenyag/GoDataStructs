// Package fixedlist - An implementation of a fixed list
// Copyright 2019, Mihir Wadekar and Chaitenya Gupta
package fixedlist

// Element - A blank interface to allow list of all datatypes
type Element interface{}

// FixedList - Struct to create a fixed list of size 6
type FixedList struct {
	// Array storing the data
	data [6]interface{}
	// Size of the list
	size uint64
}

// Size - Returns the size of the list
func (l *FixedList) Size() uint64 {
	return l.size
}

// Empty - Checks if the list is empty or not
func (l *FixedList) Empty() bool {
	return l.size == 0
}

// Insert - Inserts an element at the given position in the list
func (l *FixedList) Insert(value Element, index uint64) {
	if l.size >= 6 {
		// If the size of the list is greater than
		// or equal to 6, panic
		panic("List full")
	} else if index > l.size {
		// If the insertion index is greater than the
		// size of the list, panic
		panic("Invalid index")
	} else {
		// Copy the elements after the given index into
		// their next indexes
		for index < l.size {
			l.data[index] = l.data[index-1]
			index++
		}
		// Insert the value at the given index
		l.data[index] = value
	}
	// Increment the size
	l.size++
}

// Remove - Removes an element from the given position in the list
func (l *FixedList) Remove(index uint64) {
	if l.Empty() {
		// If the list is empty, panic
		panic("List empty")
	}
	if index > l.size {
		// If the removal index is greater than the
		// size of the list, panic
		panic("Invalid index")
	} else {
		// Copy the elements after the given index into
		// their previous indexes
		for index < l.size {
			l.data[index] = l.data[index+1]
			index++
		}
	}
	// Decrement the size
	l.size--
}

// Get - Gets an element from the given position in the list
func (l *FixedList) Get(index uint64) Element {
	if index >= l.size {
		// If the index is greater than or equal
		// to the size of the list, panic
		panic("Invalid index")
	} else {
		// Else, return the value at that index
		return l.data[index]
	}
}

// List - Make-shift constructor for FixedList
func List() *FixedList {
	// Allocate and return a new FixedList
	return new(FixedList)
}
