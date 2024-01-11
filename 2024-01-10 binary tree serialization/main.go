package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
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

	serialized_tree := serialize(&root)
	deserialized_root := deserialize(serialized_tree)

	if deserialized_root.Left.Right.Val == 4 &&
		deserialized_root.Right.Left.Left.Val == 11{
		fmt.Println("Success")
	} else {
		fmt.Println("Failure")
	}
}

func deserialize(serialized_tree []int) *TreeNode {
	panic("unimplemented")
}

func serialize(treeNode *TreeNode) []int {
	panic("unimplemented")
}
