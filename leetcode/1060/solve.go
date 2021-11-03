package main

import "fmt"

func main() {
	fmt.Println("Case 1:", missingElement([]int{4, 7, 9, 10}, 1))
	fmt.Println("Case 2:", missingElement([]int{4, 7, 9, 10}, 3))
	fmt.Println("Case 3:", missingElement([]int{1, 2, 4}, 3))
}

func missingElement(nums []int, k int) int {
	if k > missing(nums, 0, len(nums)-1) {
		return nums[len(nums)-1] + (k - missing(nums, 0, len(nums)-1))
	}

	st, end := 0, len(nums)-1
	for (st + 1) < end {
		mid := st + ((end - st) / 2)
		diff := missing(nums, st, mid)

		if diff >= k {
			end = mid
		} else {
			st, k = mid, k-diff
		}
	}
	return nums[st] + k
}

func missing(nums []int, st, end int) int {
	return (nums[end] - nums[st]) - (end - st)
}
