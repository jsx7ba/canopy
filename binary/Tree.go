package binary

import "cmp"

// Tree The interface for all tree types.
type Tree[E cmp.Ordered] interface {
	// Insert Places a value into the tree.
	// Returns true if the value was inserted, false if the value exists already.
	Insert(value E) bool

	// Delete Removes a value into the tree.
	// Returns true if the value was removed.
	Delete(value E) bool

	// Find Returns true if a value exists in the tree.
	Find(value E) bool

	Traverse(method func(node Node[E], v func(node Node[E]) bool) bool, visitor func(node Node[E]) bool)
}

type Node[E cmp.Ordered] interface {
	Value() E
	Parent() (Node[E], bool)
	Left() (Node[E], bool)
	Right() (Node[E], bool)
}

// InsertAll Inserts a slice of values into a given tree.
func InsertAll[E cmp.Ordered](t Tree[E], values ...E) {
	for _, v := range values {
		t.Insert(v)
	}
}

func PostOrder[E cmp.Ordered](node Node[E], v func(node Node[E]) bool) bool {
	if left, ok := node.Left(); ok {
		if !PostOrder(left, v) {
			return false
		}
	}

	if right, ok := node.Left(); ok {
		if !PostOrder(right, v) {
			return false
		}
	}

	return v(node)
}

func InOrder[E cmp.Ordered](node Node[E], v func(node Node[E]) bool) bool {

	if left, ok := node.Left(); ok {
		if !InOrder(left, v) {
			return false
		}
	}

	if !v(node) {
		return false
	}

	if right, ok := node.Right(); ok {
		if !InOrder(right, v) {
			return false
		}
	}

	return true
}

func PreOrder[E cmp.Ordered](node Node[E], v func(node Node[E]) bool) bool {

	if !v(node) {
		return false
	}

	if left, ok := node.Left(); ok {
		if !PreOrder(left, v) {
			return false
		}
	}

	if right, ok := node.Right(); ok {
		if !PreOrder(right, v) {
			return false
		}
	}

	return true
}

func BreadthFirst[E cmp.Ordered](node Node[E], v func(node Node[E]) bool) bool {
	nodes := make([]Node[E], 1)
	nodes[0] = node

	for {
		children := make([]Node[E], 0)
		for _, n := range nodes {
			if !v(n) {
				return true
			}

			if left, ok := n.Left(); ok {
				children = append(children, left)
			}
			if right, ok := n.Right(); ok {
				children = append(children, right)
			}
		}
		if len(children) == 0 {
			break
		}
		nodes = children
	}
	return true
}
