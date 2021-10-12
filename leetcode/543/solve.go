package main

import "fmt"

func main() {
	fmt.Println("No test case")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var ans int = 0
	findPath(root, &ans)
	return ans
}

func findPath(node *TreeNode, ans *int) int {
	if node == nil {
		return 0
	}
	left := findPath(node.Left, ans)
	right := findPath(node.Right, ans)

	if *ans < (left + right) {
		*ans = left + right
	}

	if left < right {
		return right + 1
	}
	return left + 1
}
