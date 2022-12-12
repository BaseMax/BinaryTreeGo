# Binary-Tree Go

Implentation of a binary tree in Go.

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

TODO

## License

This project is licensed under the GPL-3.0 License - see the [LICENSE](LICENSE) file for details.

© Copyright (c) 2022, Max Base
