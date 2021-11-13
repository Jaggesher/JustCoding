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

func getLonelyNodes(root *TreeNode) []int {
	var ans []int = make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		temp := make([]*TreeNode, 0)
		for _, vl := range queue {
			if vl.Left == nil && vl.Right == nil {
				continue
			} else if vl.Left != nil && vl.Right != nil {
				temp = append(temp, []*TreeNode{vl.Left, vl.Right}...)
			} else if vl.Left == nil {
				ans = append(ans, vl.Right.Val)
				temp = append(temp, vl.Right)
			} else {
				ans = append(ans, vl.Left.Val)
				temp = append(temp, vl.Left)
			}
		}
		queue = temp
	}
	return ans
}
