package tree

import "fmt"

type node struct {
	data  interface{}
	left  *node
	right *node
}

type compareFunc func(interface{}, interface{}) int

type BSTree struct {
	root    *node
	cmpFunc compareFunc
}

func NewBSTree(cmpFunc compareFunc) *BSTree {
	return &BSTree{cmpFunc: cmpFunc}
}

func (bst *BSTree) Insert(v interface{}) bool {
	if bst.root == nil {
		bst.root = &node{data: v}
		return true
	}

	p := bst.root
	for p != nil {
		res := bst.cmpFunc(v, p.data)
		if res == 0 {
			return false
		} else if res > 0 {
			if p.right == nil {
				p.right = &node{data: v}
				break
			}
			p = p.right
		} else {
			if p.left == nil {
				p.left = &node{data: v}
				break
			}

			p = p.left
		}
	}

	return true
}

func (bst *BSTree) Delete(v interface{}) bool {
	var pp *node
	p := bst.root
	deleteLeft := false
	for p != nil {
		res := bst.cmpFunc(v, p.data)
		if res > 0 {
			pp = p
			p = p.right
			deleteLeft = false
		} else if res < 0 {
			pp = p
			p = p.left
			deleteLeft = true
		} else {
			break
		}
	}

	// doesn't exist specific node
	if p == nil {
		return false
	}

	// delete the node that has no left and right child node
	if p.left == nil && p.right == nil {
		if pp != nil {
			if deleteLeft {
				pp.left = nil
			} else {
				pp.right = nil
			}
		} else {
			bst.root = nil
		}
	} else if p.right != nil {
		// delete the node that must has right child node
		qp := p
		q := p.right

		empty := true
		for q.left != nil {
			qp = p
			q = q.left
			empty = false
		}

		if empty {
			qp.right = q.right
		} else {
			qp.left = nil
		}

		q.left = p.left
		q.right = p.right

		if pp == nil {
			bst.root = q
		} else {
			if deleteLeft {
				pp.left = q
			} else {
				pp.right = q
			}
		}
	} else {
		// delete the node that has only left child node
		if pp != nil {
			if deleteLeft {
				pp.left = p.left
			} else {
				pp.right = p.left
			}
		} else {
			bst.root = p.left
		}
	}

	return true
}

func (bst *BSTree) PreOrder() {
	preOrder(bst.root)
}

func (bst *BSTree) InOrder() {
	inOrder(bst.root)
}

func (bst *BSTree) PostOrder() {
	postOrder(bst.root)
}

func (bst *BSTree) LevelOrder() {
	levelOrder(bst.root)
}

func (bst *BSTree) Depth() int {
	return calculate(bst.root, 0)
}

func preOrder(n *node) {
	if n == nil {
		return
	}

	fmt.Println(n.data)
	preOrder(n.left)
	preOrder(n.right)
}

func inOrder(n *node) {
	if n == nil {
		return
	}

	inOrder(n.left)
	fmt.Println(n.data)
	inOrder(n.right)
}

func postOrder(n *node) {
	if n == nil {
		return
	}

	postOrder(n.left)
	postOrder(n.right)
	fmt.Println(n.data)
}

func levelOrder(root *node) {
	q := make([]*node, 0)
	q = append(q, root)

	for len(q) != 0 {
		n := q[0]
		q = q[1:]

		fmt.Println(n.data)
		if n.left != nil {
			q = append(q, n.left)
		}

		if n.right != nil {
			q = append(q, n.right)
		}
	}
}

func calculate(n *node, depth int) int {
	if n == nil {
		return depth
	}

	return max(calculate(n.left, depth+1), calculate(n.right, depth+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
