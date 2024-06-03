package canopy

import (
	"cmp"
	"fmt"
	"testing"
)

func TestRedBlack_singleNode(t *testing.T) {
	tree := NewRedBlackTree[int]()
	if !tree.Insert(32) {
		t.Fatal("insert single node failed")
	}

	count := 0
	counter := func(node Node[int]) bool {
		count++
		return true
	}

	tree.Traverse(PreOrder[int], counter)
	if count != 1 {
		PrintTree[int](tree)
		t.Fatal("counted more than one node, or zero nodes")
	}
}

func TestRedBlack_Case3Right(t *testing.T) {
	tree := NewRedBlackTree[int]()
	tree.Insert(32)
	tree.Insert(42)
	tree.Insert(52)

	n := tree.root
	if fail, mesg := checkNode[int](n, 42, black); fail {
		t.Fatal(mesg)
	}
	if fail, mesg := checkNode[int](n.left, 32, red); fail {
		t.Fatal(mesg)
	}
	if fail, mesg := checkNode[int](n.right, 52, red); fail {
		t.Fatal(mesg)
	}
}

func TestRedBlack_Case2Left(t *testing.T) {

}

func checkNode[E cmp.Ordered](n *rbNode[E], value E, color color) (bool, string) {
	if n == nil {
		return true, "node is nil"
	}

	if n.value != value {
		return true, fmt.Sprintf("node has value %v expected %v", n.value, value)
	}

	if n.color != color {
		return true, fmt.Sprintf("node has color %s, expected %s", n.color, color)
	}

	return false, ""
}

func TestRedBlack_fourNodes(t *testing.T) {
	tree := NewRedBlackTree[int]()
	InsertAll(tree, 32, 42, 52, 49, 53, 54, 15, 17)
	printRBTree[int](tree)
}

// provide visual confirmation about node color
func printRBTree[E cmp.Ordered](tree *RedBlackTree[E]) {
	printer := func(n Node[E]) bool {
		node := n.(*rbNode[E])
		template := "%v(%s) L %v R %v\n"
		leftValue := "nil"
		rightValue := "nil"
		if node.left != nil {
			leftValue = fmt.Sprintf("%+v", node.left.value)
		}
		if node.right != nil {
			rightValue = fmt.Sprintf("%v", node.right.value)
		}

		fmt.Printf(template, node.value, node.color, leftValue, rightValue)
		return true
	}

	tree.Traverse(BreadthFirst[E], printer)
}
