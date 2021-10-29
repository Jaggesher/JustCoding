package main

import "fmt"

func main() {
	fmt.Println("Case 1:", sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println("Case 2:", sortedSquares([]int{-7, -3, 2, 3, 11}))
}
func sortedSquares(nums []int) []int {
	ln := len(nums)
	var ans []int = make([]int, ln)
	left, right, index := 0, ln-1, ln-1
	var abs = func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	for left <= right {
		if abs(nums[left]) < abs(nums[right]) {
			ans[index] = nums[right] * nums[right]
			right--
		} else {
			ans[index] = nums[left] * nums[left]
			left++
		}
		index--
	}
	return ans
}
