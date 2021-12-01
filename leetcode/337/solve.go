package main

import "fmt"

func main() {
	fmt.Println("No test Case")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return max(robBest(root))
}

func robBest(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftWith, leftWithout := robBest(root.Left)
	rightWith, rightWithout := robBest(root.Right)
	return root.Val + leftWithout + rightWithout, max(leftWith, leftWithout) + max(rightWith, rightWithout)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
