package bst

import (
	"testing"
)

func TestPanic(t *testing.T) {

}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestEmpty(t *testing.T) {
	tree := new(BinaryTree)
	if tree.Empty() == false {
		t.Errorf("Expected empty == true, got %t", tree.Empty())
	}
	if tree.Size() != 0 {
		t.Errorf("Expected size == 0, got %d", tree.Size())
	}
}
