package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println("Case 2: ", findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}

/***
 * Time: O(n log n)
 * Space: O(1)
 */
func findPeakElement(nums []int) int {
	st, end := 0, len(nums)-1
	for st < end {
		mid := (st + end) / 2
		if nums[mid] < nums[mid+1] {
			st = mid + 1
		} else {
			end = mid
		}
	}
	return st
}
