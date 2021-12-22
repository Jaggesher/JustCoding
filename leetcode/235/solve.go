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
 * Approach: Using Iteration
 * Time: O(n)
 * Space: O(1)
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else {
			return root
		}
	}
	return nil
}

/***
 * Approach: Using Recursion
 * Time: O(n)
 * Space: O(m)
 */

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left, right := lowestCommonAncestor2(root.Left, p, q), lowestCommonAncestor2(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
