package canopy

import (
	"cmp"
	"fmt"
)

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

	// Traverse provides a way to visit the nodes in a binary tree.
	Traverse(method func(node Node[E], v func(node Node[E]) bool) bool, visitor func(node Node[E]) bool)
}

// Node is a common interface for all binary tree nodes.
type Node[E cmp.Ordered] interface {
	Value() E
	p() (Node[E], bool)
	l() (Node[E], bool)
	r() (Node[E], bool)
}

// InsertAll Inserts a slice of values into a given tree.
func InsertAll[E cmp.Ordered](t Tree[E], values ...E) {
	for _, v := range values {
		t.Insert(v)
	}
}

// PostOrder recursively traverses a binary tree in post order.
func PostOrder[E cmp.Ordered](node Node[E], v func(node Node[E]) bool) bool {
	if left, ok := node.l(); ok {
		if !PostOrder(left, v) {
			return false
		}
	}

	if right, ok := node.r(); ok {
		if !PostOrder(right, v) {
			return false
		}
	}

	return v(node)
}

// InOrder recursively traverses a binary tree "in order".
func InOrder[E cmp.Ordered](node Node[E], v func(node Node[E]) bool) bool {

	if left, ok := node.l(); ok {
		if !InOrder(left, v) {
			return false
		}
	}

	if !v(node) {
		return false
	}

	if right, ok := node.r(); ok {
		if !InOrder(right, v) {
			return false
		}
	}

	return true
}

// PreOrder recursively traverses a binary tree with "pre order".
func PreOrder[E cmp.Ordered](node Node[E], v func(node Node[E]) bool) bool {

	if !v(node) {
		return false
	}

	if left, ok := node.l(); ok {
		if !PreOrder(left, v) {
			return false
		}
	}

	if right, ok := node.r(); ok {
		if !PreOrder(right, v) {
			return false
		}
	}

	return true
}

// BreadthFirst traverses a binary tree with breadth first ordering.
func BreadthFirst[E cmp.Ordered](node Node[E], v func(node Node[E]) bool) bool {
	nodes := make([]Node[E], 1)
	nodes[0] = node

	for {
		children := make([]Node[E], 0)
		for _, n := range nodes {
			if !v(n) {
				return true
			}

			if left, ok := n.l(); ok {
				children = append(children, left)
			}
			if right, ok := n.r(); ok {
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

// PrintTree prints a binary tree in pre-order.
func PrintTree[E cmp.Ordered](tree Tree[E]) {
	visitor := func(n Node[E]) bool {
		fmt.Println(n.Value())
		return true
	}
	tree.Traverse(PreOrder[E], visitor)
}
