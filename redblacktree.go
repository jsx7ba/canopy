package canopy

import (
	"cmp"
	"fmt"
)

// RBTree a struct for red black trees.
// Red Black Trees have the following properties:
// 1) Every node is either red or black.
// 2) nil nodes are considered black
// 3) A red node cannot have a red child
// 4) The path from any given node goes through the same number of black nodes.
// 5) If a node has a single child, it must be a red child.
type RBTree[E cmp.Ordered] struct {
	root *rbNode[E]
}

type color uint8

const (
	black = iota
	red
)

func (c color) String() string {
	if c == black {
		return "black"
	}
	return "red"
}

type rbNode[E cmp.Ordered] struct {
	value  E
	parent *rbNode[E]
	left   *rbNode[E]
	right  *rbNode[E]
	color  color
}

func (n *rbNode[E]) Value() E {
	return n.value
}

func (n *rbNode[E]) p() (Node[E], bool) {
	return n.parent, n.parent != nil
}

func (n *rbNode[E]) l() (Node[E], bool) {
	return n.left, n.left != nil
}

func (n *rbNode[E]) r() (Node[E], bool) {
	return n.right, n.right != nil
}

type RedBlackTree[E cmp.Ordered] struct {
	root *rbNode[E]
}

// NewRedBlackTree creates a new red black tree.
func NewRedBlackTree[E cmp.Ordered]() *RedBlackTree[E] {
	return &RedBlackTree[E]{root: nil}
}

func (t *RedBlackTree[E]) Insert(value E) bool {
	node := &rbNode[E]{value: value, color: red}
	if t.root == nil {
		t.root = node
		node.color = black
		return true
	}

	current := t.root
	for {
		if value < current.value {
			if current.left == nil {
				current.left = node
				node.parent = current
				break
			}
			current = current.left
		} else if value > current.value {
			if current.right == nil {
				current.right = node
				node.parent = current
				break
			}
			current = current.right
		} else {
			return false
		}
	}

	t.balance(node)
	return true
}

func (t *RedBlackTree[E]) balance(n *rbNode[E]) {
	if n == t.root || n.parent == nil {
		return
	}

	for n != t.root && n.parent.color == red {
		p := n.parent
		gp := n.parent.parent

		if gp == nil {
			// ?? not sure
			break
		}

		u := gp.right
		if p == gp.right {
			u = gp.left
		}

		if u != nil && u.color == red { // Case 1: The parent color is red, and the uncle color is red
			recolor1(p, u, gp)
			fmt.Println("case 1 recolor")
		} else { // Case 2: the parent color is red and the uncle color is black (or nil)
			if n == p.right && p == gp.left { // Case 2: n, p and gp make a triangle - rotate around parent
				t.rotateLeft(p)
				fmt.Println("case 2 rotate left")
			} else if n == p.left && p == gp.right {
				t.rotateRight(p)
				fmt.Println("case 2 rotate right")
			} else if n == p.right && p == gp.right { // Case 3: n, p, and gp are in a line: rotate around grandparent
				t.rotateLeft(gp)
				recolor3(p, gp)
				fmt.Println("case 3 rotate left/recolor")
			} else if n == p.left && p == gp.left {
				t.rotateRight(gp)
				recolor3(p, gp)
				fmt.Println("case 3 rotate right/recolor")
			}
		}
		n = p
	}
	t.root.color = black
}

func recolor1[E cmp.Ordered](p, u, gp *rbNode[E]) {
	p.color = black
	u.color = black
	gp.color = red
}

func recolor3[E cmp.Ordered](p, gp *rbNode[E]) {
	p.color = black
	gp.color = red
}

func (t *RedBlackTree[E]) rotateLeft(n *rbNode[E]) {
	p := n.parent
	c := n.right
	c.parent = n.parent
	n.parent = c
	n.right = c.left
	c.left = n
	if n.right != nil {
		n.right.parent = n
	}

	if p != nil {
		if p.left == n {
			p.left = c
		} else {
			p.right = c
		}
	}

	if c.parent == nil {
		t.root = c
	}
}

func (t *RedBlackTree[E]) rotateRight(n *rbNode[E]) {
	p := n.parent

	c := n.left
	c.parent = n.parent
	n.parent = c
	n.left = c.right
	c.right = n
	if n.left != nil {
		n.left.parent = n
	}

	if p != nil {
		if p.left == n {
			p.left = c
		} else {
			p.right = c
		}
	}

	if c.parent == nil {
		t.root = c
	}
}

func rbfind[E cmp.Ordered](n *rbNode[E], value E) *rbNode[E] {
	for n != nil && n.value != value {
		if value < n.value {
			n = n.left
		} else if value > n.value {
			n = n.right
		}
	}
	return n
}

func (t *RedBlackTree[E]) Delete(value E) bool {
	node := rbfind(t.root, value)
	if node == nil {
		return false
	}
	// todo - rebalance
	return true
}

func (t *RedBlackTree[E]) Find(value E) bool {
	node := rbfind(t.root, value)
	return node != nil
}

func (t *RedBlackTree[E]) Traverse(method func(node Node[E], v func(node Node[E]) bool) bool, v func(node Node[E]) bool) {
	if t.root != nil {
		method(t.root, v)
	}
}
