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

// Insert a new node into the binary tree
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

// Insert a new node into left side of the binary tree
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

// Insert a new node into right side of the binary tree
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
