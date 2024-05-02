package bst

import (
	"cmp"
)

// TraversalOrder Constants for defining how the tree is traversed.
type TraversalOrder uint8

const (
	PreOrder TraversalOrder = iota
	PostOrder
	InOrder
	BreadthFirst
)

func (t TraversalOrder) String() string {
	var s string
	switch t {
	case PreOrder:
		s = "PreOrder"
	case PostOrder:
		s = "PostOrder"
	case InOrder:
		s = "InOrder"
	case BreadthFirst:
		s = "BreadthFirst"
	}
	return s
}

type Node[E cmp.Ordered] struct {
	value  E
	parent *Node[E]
	left   *Node[E]
	right  *Node[E]
}

type Tree[E cmp.Ordered] struct {
	root *Node[E]
}

func New[E cmp.Ordered]() *Tree[E] {
	return new(Tree[E])
}

func (t *Tree[E]) Insert(value E) bool {
	n := &Node[E]{
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

// A common implementation for Find and Delete.  Returns the node where value is found, or nil if it is not found.
func find[E cmp.Ordered](node *Node[E], value E) *Node[E] {
	for node != nil && node.value != value {
		if value < node.value {
			node = node.left
		} else if value > node.value {
			node = node.right
		}
	}
	return node
}

func (t *Tree[E]) Delete(value E) bool {
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
	} else if node.left != nil && node.right != nil { // case 2:  node with two children
		// swap n with its inorder successor
		successor := findInorderSuccessor(node)

		// swap values and unlink the successor node
		node.value = successor.value
		if successor.value < successor.parent.value {
			successor.parent.left = nil
		} else {
			successor.parent.right = nil
		}
		successor.parent = nil
	} else { // case 3: node one child
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

func postOrderTraverse[E cmp.Ordered](node *Node[E], v Visitor[E]) bool {
	if node == nil {
		return true
	}

	if !postOrderTraverse(node.left, v) {
		return false
	}

	if !postOrderTraverse(node.right, v) {
		return false
	}
	return v(node)
}

func inOrderTraverse[E cmp.Ordered](node *Node[E], v Visitor[E]) bool {
	if node == nil {
		return true
	}

	if !inOrderTraverse(node.left, v) {
		return false
	}

	if !v(node) {
		return false
	}

	if !inOrderTraverse(node.right, v) {
		return false
	}

	return true
}

func preOrderTraverse[E cmp.Ordered](node *Node[E], v Visitor[E]) {
	if node == nil {
		return
	}

	if !v(node) {
		return
	}

	preOrderTraverse(node.left, v)
	preOrderTraverse(node.right, v)
}

func breadthFirstTraverse[E cmp.Ordered](node *Node[E], v Visitor[E]) {
	nodes := make([]*Node[E], 1)

	nodes[0] = node

	for {
		children := make([]*Node[E], 0)
		for _, n := range nodes {
			if !v(n) {
				return
			}
			if n.left != nil {
				children = append(children, n.left)
			}
			if n.right != nil {
				children = append(children, n.right)
			}
		}
		if len(children) == 0 {
			break
		}
		nodes = children
	}
}

// Visitor A function applied to each node in the tree.
// Implementations should return false when searching the tree is no longer necessary.
type Visitor[E cmp.Ordered] func(node *Node[E]) bool

// Visit Applies a Visitor function to each node in the tree.
func (t *Tree[E]) Visit(traversal TraversalOrder, v Visitor[E]) {
	switch traversal {
	case PreOrder:
		preOrderTraverse(t.root, v)
	case PostOrder:
		postOrderTraverse(t.root, v)
	case InOrder:
		inOrderTraverse(t.root, v)
	case BreadthFirst:
		breadthFirstTraverse(t.root, v)
	}
}

// Find Returns true if the tree contains value.
func (t *Tree[E]) Find(value E) bool {
	node := find(t.root, value)
	return node != nil
}
