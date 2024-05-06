package binary

import (
	"fmt"
	"testing"
)

func TestTraversal(t *testing.T) {
	tree := NewBinarySearchTree[int]()
	tree.Insert(42)

	foo := func(node Node[int]) bool {
		fmt.Println(node.Value())
		return true
	}

	tree.Traverse(PostOrder[int], foo)
}
