package bst

import (
	"cmp"
	"fmt"
	"testing"
)

func TestDepthTraversal(t *testing.T) {
	tree := NewBinarySearchTree[int]()
	tree.Insert(1)
	tree.Insert(0)
	tree.Insert(2)

	v := func(n *Node[int]) bool {
		fmt.Println(n.value)
		return true
	}

	tree.Visit(PreOrder, v)
}

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

func TestFind(t *testing.T) {

}

func TestInOrderSuccessor(t *testing.T) {
	tree := NewBinarySearchTree[int]()
	tree.InsertAll(20, 8, 22, 4, 12, 10, 14)
	type data struct {
		val  int
		succ int
	}

	testData := []data{{10, 12}, {14, 20}, {8, 10}}

	for _, d := range testData {
		var startNode *Node[int]

		findStartNode := func(n *Node[int]) bool {
			if n.value == d.val {
				startNode = n
				return false
			}
			return true
		}

		tree.Visit(PreOrder, findStartNode)

		if startNode == nil {
			t.Error("no start node found for ", d.val)
			continue
		}

		successor := findInorderSuccessor(startNode)
		if successor == nil {
			t.Error("no successor found for ", d.val)
			continue
		}
		if successor.value != d.succ {
			t.Error("expected ", d.succ, " got ", successor.value, " starting from ", startNode.value)
		}
	}
}

func TestDeleteLeaf(t *testing.T) {
	data := []int{1, 0, 2}
	expected := []int{1, 0}
	actual := make([]int, len(expected))

	tree := NewBinarySearchTree[int]()
	tree.InsertAll(data...)

	tree.Delete(2)
	index := 0
	v := func(n *Node[int]) bool {
		fmt.Println(n.value)
		actual[index] = n.value
		index++
		return true
	}

	tree.Visit(PreOrder, v)
	arrayEquals(t, "", expected, actual)
}

func TestDeleteInternal(t *testing.T) {
	expected := []int{3, 0, 1, 2, 4, 5}
	data := []int{3, 0, 1, 2, 6, 4, 5}
	actual := make([]int, len(expected))
	index := 0

	tree := NewBinarySearchTree[int]()
	tree.InsertAll(data...)

	tree.Delete(6)
	v := func(n *Node[int]) bool {
		fmt.Println(n.value)
		actual[index] = n.value
		index++
		return true
	}

	tree.Visit(PreOrder, v)
	arrayEquals(t, "", expected, actual)
}

func TestDeleteRoot(t *testing.T) {
	data := []int{20, 8, 22, 4, 12, 10, 14}
	expected := []int{22, 8, 4, 12, 10, 14}
	actual := make([]int, len(expected))
	index := 0

	tree := NewBinarySearchTree[int]()
	tree.InsertAll(data...)

	tree.Delete(20)
	v := func(n *Node[int]) bool {
		fmt.Println(n.value)
		actual[index] = n.value
		index++
		return true
	}

	tree.Visit(PreOrder, v)
	arrayEquals(t, "", expected, actual)
}

func TestTraversal(t *testing.T) {
	data := []int{100, 20, 200, 10, 30, 150, 300}
	expected := map[TraversalOrder][]int{
		PreOrder:     {100, 20, 10, 30, 200, 150, 300},
		PostOrder:    {10, 30, 20, 150, 300, 200, 100},
		InOrder:      {10, 20, 30, 100, 150, 200, 300},
		BreadthFirst: {100, 20, 200, 10, 30, 150, 300},
	}

	tree := NewBinarySearchTree[int]()
	tree.InsertAll(data...)

	actual := make([]int, 0, len(data))
	visitor := func(n *Node[int]) bool {
		actual = append(actual, n.value)
		return true
	}

	for k, v := range expected {
		tree.Visit(k, visitor)
		errorPrefix := fmt.Sprintf("%s:", k)
		arrayEquals(t, errorPrefix, v, actual)
		actual = actual[:0]
	}
}
