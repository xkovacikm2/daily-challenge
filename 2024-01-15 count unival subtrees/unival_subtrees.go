package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	root := Node{Value: 0}
	root.Left = &Node{Value: 1}
	root.Right = &Node{Value: 0}
	root.Right.Left = &Node{Value: 1}
	root.Right.Left.Left = &Node{Value: 1}
	root.Right.Left.Right = &Node{Value: 1}
	root.Right.Right = &Node{Value: 0}

	univalSubtrees := countUnivalSubtrees(&root)

	if univalSubtrees != 5 {
		fmt.Println("Failure")
	}
	fmt.Println("Success")
}

func countUnivalSubtrees(node *Node) int {
	if node == nil {
		return 0
	} else if node.Left == nil && node.Right == nil {
		return 1
	}

	value := countUnivalSubtrees(node.Left) + countUnivalSubtrees(node.Right)
	if isUnival(node) {
		value++
	}
	return value
}

func isUnival(node *Node) bool {
	if node == nil {
		return true
	}
	return (node.Left == nil || node.Value == node.Left.Value) &&
		(node.Right == nil || node.Value == node.Right.Value) &&
		isUnival(node.Left) && 
		isUnival(node.Right) 
}