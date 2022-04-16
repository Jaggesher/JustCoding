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
 * Time: O(n) => n is total nodes of the tree
 * Space: O(h) => h is the height of the tree
 */
func convertBST(root *TreeNode) *TreeNode {
	var totalSum int = 0
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Right)
		totalSum += node.Val
		node.Val = totalSum
		traverse(node.Left)
	}
	traverse(root)
	return root
}
