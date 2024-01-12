package main

import (
	"fmt"
	"math"
)

// TreeNode struct represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Creating a binary tree
	root := TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 11,
				},
			},
		},
	}

	// Serializing the binary tree
	serialized_tree := serialize(&root)
	// Deserializing the serialized binary tree
	deserialized_root := deserialize(serialized_tree)

	// Checking if the deserialized tree is correct
	if deserialized_root.Left.Right.Val == 4 &&
		deserialized_root.Right.Left.Left.Val == 11 {
		fmt.Println("Success")
	} else {
		fmt.Println("Failure")
	}
}

// deserialize function takes a serialized binary tree and returns the root of the binary tree
func deserialize(serialized_tree []int) *TreeNode {
	return deserializeTree(serialized_tree, 0)
}

// deserializeTree function is a helper function for the deserialize function
// It takes a serialized binary tree and an index, and returns the root of the binary tree
func deserializeTree(serialized_tree []int, i int) *TreeNode {
	// If the index is out of bounds or the value at the index is -1, return nil
	if i >= len(serialized_tree) || serialized_tree[i] == -1 {
		return nil
	}

	// Creating a new TreeNode with the value at the index
	// The left child is the node at the index 2*i+1
	// The right child is the node at the index 2*i+2
	root := TreeNode{
		Val:   serialized_tree[i],
		Left:  deserializeTree(serialized_tree, 2*i+1),
		Right: deserializeTree(serialized_tree, 2*i+2),
	}

	return &root
}

// serialize function takes the root of a binary tree and returns a serialized binary tree
func serialize(root *TreeNode) []int {
	// Getting the depth of the binary tree
	depth := getTreeDepth(root)
	// Creating an array with size 2^(depth+1)
	arr := make([]int, int(math.Pow(2, float64(depth+1))))
	// Serializing the binary tree into the array
	serializeTree(root, 0, &arr)
	return arr
}

// serializeTree function is a helper function for the serialize function
// It takes the root of a binary tree, an index, and an array, and serializes the binary tree into the array
func serializeTree(root *TreeNode, i int, arr *[]int) {
	if root != nil {
		// If the node is not nil, store its value in the array
		(*arr)[i] = root.Val
		// Serialize the left child at the index 2*i+1
		serializeTree(root.Left, 2*i+1, arr)
		// Serialize the right child at the index 2*i+2
		serializeTree(root.Right, 2*i+2, arr)
	} else {
		// If the node is nil, store -1 in the array
		(*arr)[i] = -1
	}
}

// getTreeDepth function takes the root of a binary tree and returns its depth
func getTreeDepth(root *TreeNode) int {
	if root == nil {
		// If the node is nil, its depth is 0
		return 0
	}

	// Getting the depth of the left child
	left_depth := getTreeDepth(root.Left)
	// Getting the depth of the right child
	right_depth := getTreeDepth(root.Right)

	// The depth of the node is the maximum depth of its children plus 1
	if left_depth > right_depth {
		return left_depth + 1
	} else {
		return right_depth + 1
	}
}
