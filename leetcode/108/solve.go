package main

import "fmt"

func main() {
	fmt.Println("Hello Jaggesher")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := (len(nums) / 2)
	return &TreeNode{Val: nums[mid], Left: sortedArrayToBST(nums[:mid]), Right: sortedArrayToBST(nums[(mid + 1):])}
}
