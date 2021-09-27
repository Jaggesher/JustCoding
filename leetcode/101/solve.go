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

func isSymmetric(root *TreeNode) bool {
	var queue []*TreeNode = make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		left, right := -1, len(queue)
		for left < right {
			left++
			right--
			if queue[left] == nil && queue[right] == nil {
				continue
			} else if queue[left] == nil || queue[right] == nil || queue[left].Val != queue[right].Val {
				return false
			}
		}
		temp := make([]*TreeNode, 0)
		for _, tm := range queue {
			if tm == nil {
				continue
			}
			temp = append(temp, tm.Left)
			temp = append(temp, tm.Right)
		}
		queue = temp
	}
	return true
}
