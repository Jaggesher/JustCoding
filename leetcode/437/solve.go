package main

import (
	"fmt"
)

func main() {
	fmt.Println("No test case")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var mp map[int]int = make(map[int]int)
	mp[0] = 1
	return dfs(root, 0, targetSum, mp)
}

func dfs(node *TreeNode, soFar, targetSum int, mp map[int]int) int {
	if node == nil {
		return 0
	}
	soFar += node.Val
	cnt := mp[(soFar - targetSum)]
	mp[soFar]++
	cnt += dfs(node.Left, soFar, targetSum, mp)
	cnt += dfs(node.Right, soFar, targetSum, mp)
	mp[soFar]--
	return cnt
}
