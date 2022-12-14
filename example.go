package main

import "fmt"

func main() {
	// Create a binary tree
	tree := GeneralTree{}
	tree.Root = &Node{Value: 1}
	tree.Root.Left = &Node{Value: 2}
	tree.Root.Right = &Node{Value: 3}
	tree.Root.Left.Left = &Node{Value: 4}
	tree.Root.Left.Right = &Node{Value: 5}
	tree.Root.Right.Left = &Node{Value: 6}
	tree.Root.Right.Right = &Node{Value: 7}
	tree.Root.Left.Left.Left = &Node{Value: 8}
	tree.Root.Left.Left.Right = &Node{Value: 9}
	tree.Root.Left.Right.Left = &Node{Value: 10}
	tree.Root.Left.Right.Right = &Node{Value: 11}
	tree.Root.Right.Left.Left = &Node{Value: 12}
	tree.Root.Right.Left.Right = &Node{Value: 13}
	tree.Root.Right.Right.Left = &Node{Value: 14}
	tree.Root.Right.Right.Right = &Node{Value: 15}
	tree.Root.Left.Left.Left.Left = &Node{Value: 16}
	tree.Root.Left.Left.Left.Right = &Node{Value: 17}
	tree.Root.Left.Left.Right.Left = &Node{Value: 18}

	// Print the binary tree in level order
	fmt.Println("Level Order")
	tree.LevelOrder()

	// Print the binary tree in reverse-level order
	fmt.Println("Reverse Level Order")
	tree.ReverseLevelOrder()

	// Print the binary tree in vertical order
	fmt.Println("Vertical Order")
	tree.VerticalOrder()

	// Print the binary tree in reverse-vertical order
	fmt.Println("Reverse Vertical Order")
	tree.ReverseVerticalOrder()

	// Print the binary tree in diagonal order
	fmt.Println("Diagonal Order")
	tree.DiagonalOrder()

	// Print the binary tree in reverse-diagonal order
	fmt.Println("Reverse Diagonal Order")
	tree.ReverseDiagonalOrder()

	// Create another tree with the creator function
	tree2 := NewGeneralTree()
	tree2.Root = &Node{Value: 1}
	tree2.Root.Left = &Node{Value: 2}
	tree2.Root.Right = &Node{Value: 3}
	tree2.Root.Left.Left = &Node{Value: 4}

	// Print the binary tree in level order
	fmt.Println("Level Order")
	tree2.LevelOrder()

	// Print the depth of the tree
	fmt.Println("Depth")
	fmt.Println(tree2.Depth())

	// Print the height of the tree
	fmt.Println("Height")
	fmt.Println(tree2.Height())

	// Print the number of nodes in the tree
	fmt.Println("Size")
	fmt.Println(tree2.Size())

	// Print the number of leaves in the tree
	fmt.Println("Leaves")
	fmt.Println(tree2.Leaves())

	// Print the number of full nodes in the tree
	fmt.Println("Full Nodes")
	fmt.Println(tree2.FullNodes())

	// Print the number of half nodes in the tree
	fmt.Println("Half Nodes")
	fmt.Println(tree2.HalfNodes())

	// Create another tree with the helper function
	tree3 := NewGeneralTreeWithNumbers(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18)

	// Print the binary tree in level order
	fmt.Println("Level Order")
	tree3.LevelOrder()
}
