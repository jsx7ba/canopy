package canopy

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
	tree := NewSplayTree[int]()
	InsertAll(tree, 1, 2, 3)
	visitor := func(n Node[int]) bool {
		fmt.Println(n.Value())
		return true
	}
	tree.Traverse(InOrder[int], visitor)
}

func TestRotateRight(t *testing.T) {
	tree := NewSplayTree[int]()
	InsertAll(tree, 1, 2, 3)
	visitor := func(n Node[int]) bool {
		fmt.Println(n.Value())
		return true
	}
	tree.Traverse(InOrder[int], visitor)
}

func TestRotation(t *testing.T) {
	tree := NewSplayTree[int]()
	InsertAll(tree, 4, 3, 6, 5, 7, 1)
	PrintTree[int](tree)
}

func TestZigZag(t *testing.T) {
	tree := NewSplayTree[int]()
	InsertAll(tree, 25, 50, 75, 30)
	visitor := func(n Node[int]) bool {
		fmt.Println(n.Value())
		return true
	}
	tree.Traverse(PreOrder[int], visitor)
}

func TestLargerTree(t *testing.T) {
	tree := NewSplayTree[int]()
	//
	InsertAll(tree, 4, 5, 6, 2, 1, 20, 17, 22, 18)
	PrintTree[int](tree)

	if !tree.Find(1) {
		t.Error("could not find 1 in splay tree")
	}

	if !tree.Find(1) {
		t.Error("could not find 1 in splay tree")
	}

	if !tree.Find(17) {
		PrintTree[int](tree)
		t.Error("could not find 17 in splay tree")
	}

	// 17 should be the first bsNode now
	calls := 0
	v := func(n Node[int]) bool {
		calls++
		return n.Value() != 17
	}
	tree.Traverse(PreOrder[int], v)
	if calls != 1 {
		PrintTree[int](tree)
		t.Error("17 was not at the top of the tree")
	}
}

func TestFind(t *testing.T) {
	tree := NewSplayTree[int]()
	InsertAll(tree, 25, 50, 75, 30)

	if !tree.Find(30) {
		t.Error("could not find element 25")
		return
	}

	calls := 0
	visitor := func(n Node[int]) bool {
		calls++
		return n.Value() != 30
	}
	tree.Traverse(PreOrder[int], visitor)

	if calls != 1 {
		PrintTree[int](tree)
		t.Error("expected 1 call to find 25, got ", calls)
	}
}

func TestNotFound(t *testing.T) {
	expected := []int{75, 50, 30, 25}
	tree := NewSplayTree[int]()
	InsertAll(tree, 30, 25, 75, 50)

	if tree.Find(200) {
		t.Error("200 shouldn't exist in the tree")
	}

	actual := make([]int, 0)
	visitor := func(n Node[int]) bool {
		actual = append(actual, n.Value())
		return true
	}
	tree.Traverse(PreOrder[int], visitor)

	arrayEquals(t, "", expected, actual)
}

func TestDeleteRoot(t *testing.T) {
	tree := NewSplayTree[int]()
	tree.Insert(42)
	tree.Delete(42)
}

func TestDeleteChild1(t *testing.T) {
	tree := NewSplayTree[int]()
	InsertAll(tree, 42, 44, 32)
	tree.Delete(42)
}

func TestDeleteLeaf(t *testing.T) {
	tree := NewSplayTree[int]()
	InsertAll(tree, 42, 44, 32)
	tree.Delete(44)
}

func TestDeleteInLargerTree(t *testing.T) {
	tree := NewSplayTree[int]()
	InsertAll(tree, 4, 5, 6, 2, 1, 20)
	tree.Delete(2)
	PrintTree[int](tree)
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
	tree := NewSplayTree[int]()
	InsertAll(tree, values...)
	tree.Delete(3)
	PrintTree[int](tree)
}
