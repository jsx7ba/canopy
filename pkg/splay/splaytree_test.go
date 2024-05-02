package splay

import (
	"cmp"
	"fmt"
	"testing"
)

func arrayEquals[E cmp.Ordered](t *testing.T, prefix string, expected, actual []E) bool {
	if len(expected) != len(actual) {
		t.Error(prefix, " actual array length does not match expected array")
		return false
	}

	for i := range expected {
		if expected[i] != actual[i] {
			t.Error(prefix, "expected", expected[i], "got", actual[i], "at index", i)
			return false
		}
	}

	return true
}

func TestRotateLeft(t *testing.T) {
	tree := New[int]()
	tree.InsertAll(1, 2, 3)
	visitor := func(n *Node[int]) bool {
		fmt.Println(n.value)
		return true
	}
	tree.Visit(visitor)
}

func TestRotateRight(t *testing.T) {
	tree := New[int]()
	tree.InsertAll(1, 3, 2)
	visitor := func(n *Node[int]) bool {
		fmt.Println(n.value)
		return true
	}
	tree.Visit(visitor)
}

func TestRotation(t *testing.T) {
	tree := New[int]()
	tree.InsertAll(4, 3, 6, 5, 7, 1)
	visitor := func(n *Node[int]) bool {
		fmt.Println(n.value)
		return true
	}
	tree.Visit(visitor)
}

func TestZigZag(t *testing.T) {
	tree := New[int]()
	tree.InsertAll(25, 50, 75, 30)
	visitor := func(n *Node[int]) bool {
		fmt.Println(n.value)
		return true
	}
	tree.Visit(visitor)
}

func TestLargerTree(t *testing.T) {
	tree := New[int]()
	//
	tree.InsertAll(4, 5, 6, 2, 1, 20, 17, 22, 18)
	PrintTree(tree)

	if !tree.Find(1) {
		t.Error("could not find 1 in splay tree")
	}

	if !tree.Find(1) {
		t.Error("could not find 1 in splay tree")
	}

	if !tree.Find(17) {
		PrintTree(tree)
		t.Error("could not find 17 in splay tree")
	}

	// 17 should be the first node now
	calls := 0
	v := func(n *Node[int]) bool {
		calls++
		return n.value != 17
	}
	tree.Visit(v)
	if calls != 1 {
		PrintTree(tree)
		t.Error("17 was not at the top of the tree")
	}
}

func TestFind(t *testing.T) {
	tree := New[int]()
	tree.InsertAll(25, 50, 75, 30)

	if !tree.Find(30) {
		t.Error("could not find element 25")
		return
	}

	calls := 0
	visitor := func(n *Node[int]) bool {
		calls++
		return n.value != 30
	}
	tree.Visit(visitor)

	if calls != 1 {
		t.Error("expected 1 call to find 25, got ", calls)
	}
}

func TestNotFound(t *testing.T) {
	expected := []int{75, 50, 30, 25}
	tree := New[int]()
	tree.InsertAll(30, 25, 75, 50)

	if tree.Find(200) {
		t.Error("200 shouldn't exist in the tree")
	}

	actual := make([]int, 0)
	visitor := func(n *Node[int]) bool {
		actual = append(actual, n.value)
		return true
	}
	tree.Visit(visitor)

	arrayEquals(t, "", expected, actual)
}

func TestDeleteRoot(t *testing.T) {
	tree := New[int]()
	tree.Insert(42)
	tree.Delete(42)
}

func TestDeleteChild1(t *testing.T) {
	tree := New[int]()
	tree.InsertAll(42, 44, 32)
	tree.Delete(42)
}

func TestDeleteLeaf(t *testing.T) {
	tree := New[int]()
	tree.InsertAll(42, 44, 32)
	tree.Delete(44)
}

func TestDeleteInLargerTree(t *testing.T) {
	tree := New[int]()
	tree.InsertAll(4, 5, 6, 2, 1, 20)
	tree.Delete(2)
	PrintTree(tree)
}

func TestDeleteInBigTree(t *testing.T) {
	values := []int{
		3, 36, 93, 61, 23, 83, 6, 25, 13, 66,
		39, 63, 30, 20, 19, 21, 78, 72, 46, 40,
		92, 84, 47, 24, 58, 89, 96, 26, 53, 98,
		9, 10, 45, 11, 79, 55, 42, 90, 37, 17,
		86, 12, 76, 28, 65, 99, 70, 44, 100,
		29, 43, 87, 56, 51, 95, 7, 5, 50,
	}
	tree := New[int]()
	tree.InsertAll(values...)
	tree.Delete(3)
	PrintTree(tree)
}

func PrintTree[E cmp.Ordered](tree *Tree[E]) {
	visitor := func(n *Node[E]) bool {
		fmt.Println(n.value)
		return true
	}
	tree.Visit(visitor)
}
