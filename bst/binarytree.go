package bst

import (
	"cmp"
)

type Node[E cmp.Ordered] struct {
	value  E
	parent *Node[E]
	left   *Node[E]
	right  *Node[E]
}

type Tree[E cmp.Ordered] struct {
	root *Node[E]
}

func NewTree[E cmp.Ordered]() *Tree[E] {
	return new(Tree[E])
}

func (t *Tree[E]) Insert(value E) {
	n := &Node[E]{
		value: value,
	}

	if t.root == nil {
		t.root = n
		return
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
			// inserting the same value is a no-op
			break
		}
	}
}

func (t *Tree[E]) InsertAll(values ...E) {
	for _, v := range values {
		t.Insert(v)
	}
}

func (t *Tree[E]) Balance() {

}

// Find the "inorder successor starting from node n.
// The inorder successor is "the smallest key that is greater than the input node
func findInorderSuccessor[E cmp.Ordered](n *Node[E]) *Node[E] {
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

func (t *Tree[E]) Delete(value E) {
	deleter := func(n *Node[E]) bool {
		if n.value != value {
			return true
		}

		if n.left == nil && n.right == nil { // case 1: leaf n
			if n.parent.value < n.value {
				n.parent.right = nil
			} else {
				n.parent.left = nil
			}
		} else if n.left != nil && n.right != nil { // case 2:  node with two children
			// swap n with its inorder successor
			successor := findInorderSuccessor(n)

			// swap values and unlink the successor node
			n.value = successor.value
			if successor.value < successor.parent.value {
				successor.parent.left = nil
			} else {
				successor.parent.right = nil
			}
			successor.parent = nil
		} else { // case 3: node one child
			child := n.left
			if n.left == nil {
				child = n.right
			}

			if n == t.root {
				t.root = child
			} else {
				if n.value < n.parent.value {
					n.parent.left = child
				} else {
					n.parent.right = child
				}
			}
		}

		return false
	}

	recurse(t.root, deleter)
}

func recurse[E cmp.Ordered](node *Node[E], v Visitor[E]) {
	if node == nil {
		return
	}

	if !v(node) {
		return
	}

	recurse(node.left, v)
	recurse(node.right, v)
}

// Visitor A function applied to each node in the tree.
// Implementations should return false when searching the tree is no longer necessary.
type Visitor[E cmp.Ordered] func(node *Node[E]) bool

// Visit Applies a Visitor function to each node in the tree.
func (t *Tree[E]) Visit(v Visitor[E]) {
	recurse(t.root, v)
}

// Find Returns true if the tree contains value.
func (t *Tree[E]) Find(value E) bool {
	found := false
	finder := func(n *Node[E]) bool {
		if n.value == value {
			found = true
			return false
		}
		return true
	}
	t.Visit(finder)
	return found
}
