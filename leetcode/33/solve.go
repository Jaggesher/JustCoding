package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println("Case 2: ", search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println("Case 3: ", search([]int{1}, 0))
}

/***
 * Time: O(n log n)
 * Space: O(1)
 */
func search(nums []int, target int) int {
	st, end := 0, len(nums)-1
	for st < end {
		mid := (st + end) / 2
		if nums[st] <= nums[mid] {
			if nums[st] <= target && target <= nums[mid] {
				end = mid
			} else {
				st = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[end] {
				st = mid
			} else {
				end = mid - 1
			}
		}
	}
	if nums[st] == target {
		return st
	}
	return -1
}
