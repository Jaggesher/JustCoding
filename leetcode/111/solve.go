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

/***
 * Time: O(n) => n is the number of node
 * Space: O(m) => m is the max queue size
 */
func minDepth(root *TreeNode) int {
	var ans int = 0
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		temp := make([]*TreeNode, 0)
		ans++
		for _, vl := range queue {
			if vl.Left == nil && vl.Right == nil {
				return ans
			}
			if vl.Left != nil {
				temp = append(temp, vl.Left)
			}
			if vl.Right != nil {
				temp = append(temp, vl.Right)
			}
		}
		queue = temp
	}
	return ans
}
