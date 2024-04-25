package splay

import (
	"fmt"
	"testing"
)

func TestRotateLeft(t *testing.T) {
	tree := NewSplayTree[int]()
	tree.InsertAll(1, 2, 3)
	visitor := func(n *Node[int]) bool {
		fmt.Println(n.value)
		return true
	}
	tree.Visit(visitor)
}

func TestRotateRight(t *testing.T) {
	tree := NewSplayTree[int]()
	tree.InsertAll(1, 3, 2)
	visitor := func(n *Node[int]) bool {
		fmt.Println(n.value)
		return true
	}
	tree.Visit(visitor)
}

func TestRotation(t *testing.T) {
	tree := NewSplayTree[int]()
	tree.InsertAll(4, 3, 6, 5, 7, 1)
	visitor := func(n *Node[int]) bool {
		fmt.Println(n.value)
		return true
	}
	tree.Visit(visitor)
}

func TestZigZag(t *testing.T) {
	tree := NewSplayTree[int]()
	tree.InsertAll(25, 50, 75, 30)
	visitor := func(n *Node[int]) bool {
		fmt.Println(n.value)
		return true
	}
	tree.Visit(visitor)
}

func TestLargerTree(t *testing.T) {
	tree := NewSplayTree[int]()
	//
	tree.InsertAll(4, 5, 6, 2, 1, 20, 17, 22, 18)

	if !tree.Find(1) {
		t.Error("could not find 1 in splay tree")
	}

	if !tree.Find(1) {
		t.Error("could not find 1 in splay tree")
	}

	if !tree.Find(17) {
		t.Error("could not find 1 in splay tree")
	}

	// 17 should be the first node now
	calls := 0
	v := func(n *Node[int]) bool {
		calls++
		return n.value != 17
	}
	tree.Visit(v)
	if calls != 1 {
		t.Error("17 was not at the top of the tree")
	}
}

func TestFind(t *testing.T) {
	tree := NewSplayTree[int]()
	tree.InsertAll(25, 50, 75, 30)

	if !tree.Find(25) {
		t.Error("could not find element 25")
		return
	}

	calls := 0
	visitor := func(n *Node[int]) bool {
		calls++
		return n.value != 25
	}
	tree.Visit(visitor)

	if calls != 1 {
		t.Error("expected 1 call to find 25, got ", calls)
	}
}
