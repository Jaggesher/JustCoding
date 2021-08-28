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

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int = make([][]int, 0)
	var queue []*TreeNode = make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	flag := false
	for len(queue) > 0 {
		tempQueue := make([]*TreeNode, 0)
		arr := make([]int, len(queue))
		for index, vl := range queue {
			arr[index] = vl.Val
			if vl.Left != nil {
				tempQueue = append(tempQueue, vl.Left)
			}
			if vl.Right != nil {
				tempQueue = append(tempQueue, vl.Right)
			}
		}
		if flag {
			for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}

		ans = append(ans, arr)
		queue = tempQueue
		flag = !flag
	}

	return ans
}
