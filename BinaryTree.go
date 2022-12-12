/**
 *
 * @file BinaryTree.go
 * @author Max Base (maxbasecode@gmail.com)
 * @brief Binary Tree Implementation in gO
 * @version 0.1
 * @date 2022-12-12
 * @copyright Copyright (c) 2022
 *
 */

package main

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

// Create a new node
func NewNode(value interface{}) *Node {
	return &Node{Value: value}
}

// Create a new binary tree
func NewGeneralTree() *GeneralTree {
	return &GeneralTree{}
}

// Insert a new node into the binary tree
func (tree *GeneralTree) Insert(value interface{}) {
	if tree.Root == nil {
		tree.Root = NewNode(value)
	} else {
		tree.Root.Insert(value)
	}
}

func (node *Node) Insert(value interface{}) {
	if node == nil {
		return
	} else if value.(int) <= node.Value.(int) {
		if node.Left == nil {
			node.Left = NewNode(value)
		} else {
			node.Left.Insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = NewNode(value)
		} else {
			node.Right.Insert(value)
		}
	}
}

// Insert a new node into left side of the binary tree
func (tree *GeneralTree) InsertLeft(value interface{}) {
	if tree.Root == nil {
		tree.Root = NewNode(value)
	} else {
		tree.Root.InsertLeft(value)
	}
}

func (node *Node) InsertLeft(value interface{}) {
	if node == nil {
		return
	} else if node.Left == nil {
		node.Left = NewNode(value)
	} else {
		node.Left.InsertLeft(value)
	}
}

// Insert a new node into right side of the binary tree
func (tree *GeneralTree) InsertRight(value interface{}) {
	if tree.Root == nil {
		tree.Root = NewNode(value)
	} else {
		tree.Root.InsertRight(value)
	}
}

func (node *Node) InsertRight(value interface{}) {
	if node == nil {
		return
	} else if node.Right == nil {
		node.Right = NewNode(value)
	} else {
		node.Right.InsertRight(value)
	}
}

// Search a node in the binary tree
func (tree *GeneralTree) Search(value interface{}) *Node {
	if tree.Root == nil {
		return nil
	} else {
		return tree.Root.Search(value)
	}
}

// Search a node in the binary tree
func (node *Node) Search(value interface{}) *Node {
	if node == nil {
		return nil
	} else if node.Value == value {
		return node
	} else if value.(int) <= node.Value.(int) {
		return node.Left.Search(value)
	} else {
		return node.Right.Search(value)
	}
}

// Find the first value in the binary tree
func (tree *GeneralTree) FindFirst() *Node {
	if tree.Root == nil {
		return nil
	} else {
		return tree.Root.FindFirst()
	}
}

func (node *Node) FindFirst() *Node {
	if node == nil {
		return nil
	} else if node.Left == nil {
		return node
	} else {
		return node.Left.FindFirst()
	}
}

// Find the last value in the binary tree
func (tree *GeneralTree) FindLast() *Node {
	if tree.Root == nil {
		return nil
	} else {
		return tree.Root.FindLast()
	}
}

func (node *Node) FindLast() *Node {
	if node == nil {
		return nil
	} else if node.Right == nil {
		return node
	} else {
		return node.Right.FindLast()
	}
}

// Find the depth of the binary tree
func (tree *GeneralTree) Depth() int {
	if tree.Root == nil {
		return 0
	} else {
		return tree.Root.Depth()
	}
}

func (node *Node) Depth() int {
	if node == nil {
		return 0
	} else {
		leftDepth := node.Left.Depth()
		rightDepth := node.Right.Depth()

		if leftDepth > rightDepth {
			return leftDepth + 1
		} else {
			return rightDepth + 1
		}
	}
}

// Find the number of nodes in the binary tree
func (tree *GeneralTree) Size() int {
	if tree.Root == nil {
		return 0
	} else {
		return tree.Root.Size()
	}
}

func (node *Node) Size() int {
	if node == nil {
		return 0
	} else {
		return node.Left.Size() + node.Right.Size() + 1
	}
}

// Find the number of leaves in the binary tree
func (tree *GeneralTree) Leaves() int {
	if tree.Root == nil {
		return 0
	} else {
		return tree.Root.Leaves()
	}
}

