package avl

import (
	"fmt"
	"math"
)

type Node struct {
	left   *Node
	right  *Node
	height int32
	value  float64
}

type BinaryTree struct {
	root *Node
	size uint64
}

func (b *BinaryTree) Size() uint64 {
	return b.size
}

func (b *BinaryTree) Empty() bool {
	return b.size == 0
}

func (b *BinaryTree) height(n *Node) int32 {
	if n == nil {
		return -1
	} else {
		return n.height
	}
}

func (b *BinaryTree) rotateRight(prt **Node) {
	chd := (*prt).left
	(*prt).left = chd.right
	(*prt).height = int32(1.0 + math.Max(float64(b.height((*prt).left)), float64(b.height((*prt).right))))
	chd.right = *prt
	chd.height = int32(1.0 + math.Max(float64(b.height((chd).left)), float64(b.height((chd).right))))
	*prt = chd
}

func (b *BinaryTree) rotateLeft(prt **Node) {
	chd := (*prt).right
	(*prt).right = chd.left
	(*prt).height = int32(1.0 + math.Max(float64(b.height((*prt).left)), float64(b.height((*prt).right))))
	chd.left = *prt
	chd.height = int32(1.0 + math.Max(float64(b.height((chd).left)), float64(b.height((chd).right))))
	*prt = chd
}

func (b *BinaryTree) retrace(n **Node) {
	if (*n) == nil {
		return
	}

	if b.height((*n).left)-b.height((*n).right) > 1 {
		if b.height((*n).left.left) < b.height((*n).left.right) {
			b.rotateLeft(&(*n).left)
		}
		b.rotateRight(n)
	} else if b.height((*n).right)-b.height((*n).left) > 1 {
		if b.height((*n).right.right) < b.height((*n).right.left) {
			b.rotateRight(&(*n).right)
		}
		b.rotateLeft(n)
	}

	(*n).height = int32(1.0 + math.Max(float64(b.height((*n).left)), float64(b.height((*n).right))))
}

func (b *BinaryTree) insertHelper(n **Node, value float64) {
	if *n == nil {
		*n = &Node{left: nil, right: nil, value: value}
	} else if value < (*n).value {
		b.insertHelper(&((*n).left), value)
	} else if value > (*n).value {
		b.insertHelper(&((*n).right), value)
	} else {
		panic(fmt.Sprintf("Duplicate Key!"))
	}
	b.retrace(n)
}

func (b *BinaryTree) Insert(value float64) {
	b.insertHelper(&b.root, value)
	b.size++
}

func (b *BinaryTree) removeHelper(n **Node, value float64) {
	if *n == nil {
		panic("Invalid Key")
	} else if value < (*n).value {
		b.removeHelper(&((*n).left), value)
	} else if value > (*n).value {
		b.removeHelper(&((*n).right), value)
	} else {
		if ((*n).left) != nil && ((*n).right) != nil {
			(*n).value = b.min((*n).right).value
			b.removeHelper(&((*n).right), value)
		} else {
			if (*n).left != nil {
				*n = (*n).left
			} else {
				*n = (*n).right
			}
		}
	}
	b.retrace(n)
}

func (b *BinaryTree) Remove(value float64) {
	b.removeHelper(&b.root, value)
	b.size--
}

func (b *BinaryTree) Contains(value float64) bool {
	n := b.root
	for n != nil {
		if value == n.value {
			return true
		} else if value < n.value {
			n = n.left
		} else if value > n.value {
			n = n.right
		}
	}
	return false
}

func (b *BinaryTree) min(n *Node) *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (b *BinaryTree) Min() float64 {
	if b.root == nil {
		panic("No min of empty tree")
	}
	n := b.min(b.root)
	return n.value
}

func (b *BinaryTree) max(n *Node) *Node {
	for n.right != nil {
		n = n.right
	}
	return n
}

func (b *BinaryTree) Max() float64 {
	if b.root == nil {
		panic("No min of empty tree")
	}
	n := b.max(b.root)
	return n.value
}

func main() {
	b := new(BinaryTree)
	b.Insert(301.89)
	b.Insert(1.0)
	b.Insert(2.0)
	b.Insert(5.0)
	b.Insert(10.0)
	b.Insert(20.0)
	b.Insert(401.0)

	b.Remove(1.0)
	b.Remove(2.0)
	b.Remove(5.0)
	b.Insert(5.0)

	fmt.Printf("Binary Tree Size is empty: %d\n", b.height(b.root.right))
}
