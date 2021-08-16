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

var flag []bool
var ans []int

func rightSideView(root *TreeNode) []int {
	flag = make([]bool, 101)
	ans = make([]int, 0)
	dfs(root, 0)
	return ans
}

func dfs(node *TreeNode, level int) {
	if node == nil {
		return
	}
	if !flag[level] {
		ans = append(ans, node.Val)
		flag[level] = true
	}
	dfs(node.Right, level+1)
	dfs(node.Left, level+1)
}
