package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", numberOfArithmeticSlices([]int{1, 2, 3, 4}))
	fmt.Println("Case 2: ", numberOfArithmeticSlices([]int{1}))
}

/***
 * Time: O(n)
 * Space: O(1)
 */

func numberOfArithmeticSlices(nums []int) int {
	var ans int = 0
	dp := 0
	for i := 2; i < len(nums); i++ {
		if nums[i-1]-nums[i] == nums[i-2]-nums[i-1] {
			ans, dp = ans+dp+1, dp+1
		} else {
			dp = 0
		}
	}
	return ans
}
