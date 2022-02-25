package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println("Case 2: ", minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println("Case 3: ", minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}

/***
 * Time: O(n)
 * Space: O(1)
 */
func minSubArrayLen(target int, nums []int) int {
	var ans = 0
	st, tempSum := 0, nums[0]
	if tempSum >= target {
		ans = 1
	}

	for i := 1; i < len(nums); i++ {
		tempSum += nums[i]
		for st < i && (tempSum-nums[st]) >= target {
			tempSum -= nums[st]
			st++
		}
		if tempSum >= target && (ans == 0 || ans > (i-st+1)) {
			ans = i - st + 1
		}
	}

	return ans
}
