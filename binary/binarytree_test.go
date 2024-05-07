package binary

import (
	"fmt"
	"testing"
)

func TestDepthTraversal(t *testing.T) {
	tree := NewBinarySearchTree[int]()
	tree.Insert(1)
	tree.Insert(0)
	tree.Insert(2)

	v := func(n Node[int]) bool {
		fmt.Println(n.Value())
		return true
	}

	tree.Traverse(PreOrder[int], v)
}

func TestInOrderSuccessor(t *testing.T) {
	tree := NewBinarySearchTree[int]()
	InsertAll(tree, 20, 8, 22, 4, 12, 10, 14)
	type data struct {
		val  int
		succ int
	}

	testData := []data{{10, 12}, {14, 20}, {8, 10}}

	for _, d := range testData {
		startNode := find(tree.root, d.val)

		if startNode == nil {
			t.Error("no start bsNode found for ", d.val)
			continue
		}

		successor := findInorderSuccessor(startNode)
		if successor == nil {
			t.Error("no successor found for ", d.val)
			continue
		}
		if successor.value != d.succ {
			t.Error("expected ", d.succ, " got ", successor.value)
		}
	}
}

func TestDeleteBSLeaf(t *testing.T) {
	data := []int{1, 0, 2}
	expected := []int{1, 0}
	actual := make([]int, len(expected))

	tree := NewBinarySearchTree[int]()
	InsertAll(tree, data...)

	tree.Delete(2)
	index := 0
	v := func(n Node[int]) bool {
		fmt.Println(n.Value())
		actual[index] = n.Value()
		index++
		return true
	}

	tree.Traverse(PreOrder[int], v)
	arrayEquals(t, "", expected, actual)
}

func TestDeleteInternal(t *testing.T) {
	expected := []int{3, 0, 1, 2, 4, 5}
	data := []int{3, 0, 1, 2, 6, 4, 5}
	actual := make([]int, len(expected))
	index := 0

	tree := NewBinarySearchTree[int]()
	InsertAll(tree, data...)

	tree.Delete(6)
	v := func(n Node[int]) bool {
		fmt.Println(n.Value())
		actual[index] = n.Value()
		index++
		return true
	}

	tree.Traverse(PreOrder[int], v)
	arrayEquals(t, "", expected, actual)
}

func TestDeleteBSRoot(t *testing.T) {
	data := []int{20, 8, 22, 4, 12, 10, 14}
	expected := []int{22, 8, 4, 12, 10, 14}
	actual := make([]int, len(expected))
	index := 0

	tree := NewBinarySearchTree[int]()
	InsertAll(tree, data...)

	tree.Delete(20)
	v := func(n Node[int]) bool {
		fmt.Println(n.Value())
		actual[index] = n.Value()
		index++
		return true
	}

	tree.Traverse(PreOrder[int], v)
	arrayEquals(t, "", expected, actual)
}
