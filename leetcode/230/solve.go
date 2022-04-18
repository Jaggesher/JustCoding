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
 * Space: O(h)
 */
func kthSmallest(root *TreeNode, k int) int {
	var ans int
	bfs(root, k, 0, &ans)
	return ans
}

func bfs(node *TreeNode, k, soFar int, ans *int) int {
	if node == nil || k <= soFar {
		return soFar
	}
	soFar = bfs(node.Left, k, soFar, ans)
	if (soFar + 1) == k {
		*ans = node.Val
	}
	return bfs(node.Right, k, soFar+1, ans)
}
