package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println("Case 2: ", canJump([]int{3, 2, 1, 0, 4}))
}

/***
 * Time: O(n) => n is the length of nums
 * Space: O(1) => As we are using constant space
 */
func canJump(nums []int) bool {
	maxJump := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxJump <= 0 {
			return false
		}
		maxJump--
		if maxJump < nums[i] {
			maxJump = nums[i]
		}
	}
	return true
}
