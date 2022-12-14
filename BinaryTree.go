/**
 *
 * @file BinaryTree.go
 * @author Max Base (maxbasecode@gmail.com)
 * @brief Binary Tree Implementation in Go
 * @version 0.1
 * @date 2022-12-12
 * @copyright Copyright (c) 2022
 *
 */

package main

import "fmt"

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

// Create a new tree and insert the given values (in order, root, left and right), number of arguments in the function is not limited
func NewGeneralTreeWithNumbers(values ...interface{}) *GeneralTree {
	tree := NewGeneralTree()
	for _, value := range values {
		tree.Insert(value)
	}
	return tree
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

// Find the height of the binary tree
func (tree *GeneralTree) Height() int {
	if tree.Root == nil {
		return 0
	} else {
		return tree.Root.Height()
	}
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	} else {
		leftHeight := node.Left.Height()
		rightHeight := node.Right.Height()

		if leftHeight > rightHeight {
			return leftHeight + 1
		} else {
			return rightHeight + 1
		}
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

// Print the binary tree in pre-order
func (tree *GeneralTree) PreOrder() {
	if tree.Root == nil {
		return
	} else {
		tree.Root.PreOrder()
	}
}

func (node *Node) PreOrder() {
	if node == nil {
		return
	} else {
		fmt.Println(node.Value)
		node.Left.PreOrder()
		node.Right.PreOrder()
	}
}

// Print the binary tree in in-order
func (tree *GeneralTree) InOrder() {
	if tree.Root == nil {
		return
	} else {
		tree.Root.InOrder()
	}
}

func (node *Node) InOrder() {
	if node == nil {
		return
	} else {
		node.Left.InOrder()
		fmt.Println(node.Value)
		node.Right.InOrder()
	}
}

// Print the binary tree in post-order
func (tree *GeneralTree) PostOrder() {
	if tree.Root == nil {
		return
	} else {
		tree.Root.PostOrder()
	}
}

func (node *Node) PostOrder() {
	if node == nil {
		return
	} else {
		node.Left.PostOrder()
		node.Right.PostOrder()
		fmt.Println(node.Value)
	}
}

// Print the binary tree in level-order
func (tree *GeneralTree) LevelOrder() {
	if tree.Root == nil {
		return
	} else {
		tree.Root.LevelOrder()
	}
}

func (node *Node) LevelOrder() {
	if node == nil {
		return
	} else {
		queue := []Node{*node}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			fmt.Println(current.Value)
			if current.Left != nil {
				queue = append(queue, *current.Left)
			}
			if current.Right != nil {
				queue = append(queue, *current.Right)
			}
		}
	}
}

// Print the binary tree in reverse level-order
func (tree *GeneralTree) ReverseLevelOrder() {
	if tree.Root == nil {
		return
	} else {
		tree.Root.ReverseLevelOrder()
	}
}

func (node *Node) ReverseLevelOrder() {
	if node == nil {
		return
	} else {
		queue := []Node{*node}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			fmt.Println(current.Value)
			if current.Right != nil {
				queue = append(queue, *current.Right)
			}
			if current.Left != nil {
				queue = append(queue, *current.Left)
			}
		}
	}
}

// Print the binary tree in spiral order
func (tree *GeneralTree) SpiralOrder() {
	if tree.Root == nil {
		return
	} else {
		tree.Root.SpiralOrder()
	}
}

func (node *Node) SpiralOrder() {
	if node == nil {
		return
	} else {
		stack1 := []Node{*node}
		stack2 := []Node{}
		for len(stack1) > 0 || len(stack2) > 0 {
			for len(stack1) > 0 {
				current := stack1[len(stack1)-1]
				stack1 = stack1[:len(stack1)-1]
				fmt.Println(current.Value)
				if current.Left != nil {
					stack2 = append(stack2, *current.Left)
				}
				if current.Right != nil {
					stack2 = append(stack2, *current.Right)
				}
			}
			for len(stack2) > 0 {
				current := stack2[len(stack2)-1]
				stack2 = stack2[:len(stack2)-1]
				fmt.Println(current.Value)
				if current.Right != nil {
					stack1 = append(stack1, *current.Right)
				}
				if current.Left != nil {
					stack1 = append(stack1, *current.Left)
				}
			}
		}
	}
}

// Print the binary tree in zig-zag order
func (tree *GeneralTree) ZigZagOrder() {
	if tree.Root == nil {
		return
	} else {
		tree.Root.ZigZagOrder()
	}
}

func (node *Node) ZigZagOrder() {
	if node == nil {
		return
	} else {
		stack1 := []Node{*node}
		stack2 := []Node{}
		for len(stack1) > 0 || len(stack2) > 0 {
			for len(stack1) > 0 {
				current := stack1[len(stack1)-1]
				stack1 = stack1[:len(stack1)-1]
				fmt.Println(current.Value)
				if current.Left != nil {
					stack2 = append(stack2, *current.Left)
				}
				if current.Right != nil {
					stack2 = append(stack2, *current.Right)
				}
			}
			for len(stack2) > 0 {
				current := stack2[len(stack2)-1]
				stack2 = stack2[:len(stack2)-1]
				fmt.Println(current.Value)
				if current.Right != nil {
					stack1 = append(stack1, *current.Right)
				}
				if current.Left != nil {
					stack1 = append(stack1, *current.Left)
				}
			}
		}
	}
}

// Print the binary tree in vertical order
func (tree *GeneralTree) VerticalOrder() {
	if tree.Root == nil {
		return
	} else {
		tree.Root.VerticalOrder()
	}
}

func (node *Node) VerticalOrder() {
	if node == nil {
		return
	} else {
		queue := []Node{*node}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			fmt.Println(current.Value)
			if current.Left != nil {
				queue = append(queue, *current.Left)
			}
			if current.Right != nil {
				queue = append(queue, *current.Right)
			}
		}
	}
}

// Print the binary tree in reverse-vertical order
func (tree *GeneralTree) ReverseVerticalOrder() {
	if tree.Root == nil {
		return
	} else {
		tree.Root.ReverseVerticalOrder()
	}
}

func (node *Node) ReverseVerticalOrder() {
	if node == nil {
		return
	} else {
		queue := []Node{*node}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			fmt.Println(current.Value)
			if current.Right != nil {
				queue = append(queue, *current.Right)
			}
			if current.Left != nil {
				queue = append(queue, *current.Left)
			}
		}
	}
}

// Print the binary tree in diagonal order
func (tree *GeneralTree) DiagonalOrder() {
	if tree.Root == nil {
		return
	} else {
		tree.Root.DiagonalOrder()
	}
}

func (node *Node) DiagonalOrder() {
	if node == nil {
		return
	} else {
		queue := []Node{*node}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			fmt.Println(current.Value)
			if current.Left != nil {
				queue = append(queue, *current.Left)
			}
			if current.Right != nil {
				queue = append(queue, *current.Right)
			}
		}
	}
}

// Print the binary tree in reverse-diagonal order
func (tree *GeneralTree) ReverseDiagonalOrder() {
	if tree.Root == nil {
		return
	} else {
		tree.Root.ReverseDiagonalOrder()
	}
}

func (node *Node) ReverseDiagonalOrder() {
	if node == nil {
		return
	} else {
		queue := []Node{*node}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			fmt.Println(current.Value)
			if current.Right != nil {
				queue = append(queue, *current.Right)
			}
			if current.Left != nil {
				queue = append(queue, *current.Left)
			}
		}
	}
}
