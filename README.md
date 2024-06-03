# Canopy

A collection of tree data structures.  This is a repo for learning how structure go projects.

## Package binary

Contains a collection of binary trees.

Example usage:

```go
tree := binary.NewBinarySearchTree[int]()
binary.InsertAll(tree, 3, 2, 4)
printer := func(n binary.Node[int]) bool {
    fmt.Println("node value: ", n.Value())
    return true
}
tree.Traverse(binary.InOrder[int], printer)
```

### Binary Search Tree

A standard binary search tree.

#### Resources
* https://www.cs.usfca.edu/~galles/visualization/SplayTree.html

### Splay Tree
A Splay Tree is a binary tree which has the following properties:
- Binary tree
- Not necessarily balanced
- Any node with the most recent access is brought to the root

#### Resources
* https://www.cs.usfca.edu/~galles/visualization/SplayTree.html
* Adam Gaweda's [Splay Tree Lectures](https://youtube.com/playlist?list=PLK7dyt8j81q2QUEKr-38V0M8XdGQnAaKr&si=XKuHiiBSI_vT-YuI)

### Red Black Tree
A self-balancing binary tree.

#### Resources
https://www.cs.usfca.edu/~galles/visualization/RedBlack.html
* Michael Sambol's videos on [Red-Black Trees](https://www.youtube.com/playlist?list=PL9xmBV_5YoZNqDI8qfOZgzbqahCUmUEin)




