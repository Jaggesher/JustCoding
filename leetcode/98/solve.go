package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("No testcase available")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return BST(root, math.MinInt64, math.MaxInt64)
}

func BST(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	return BST(node.Left, min, node.Val) && BST(node.Right, node.Val, max)
}