func (node *Node) Leaves() int {
	if node == nil {
		return 0
	} else if node.Left == nil && node.Right == nil {
		return 1
	} else {
		return node.Left.Leaves() + node.Right.Leaves()
	}
}

// Find the number of full nodes in the binary tree
func (tree *GeneralTree) FullNodes() int {
	if tree.Root == nil {
		return 0
	} else {
		return tree.Root.FullNodes()
	}
}

func (node *Node) FullNodes() int {
	if node == nil {
		return 0
	} else if node.Left != nil && node.Right != nil {
		return node.Left.FullNodes() + node.Right.FullNodes() + 1
	} else {
		return node.Left.FullNodes() + node.Right.FullNodes()
	}
}

// Find the number of half nodes in the binary tree
func (tree *GeneralTree) HalfNodes() int {
	if tree.Root == nil {
		return 0
	} else {
		return tree.Root.HalfNodes()
	}
}

func (node *Node) HalfNodes() int {
	if node == nil {
		return 0
	} else if (node.Left != nil && node.Right == nil) || (node.Left == nil && node.Right != nil) {
		return node.Left.HalfNodes() + node.Right.HalfNodes() + 1
	} else {
		return node.Left.HalfNodes() + node.Right.HalfNodes()
	}
}

// Find the number of nodes in the binary tree
func (tree *GeneralTree) Nodes() int {
	if tree.Root == nil {
		return 0
	} else {
		return tree.Root.Nodes()
	}
}

func (node *Node) Nodes() int {
	if node == nil {
		return 0
	} else {
		return node.Left.Nodes() + node.Right.Nodes() + 1
	}
}

// Find the level of a node in the binary tree
func (tree *GeneralTree) Level(value interface{}) int {
	if tree.Root == nil {
		return 0
	} else {
		return tree.Root.Level(value)
	}
}

func (node *Node) Level(value interface{}) int {
	if node == nil {
		return 0
	} else if node.Value == value {
		return 1
	} else {
		leftLevel := node.Left.Level(value)
		rightLevel := node.Right.Level(value)

		if leftLevel > rightLevel {
			return leftLevel + 1
		} else {
			return rightLevel + 1
		}
	}
}

// Find the parent of a node in the binary tree
func (tree *GeneralTree) Parent(value interface{}) *Node {
	if tree.Root == nil {
		return nil
	} else {
		return tree.Root.Parent(value)
	}
}

func (node *Node) Parent(value interface{}) *Node {
	if node == nil {
		return nil
	} else if node.Left != nil && node.Left.Value == value {
		return node
	} else if node.Right != nil && node.Right.Value == value {
		return node
	} else {
		leftParent := node.Left.Parent(value)
		rightParent := node.Right.Parent(value)

		if leftParent != nil {
			return leftParent
		} else {
			return rightParent
		}
	}
}

// Find the sibling of a node in the binary tree
func (tree *GeneralTree) Sibling(value interface{}) *Node {
	if tree.Root == nil {
		return nil
	} else {
		return tree.Root.Sibling(value)
	}
}

func (node *Node) Sibling(value interface{}) *Node {
	if node == nil {
		return nil
	} else if node.Left != nil && node.Left.Value == value {
		return node.Right
	} else if node.Right != nil && node.Right.Value == value {
		return node.Left
	} else {
		leftSibling := node.Left.Sibling(value)
		rightSibling := node.Right.Sibling(value)

		if leftSibling != nil {
			return leftSibling
		} else {
			return rightSibling
		}
	}
}

// Find the uncle of a node in the binary tree
func (tree *GeneralTree) Uncle(value interface{}) *Node {
	if tree.Root == nil {
		return nil
	} else {
		return tree.Root.Uncle(value)
	}
}

func (node *Node) Uncle(value interface{}) *Node {
	if node == nil {
		return nil
	} else {
		parent := node.Parent(value)
		if parent != nil {
			return parent.Sibling(value)
		} else {
			return nil
		}
	}
}

// Find the number of nodes in the binary tree
func (tree *GeneralTree) Count() int {
	if tree.Root == nil {
		return 0
	} else {
		return tree.Root.Count()
	}
}

func (node *Node) Count() int {
	if node == nil {
		return 0
	} else {
		return node.Left.Count() + node.Right.Count() + 1
	}
}
