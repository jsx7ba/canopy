package canopy

import (
	"cmp"
)

// SplayTree A splay tree where the most recently accessed bsNode is rotated to the root. A splay tree does
// not have to be in strict balance.
type SplayTree[E cmp.Ordered] struct {
	root *bsNode[E]
}

func NewSplayTree[E cmp.Ordered]() *SplayTree[E] {
	return &SplayTree[E]{}
}

func (t *SplayTree[E]) Insert(value E) bool {
	node := &bsNode[E]{
		value: value,
	}

	if t.root == nil {
		t.root = node
		return true
	}

	inserted := true
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
			inserted = false
			break
		}
	}

	t.splay(node) // bring the newly inserted bsNode to the root

	return inserted
}

// Delete Remove nodes from the splay tree.
// Based off the wikipedia description: https://en.wikipedia.org/wiki/Splay_tree#Deletion
func (t *SplayTree[E]) Delete(value E) bool {
	node := find(t.root, value)
	t.splay(node)

	if node.value != value {
		return false
	}

	left := node.left
	right := node.right

	// unlink the bsNode
	node.left = nil
	node.right = nil

	// connect the subtrees
	var smax *bsNode[E] = nil
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

	return true
}

func subtreeMax[E cmp.Ordered](n *bsNode[E]) *bsNode[E] {
	for n.right != nil {
		n = n.right
	}
	return n
}

// Find - Returns true if the tree contains value.  Note that the tree will splay on the bsNode
// containing the value, and in the case the value isn't found, on the leaf bsNode with the closest value.
func (t *SplayTree[E]) Find(value E) bool {
	node := splayFind(t.root, value)
	t.splay(node)
	return node.value == value
}

func (t *SplayTree[E]) Traverse(method func(node Node[E], v func(node Node[E]) bool) bool, v func(node Node[E]) bool) {
	if t.root != nil {
		method(t.root, v)
	}
}

// common implementation between find and delete
func splayFind[E cmp.Ordered](node *bsNode[E], value E) *bsNode[E] {
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

// rotate the tree until n is the root bsNode
func (t *SplayTree[E]) splay(n *bsNode[E]) {
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

func (t *SplayTree[E]) trinodeLeft(n, p, gp *bsNode[E]) {
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

func (t *SplayTree[E]) trinodeRight(n, p, gp *bsNode[E]) {
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
func (t *SplayTree[E]) zigzig(n *bsNode[E]) {
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
func (t *SplayTree[E]) zigzag(n *bsNode[E]) {
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

func (t *SplayTree[E]) rotateLeft(n *bsNode[E]) {
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

func (t *SplayTree[E]) rotateRight(n *bsNode[E]) {
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
