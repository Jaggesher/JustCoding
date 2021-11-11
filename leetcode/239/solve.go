package main

import "fmt"

func main() {
	fmt.Println("Case 1:", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println("Case 2:", maxSlidingWindow([]int{1}, 1))
	fmt.Println("Case 3:", maxSlidingWindow([]int{1, -1}, 1))
	fmt.Println("Case 4:", maxSlidingWindow([]int{9, 11}, 2))
	fmt.Println("Case 4:", maxSlidingWindow([]int{4, -2}, 2))
}

func maxSlidingWindow(nums []int, k int) []int {
	var ans []int = make([]int, len(nums)-k+1)
	leftMax, rightMax := make([]int, len(nums)), make([]int, len(nums))
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		leftMax[i], rightMax[j] = nums[i], nums[j]
		if (i % k) > 0 {
			leftMax[i] = max(leftMax[i], leftMax[i-1])
		}
		if ((j+1)%k) > 0 && (j+1) < len(nums) {
			rightMax[j] = max(rightMax[j], rightMax[j+1])
		}
	}

	for index, i, j := 0, 0, k-1; j < len(nums); index, i, j = index+1, i+1, j+1 {
		ans[index] = max(leftMax[j], rightMax[i])
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
