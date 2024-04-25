package splay

import (
	"cmp"
	"fmt"
)

// SplayTree where the most recently accessed node is rotated to the root. A splay tree does
// not have to be in strict balance.
type SplayTree[E cmp.Ordered] struct {
	root *Node[E]
}

type Node[E cmp.Ordered] struct {
	value  E
	parent *Node[E]
	left   *Node[E]
	right  *Node[E]
}

func NewSplayTree[E cmp.Ordered]() *SplayTree[E] {
	return &SplayTree[E]{}
}

func (t *SplayTree[E]) Insert(value E) {
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

func (t *SplayTree[E]) InsertAll(values ...E) {
	for i, v := range values {
		if i == 2 {
			fmt.Println("last")
		}
		t.Insert(v)
	}
}

func (t *SplayTree[E]) Delete(value E) {

}

func (t *SplayTree[E]) Find(value E) bool {
	var found *Node[E]
	finder := func(n *Node[E]) bool {
		if n.value == value {
			found = n
			return false
		}
		return true
	}
	traverse(t.root, finder)
	t.splay(found)
	return found != nil
}

// rotate the tree until n is the root node
func (t *SplayTree[E]) splay(n *Node[E]) {
	for n != t.root {
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

func (t *SplayTree[E]) trinodeLeft(n, p, gp *Node[E]) {
	p.right = n.left
	n.parent = gp.parent
	gp.parent = n
	gp.left = n.right
	p.parent = n

	n.right = gp
	n.left = p

	if n.parent == nil {
		t.root = n
	}
}

func (t *SplayTree[E]) trinodeRight(n, p, gp *Node[E]) {
	p.left = n.right
	n.parent = gp.parent
	gp.parent = n
	gp.right = n.left
	p.parent = n

	n.left = gp
	n.right = p

	if n.parent == nil {
		t.root = n
	}
}

func (t *SplayTree[E]) zigzig(n *Node[E]) {
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
		p.left = gp
		p.right = n.left
		n.left = p
	} else {
		gp.left = p.right
		p.right = gp
		p.left = n.right
		n.right = p
	}

	if n.parent == nil {
		t.root = n
	}
}

func (t *SplayTree[E]) zigzag(n *Node[E]) {
	gp := n.parent.parent
	p := n.parent

	if gp.parent != nil {
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

func (t *SplayTree[E]) rotateLeft(n *Node[E]) {
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

func (t *SplayTree[E]) rotateRight(n *Node[E]) {
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

func (t *SplayTree[E]) Visit(v Visitor[E]) {
	traverse(t.root, v)
}

func traverse[E cmp.Ordered](n *Node[E], v Visitor[E]) {
	if n == nil {
		return
	}

	if !v(n) {
		return
	}

	traverse(n.left, v)
	traverse(n.right, v)
}
