# Canopy

A collection of binary tree data structures.

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

```go
tree := canopy.NewBinarySearchTree[int]()
```

### Splay Tree
A Splay Tree is a binary tree which has the following properties:
- Binary tree
- Not necessarily balanced
- Any node with the most recent access is brought to the root

```go
tree := canopy.NewSplayTree[int]()
```

#### Resources
* https://www.cs.usfca.edu/~galles/visualization/SplayTree.html
* Adam Gaweda's [Splay Tree Lectures](https://youtube.com/playlist?list=PLK7dyt8j81q2QUEKr-38V0M8XdGQnAaKr&si=XKuHiiBSI_vT-YuI)

### Red Black Tree
A self-balancing binary tree where each node is colored red or black.

```go
tree := canopy.NewRedBlackTree[int]()
```

#### Resources
https://www.cs.usfca.edu/~galles/visualization/RedBlack.html
* Michael Sambol's videos on [Red-Black Trees](https://www.youtube.com/playlist?list=PL9xmBV_5YoZNqDI8qfOZgzbqahCUmUEin)




