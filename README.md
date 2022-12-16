# Binary-Tree Go

Implementation of a binary tree in Go. A Binary Tree is a tree data structure in which each node has at most two children, which are referred to as the left child and the right child. A node with no children is called a leaf node. A node cannot have more than two children.

## Structure

**Types:**
```go
// Type for storing a node of a binary tree
type Node struct {
	Value interface{}
	Left  *Node
	Right *Node
}

// Type for storing a binary tree
type GeneralTree struct {
	Root *Node
}
```

**Functions:**

- `func NewNode(value interface{}) *Node`: Create a new node
- `func NewGeneralTree() *GeneralTree`: Create a new binary tree
- `func NewGeneralTreeWithNumbers(values ...interface{}) *GeneralTree`: Create a new tree and insert the given values (in order, root, left and right), number of arguments in the function is not limited

**Methods:**

```go
// Insert a new node into the binary tree
func (tree *GeneralTree) Insert(value interface{})
func (node *Node) Insert(value interface{})

// Insert a new node into left side of the binary tree
func (tree *GeneralTree) InsertLeft(value interface{})
func (node *Node) InsertLeft(value interface{})

// Insert a new node into right side of the binary tree
func (tree *GeneralTree) InsertRight(value interface{})
func (node *Node) InsertRight(value interface{})

// Search a node in the binary tree
func (tree *GeneralTree) Search(value interface{}) *Node
func (node *Node) Search(value interface{}) *Node

// Find the first value in the binary tree
func (tree *GeneralTree) FindFirst() *Node
func (node *Node) FindFirst() *Node

// Find the last value in the binary tree
func (tree *GeneralTree) FindLast() *Node
func (node *Node) FindLast() *Node

	// Find the height of the binary tree
func (tree *GeneralTree) Height() int
func (node *Node) Height() int

	// Find the depth of the binary tree
func (tree *GeneralTree) Depth() int
func (node *Node) Depth() int

	// Find the number of nodes in the binary tree
func (tree *GeneralTree) Size() int
func (node *Node) Size() int

	// Find the number of leaves in the binary tree
func (tree *GeneralTree) Leaves() int
func (node *Node) Leaves() int

	// Find the number of full nodes in the binary tree
func (tree *GeneralTree) FullNodes() int
func (node *Node) FullNodes() int

// Find the number of half nodes in the binary tree
func (tree *GeneralTree) HalfNodes() int
func (node *Node) HalfNodes() int

// Find the number of nodes in the binary tree
func (tree *GeneralTree) Nodes() int
func (node *Node) Nodes() int

// Find the level of a node in the binary tree
func (tree *GeneralTree) Level(value interface{}) int
func (node *Node) Level(value interface{}) int

// Find the parent of a node in the binary tree
func (tree *GeneralTree) Parent(value interface{}) *Node
func (node *Node) Parent(value interface{}) *Node

// Find the sibling of a node in the binary tree
func (tree *GeneralTree) Sibling(value interface{}) *Node
func (node *Node) Sibling(value interface{}) *Node

// Find the uncle of a node in the binary tree
func (tree *GeneralTree) Uncle(value interface{}) *Node
func (node *Node) Uncle(value interface{}) *Node

// Find the number of nodes in the binary tree
func (tree *GeneralTree) Count() int
func (node *Node) Count() int

// Print the binary tree in pre-order
func (tree *GeneralTree) PreOrder()
func (node *Node) PreOrder()

// Print the binary tree in in-order
func (tree *GeneralTree) InOrder()
func (node *Node) InOrder()

// Print the binary tree in post-order
func (tree *GeneralTree) PostOrder()
func (node *Node) PostOrder()

// Print the binary tree in level-order
func (tree *GeneralTree) LevelOrder()
func (node *Node) LevelOrder()

// Print the binary tree in reverse level-order
func (tree *GeneralTree) ReverseLevelOrder()
func (node *Node) ReverseLevelOrder()

// Print the binary tree in spiral order
func (tree *GeneralTree) SpiralOrder()
func (node *Node) SpiralOrder()

// Print the binary tree in zig-zag order
func (tree *GeneralTree) ZigZagOrder()
func (node *Node) ZigZagOrder()

// Print the binary tree in vertical order
func (tree *GeneralTree) VerticalOrder()
func (node *Node) VerticalOrder()

// Print the binary tree in reverse-vertical order
func (tree *GeneralTree) ReverseVerticalOrder()
func (node *Node) ReverseVerticalOrder()

// Print the binary tree in diagonal order
func (tree *GeneralTree) DiagonalOrder()
func (node *Node) DiagonalOrder()

// Print the binary tree in reverse-diagonal order
func (tree *GeneralTree) ReverseDiagonalOrder()
func (node *Node) ReverseDiagonalOrder()
```

## License

This project is licensed under the GPL-3.0 License - see the [LICENSE](LICENSE) file for details.

© Copyright (c) 2022, Max Base
