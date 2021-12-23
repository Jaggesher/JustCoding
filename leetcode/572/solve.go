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
 * Time: O(n*m)
 * Space: O(n+m) => O(n)
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}
	return matchTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func matchTree(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return root1 == root2
	}
	return root1.Val == root2.Val && matchTree(root1.Left, root2.Left) && matchTree(root1.Right, root2.Right)
}
