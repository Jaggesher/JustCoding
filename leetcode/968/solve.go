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

func minCameraCover(root *TreeNode) int {
	var ans int = 0
	tm := dfs(root, &ans)
	if tm == 1 {
		ans++
	}
	return ans
}

/***
 * Value 1: No camera
 * Value 2: Has coverage
 * Value 3: Camera installed here
 * Value 4: No need
 */
func dfs(node *TreeNode, ans *int) int {
	if node == nil {
		return 4
	}
	left, right := dfs(node.Left, ans), dfs(node.Right, ans)
	if left == 1 || right == 1 { // Yes we must need to install here
		*ans++
		return 3
	}

	if left == 4 && left == right { // This means edge node
		return 1
	}

	if left == 3 || right == 3 { // Yes I'm covered as my children have camera
		return 2
	}

	return 1 //Need to cover
}
