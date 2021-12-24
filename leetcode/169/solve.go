package main

import "fmt"

func main() {
	fmt.Println("Case 1:", majorityElement([]int{3, 2, 3}))
	fmt.Println("Case 2:", majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

/***
 * Algorithm: Boyer-Moore Voting Algorithm
 * Time: O(n)
 * Space: O(1)
 */
func majorityElement(nums []int) int {
	var majority, count int = nums[0], 1
	for i := 1; i < len(nums); i++ {
		if majority == nums[i] {
			count++
		} else {
			count--
			if count == 0 {
				majority, count = nums[i], 1
			}
		}
	}
	return majority
}
