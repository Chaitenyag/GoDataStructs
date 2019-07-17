package Hash

import (
	"fmt"
	"hash"
)

const capacity = 15

type Node struct {
	key int
	nextNode *Node
}

type Table struct {
	ht map[int]*Node
	curSize int
}

func Hash(key, ts int) int {
	return key % ts
}

func Insert(hash *Table, key int) {
	pos := Hash(key, hash.curSize)
	element := Node{key: key, nextNode: hash.ht[pos]}
	hash.ht[pos] = &element
	hash.curSize++;
}