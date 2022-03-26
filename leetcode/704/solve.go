package main

import "fmt"

func main() {
	fmt.Println("Case 1:", search([]int{1, 0, 3, 5, 9, 12}, 9))
	fmt.Println("Case 2:", search([]int{1, 0, 3, 5, 9, 12}, 2))
}

/***
 * Approach: Binary Search
 * Time: O(n Log n)
 * Space: O(1)
 */
func search(nums []int, target int) int {
	st, end := 0, len(nums)-1

	for st <= end {
		mid := (st + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			st = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
