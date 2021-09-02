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

var tracker *TreeNode = nil

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	left, right := root.Left, root.Right
	root.Left, root.Right = nil, nil

	if tracker != nil {
		tracker.Right = root
	}

	tracker = root
	flatten(left)
	flatten(right)
}
