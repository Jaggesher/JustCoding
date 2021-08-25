package main

import "fmt"

func main() {
	fmt.Println("No Test Case")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int = make([][]int, 0)
	var queue []*TreeNode = make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		tempQueue := make([]*TreeNode, 0)
		temp := make([]int, len(queue))
		for index, item := range queue {
			temp[index] = item.Val
			if item.Left != nil {
				tempQueue = append(tempQueue, item.Left)
			}
			if item.Right != nil {
				tempQueue = append(tempQueue, item.Right)
			}
		}
		ans = append(ans, temp)
		queue = tempQueue
	}
	return ans
}
