package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", maxProduct([]int{2, 3, -2, 4}))
	fmt.Println("Case 2: ", maxProduct([]int{-2, 0, -1}))
}

/***
 * Time: O(n) => n is the length of nums array
 * Space: O(1) => here we need constant space
 */
func maxProduct(nums []int) int {
	var ans int = nums[0]
	for i, mn, mx := 1, ans, ans; i < len(nums); i++ {
		tm_mn, tm_mx := mn*nums[i], mx*nums[i]
		mn = Min(nums[i], Min(tm_mn, tm_mx))
		mx = Max(nums[i], Max(tm_mn, tm_mx))
		ans = Max(ans, mx)
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
