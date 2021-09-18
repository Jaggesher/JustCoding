package main

import "fmt"

func main() {
	fmt.Println("No test case.")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var node1, node2 *TreeNode

func recoverTree(root *TreeNode) {
	node1, node2 = nil, nil
	inorderTraverse(root, nil)
	node1.Val, node2.Val = node2.Val, node1.Val
}

func inorderTraverse(node *TreeNode, prev *TreeNode) *TreeNode {
	if node == nil {
		return prev
	}
	prev = inorderTraverse(node.Left, prev)
	if prev != nil && prev.Val > node.Val {
		node2 = node
		if node1 == nil {
			node1 = prev
		}
	}
	return inorderTraverse(node.Right, node)
}
