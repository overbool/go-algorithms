package tree

import (
	"testing"
)

/**
				8
		5				10
	3						15
 */
func TestBinaryTree(t *testing.T) {
	cmp := func(a, b interface{}) int {
		x := a.(int)
		y := b.(int)

		if x == y {
			return 0
		} else if x > y {
			return 1
		} else {
			return -1
		}
	}

	tree := NewBSTree(cmp)
	tree.Insert(8)
	tree.Insert(5)
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(3)

	tree.InOrder()

	//tree.Delete(8)
	//tree.InOrder()
	tree.Delete(3)
	tree.InOrder()
}

