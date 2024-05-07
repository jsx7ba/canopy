package binary

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func TestTraversals(t *testing.T) {
	type traverseFunc func(node Node[int], v func(node Node[int]) bool) bool
	funcs := []traverseFunc{PreOrder[int], PostOrder[int], InOrder[int], BreadthFirst[int]}
	data := []int{100, 20, 200, 10, 30, 150, 300}
	expected := [][]int{
		{100, 20, 10, 30, 200, 150, 300},
		{10, 30, 20, 150, 300, 200, 100},
		{10, 20, 30, 100, 150, 200, 300},
		{100, 20, 200, 10, 30, 150, 300},
	}

	tree := NewBinarySearchTree[int]()
	InsertAll(tree, data...)

	actual := make([]int, 0, len(data))
	visitor := func(n Node[int]) bool {
		actual = append(actual, n.Value())
		return true
	}

	for i, e := range expected {
		tree.Traverse(funcs[i], visitor)
		errorPrefix := fmt.Sprintf("%s:", runtime.FuncForPC(reflect.ValueOf(funcs[i]).Pointer()).Name())
		arrayEquals(t, errorPrefix, e, actual)
		actual = actual[:0]
	}
}
