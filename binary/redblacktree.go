package binary

import "cmp"

type RBTree[E cmp.Ordered] struct {
	root *RBNode[E]
}

type Color uint8

const (
	Black = iota
	Red
)

type RBNode[E cmp.Ordered] struct {
	value  E
	parent *RBNode[E]
	left   *RBNode[E]
	right  *RBNode[E]
	color  Color
}

func (n *RBNode[E]) Value() E {
	return n.value
}

func (n *RBNode[E]) Parent() *RBNode[E] {
	return n.parent
}

func (n *RBNode[E]) Left() *RBNode[E] {
	return n.left
}

func (n *RBNode[E]) Right() *RBNode[E] {
	return n.right
}

type RedBlackTree[E cmp.Ordered] struct {
	root *RBNode[E]
}

func NewRedBlackTree[E cmp.Ordered]() *RedBlackTree[E] {
	return &RedBlackTree[E]{root: nil}
}
