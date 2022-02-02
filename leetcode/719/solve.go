package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1: ", smallestDistancePair([]int{1, 3, 1}, 1))
	fmt.Println("Case 2: ", smallestDistancePair([]int{1, 1, 1}, 2))
	fmt.Println("Case 3: ", smallestDistancePair([]int{1, 6, 1}, 3))
}

/***
 * Time : O(n log n) + n Log W => O(n log n)
 * Space: O(1)
 */

func smallestDistancePair(nums []int, k int) int { // n Log n + n Log(W)
	sort.Ints(nums) // n log n
	low, high := 0, nums[len(nums)-1]-nums[0]
	for low < high { // O(log W)
		mid := (low + high) / 2
		count := 0
		for left, right := 0, 0; right < len(nums); right++ { // O(n)
			for (nums[right] - nums[left]) > mid {
				left++
			}
			count += right - left
		}
		if count < k {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
