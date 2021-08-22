package main

import (
	"fmt"
)

func main() {
	fmt.Println("No Testcase Problem")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var ans []int = make([]int, 0)
	findAndBack(root, target.Val, k, &ans)
	return ans
}

func findAndBack(node *TreeNode, target, k int, ans *[]int) int {
	if node == nil {
		return -1
	}
	if node.Val == target {
		findAns(node, 0, k, ans)
		return 1
	}
	left := findAndBack(node.Left, target, k, ans)
	if left > -1 {
		if left == k {
			*ans = append(*ans, node.Val)
		} else {
			findAns(node.Right, left+1, k, ans)
		}
		return left + 1
	}

	right := findAndBack(node.Right, target, k, ans)
	if right > -1 {
		if right == k {
			*ans = append(*ans, node.Val)
		} else {
			findAns(node.Left, right+1, k, ans)
		}
		return right + 1
	}
	return -1
}

func findAns(node *TreeNode, distance, k int, ans *[]int) {
	if node == nil || distance > k {
		return
	}
	if distance == k {
		*ans = append(*ans, node.Val)
	}
	findAns(node.Left, distance+1, k, ans)
	findAns(node.Right, distance+1, k, ans)
}
