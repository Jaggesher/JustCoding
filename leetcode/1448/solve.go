package main

import (
	"fmt"
)

func main() {
	case1 := TreeNode{Val: 1, Left: nil, Right: nil}
	fmt.Println("Case 1:", goodNodes(&case1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(root *TreeNode, mx int) int {
	if root == nil {
		return 0
	}
	count := 0
	if root.Val >= mx {
		count++
		mx = root.Val
	}
	return count + dfs(root.Right, mx) + dfs(root.Left, mx)
}
