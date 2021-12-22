package main

import "fmt"

func main() {
	fmt.Println("No Test case")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/***
 * Time: O(n) => n is total number of node
 * Space: O(m) => size of queue
 */

func averageOfLevels(root *TreeNode) []float64 {
	var queue []*TreeNode
	var ans []float64
	queue = append(queue, root)
	for len(queue) > 0 {
		tmep := make([]*TreeNode, 0)
		sum := float64(0)
		for _, vl := range queue {
			sum += float64(vl.Val)
			if vl.Left != nil {
				tmep = append(tmep, vl.Left)
			}

			if vl.Right != nil {
				tmep = append(tmep, vl.Right)
			}
		}
		ans = append(ans, sum/float64(len(queue)))
		queue = tmep
	}
	return ans
}
