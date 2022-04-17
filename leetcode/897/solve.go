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

/***
 * Time: O(n)
 * Space: O(h) => Height of tree
 */
func increasingBST(root *TreeNode) *TreeNode {
	var dummy *TreeNode = &TreeNode{Val: 0, Left: nil, Right: nil}
	current := dummy

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		node.Left, current.Right = nil, node
		current = current.Right
		traverse(node.Right)
	}
	traverse(root)

	return dummy.Right
}
