package binary

import (
	"cmp"
)

type bsNode[E cmp.Ordered] struct {
	value  E
	parent *bsNode[E]
	left   *bsNode[E]
	right  *bsNode[E]
}

func (n *bsNode[E]) Value() E {
	return n.value
}

func (n *bsNode[E]) Parent() (Node[E], bool) {
	return n.parent, n.parent != nil
}

func (n *bsNode[E]) Left() (Node[E], bool) {
	return n.left, n.left != nil
}

func (n *bsNode[E]) Right() (Node[E], bool) {
	return n.right, n.right != nil
}

// BSTree is a binary search tree.
type BSTree[E cmp.Ordered] struct {
	root *bsNode[E]
}

// NewBinarySearchTree creates a binary search tree.
func NewBinarySearchTree[E cmp.Ordered]() *BSTree[E] {
	return new(BSTree[E])
}

func (t *BSTree[E]) Insert(value E) bool {
	n := &bsNode[E]{
		value: value,
	}

	if t.root == nil {
		t.root = n
		return false
	}

	current := t.root
	for {
		x := cmp.Compare(value, current.value)
		if x < 0 {
			if current.left == nil {
				current.left = n
				n.parent = current
				break
			}
			current = current.left
		} else if x > 0 {
			if current.right == nil {
				current.right = n
				n.parent = current
				break
			}
			current = current.right
		} else {
			return false
		}
	}
	return true
}

func (t *BSTree[E]) Balance() {

}

// Find the "inorder successor starting from bsNode n.
// The inorder successor is "the smallest key that is greater than the input bsNode
func findInorderSuccessor[E cmp.Ordered](n *bsNode[E]) *bsNode[E] {
	if n.right != nil {
		n = n.right
		for {
			if n.left == nil {
				return n
			}

			if n.left != nil {
				n = n.left
			}
		}
	} else {
		p := n.parent
		for p != nil && n == p.right {
			n = p
			p = n.parent
		}
		return p
	}
}

// A common implementation for Find and Delete.  Returns the bsNode where value is found, or nil if it is not found.
func find[E cmp.Ordered](node *bsNode[E], value E) *bsNode[E] {
	for node != nil && node.value != value {
		if value < node.value {
			node = node.left
		} else if value > node.value {
			node = node.right
		}
	}
	return node
}

func (t *BSTree[E]) Delete(value E) bool {
	node := find(t.root, value)
	if node == nil {
		return false
	}

	if node.left == nil && node.right == nil { // case 1: leaf n
		if node.parent.value < node.value {
			node.parent.right = nil
		} else {
			node.parent.left = nil
		}
	} else if node.left != nil && node.right != nil { // case 2:  bsNode with two children
		// swap n with its inorder successor
		successor := findInorderSuccessor(node)

		// swap values and unlink the successor bsNode
		node.value = successor.value
		if successor.value < successor.parent.value {
			successor.parent.left = nil
		} else {
			successor.parent.right = nil
		}
		successor.parent = nil
	} else { // case 3: bsNode one child
		child := node.left
		if node.left == nil {
			child = node.right
		}

		if node == t.root {
			t.root = child
		} else {
			if node.value < node.parent.value {
				node.parent.left = child
			} else {
				node.parent.right = child
			}
		}
	}
	return true
}

func (t *BSTree[E]) Traverse(method func(node Node[E], v func(node Node[E]) bool) bool, v func(node Node[E]) bool) {
	if t.root != nil {
		method(t.root, v)
	}
}

// Find Returns true if the tree contains value.
func (t *BSTree[E]) Find(value E) bool {
	node := find(t.root, value)
	return node != nil
}
