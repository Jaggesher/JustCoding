package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println("Case 2: ", search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
}

/***
 * Time: O(n) in worse cse & O(n log n) in best case
 * Space: O(1)
 */

func search(nums []int, target int) bool {
	st, end := 0, len(nums)-1
	for st < end {
		mid := (st + end) / 2
		for st < mid && nums[st] == nums[mid] {
			st++
		}

		for mid < end && nums[mid] == nums[end] {
			end--
		}

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
	return nums[st] == target
}
