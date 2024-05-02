package splay

import (
	"cmp"
)

// Tree A splay tree where the most recently accessed node is rotated to the root. A splay tree does
// not have to be in strict balance.
type Tree[E cmp.Ordered] struct {
	root *Node[E]
}

type Node[E cmp.Ordered] struct {
	value  E
	parent *Node[E]
	left   *Node[E]
	right  *Node[E]
}

func NewSplayTree[E cmp.Ordered]() *Tree[E] {
	return &Tree[E]{}
}

func (t *Tree[E]) Insert(value E) {
	node := &Node[E]{
		value: value,
	}

	if t.root == nil {
		t.root = node
		return
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
		} else {
			if current.right == nil {
				current.right = node
				node.parent = current
				break
			}
			current = current.right
		}
	}

	// bring the newly inserted node to the root
	t.splay(node)
}

func (t *Tree[E]) InsertAll(values ...E) {
	for _, v := range values {
		t.Insert(v)
	}
}

// Delete Remove nodes from the splay tree.
// Based off the wikipedia description: https://en.wikipedia.org/wiki/Splay_tree#Deletion
func (t *Tree[E]) Delete(value E) {
	node := find(t.root, value)
	t.splay(node)

	if node.value != value {
		return
	}

	left := node.left
	right := node.right

	// unlink the node
	node.left = nil
	node.right = nil

	// connect the subtrees
	var smax *Node[E] = nil
	if left != nil {
		left.parent = nil
		smax = subtreeMax(left)
		t.splay(smax)
	}

	if right != nil {
		right.parent = nil
		if left != nil {
			smax.right = right
		} else {
			t.root = right
		}
		right.parent = smax
	}
}

func subtreeMax[E cmp.Ordered](n *Node[E]) *Node[E] {
	for n.right != nil {
		n = n.right
	}
	return n
}

// Find - Returns true if the tree contains value.  Note that the tree will splay on the node
// containing the value, and in the case the value isn't found, on the leaf node with the closest value.
func (t *Tree[E]) Find(value E) bool {
	node := find(t.root, value)
	t.splay(node)
	return node.value == value
}

// common implementation between find and delete
func find[E cmp.Ordered](node *Node[E], value E) *Node[E] {
	for node != nil && value != node.value {
		if value < node.value {
			if node.left == nil {
				break
			}
			node = node.left
		} else {
			if node.right == nil {
				break
			}
			node = node.right
		}
	}
	return node
}

// rotate the tree until n is the root node
func (t *Tree[E]) splay(n *Node[E]) {
	for n != t.root {
		if n.parent == nil {
			t.root = n
			break
		}
		haveGrandparent := n.parent != nil && n.parent.parent != nil
		if haveGrandparent {
			p := n.parent
			gp := n.parent.parent
			if n == p.left && p == gp.left || n == p.right && p == gp.right {
				t.zigzig(n)
			} else {
				t.zigzag(n)
			}
		} else { // zig
			if n.parent.right == n {
				t.rotateLeft(n)
			} else {
				t.rotateRight(n)
			}
		}
	}
}

func (t *Tree[E]) trinodeLeft(n, p, gp *Node[E]) {
	p.right = n.left
	if n.left != nil {
		n.left.parent = p
	}
	n.parent = gp.parent
	gp.parent = n
	gp.left = n.right
	if n.right != nil {
		n.right.parent = gp
	}
	p.parent = n

	if n.right != nil {
		n.right.parent = gp
	}
	n.right = gp
	n.left = p

	if n.parent == nil {
		t.root = n
	}
}

func (t *Tree[E]) trinodeRight(n, p, gp *Node[E]) {
	p.left = n.right
	if n.right != nil {
		n.right.parent = p
	}
	n.parent = gp.parent
	gp.parent = n
	gp.right = n.left
	if n.left != nil {
		n.left.parent = gp
	}
	p.parent = n

	if n.left != nil {
		n.left.parent = gp
	}
	n.left = gp
	n.right = p

	if n.parent == nil {
		t.root = n
	}
}

// Restructure a left child of a left child to a right child of a right child, or vise versa.
func (t *Tree[E]) zigzig(n *Node[E]) {
	p := n.parent
	gp := n.parent.parent

	if gp.parent != nil {
		if gp == gp.parent.left {
			gp.parent.left = n
		} else {
			gp.parent.right = n
		}
	}

	n.parent = gp.parent
	p.parent = n
	gp.parent = p

	if n == p.right {
		gp.right = p.left
		if p.left != nil {
			p.left.parent = gp
		}
		p.left = gp
		p.right = n.left
		if n.left != nil {
			n.left.parent = p
		}
		n.left = p
	} else {
		gp.left = p.right
		if p.right != nil {
			p.right.parent = gp
		}
		p.right = gp
		p.left = n.right
		if n.right != nil {
			n.right.parent = p
		}
		n.right = p
	}

	if n.parent == nil {
		t.root = n
	}
}

// Restructure a left child of a right child or vise versa.
func (t *Tree[E]) zigzag(n *Node[E]) {
	gp := n.parent.parent
	p := n.parent

	if gp.parent != nil {
		n.parent = gp.parent
		if gp.parent.left == gp {
			gp.parent.left = n
		} else {
			gp.parent.right = n
		}
	}

	if n == p.right {
		t.trinodeLeft(n, p, gp)
	} else {
		t.trinodeRight(n, p, gp)
	}
}

func (t *Tree[E]) rotateLeft(n *Node[E]) {
	p := n.parent
	n.parent = p.parent
	if p.parent != nil {
		if p.parent.left == p {
			p.parent.left = n
		} else {
			p.parent.right = n
		}
	}

	if n.parent == nil {
		t.root = n
	}

	p.right = n.left
	if p.right != nil {
		p.right.parent = p
	}
	p.parent = n
	n.left = p
}

func (t *Tree[E]) rotateRight(n *Node[E]) {
	p := n.parent
	n.parent = p.parent
	if p.parent != nil {
		if p.parent.left == p {
			p.parent.left = n
		} else {
			p.parent.right = n
		}
	}

	if n.parent == nil {
		t.root = n
	}

	p.left = n.right
	if n.right != nil {
		p.left.parent = p
	}
	p.parent = n
	n.right = p
}

type Visitor[E cmp.Ordered] func(n *Node[E]) bool

func (t *Tree[E]) Visit(v Visitor[E]) {
	traverse(t.root, v)
}

func traverse[E cmp.Ordered](n *Node[E], v Visitor[E]) bool {
	if n == nil {
		return true
	}

	if !v(n) {
		return false
	}

	if !traverse(n.left, v) {
		return false
	}

	if !traverse(n.right, v) {
		return false
	}

	return true
}
