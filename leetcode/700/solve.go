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

func searchBST(root *TreeNode, val int) *TreeNode {
	var ans *TreeNode = root
	for ans != nil && ans.Val != val {
		if ans.Val < val {
			ans = ans.Right
		} else {
			ans = ans.Left
		}
	}
	return ans
}
