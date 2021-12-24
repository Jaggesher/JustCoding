package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println("Case 2: ", productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

/***
 * Time: O(n) + O(n) => O(n)
 * Space: O(n)
 */
func productExceptSelf(nums []int) []int {
	var ans []int = make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	tempProduct := 1
	for i := len(nums) - 2; i >= 0; i-- {
		tempProduct *= nums[i+1]
		ans[i] *= tempProduct
	}
	return ans
}
