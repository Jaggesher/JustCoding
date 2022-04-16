package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", removeDuplicates([]int{1, 1, 2}))
	fmt.Println("Case 2: ", removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

/***
 * Time: O(n)
 * Space: O(1)
 */
func removeDuplicates(nums []int) int {
	pos, val := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != val {
			nums[pos] = nums[i]
			pos, val = pos+1, nums[i]
		}
	}
	return pos
}
