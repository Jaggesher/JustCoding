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

func widthOfBinaryTree(root *TreeNode) int {
	var tracker [3002]int
	var ans int = 0
	bfs(root, 0, 1, &ans, &tracker)
	return ans
}

func bfs(node *TreeNode, level, num int, ans *int, tracker *[3002]int) {
	if node == nil {
		return
	}
	if tracker[level] == 0 {
		tracker[level] = num
	}
	tm := (num - tracker[level]) + 1
	if tm > *ans {
		*ans = tm
	}
	bfs(node.Left, level+1, (num*2)-1, ans, tracker)
	bfs(node.Right, level+1, (num * 2), ans, tracker)
}
