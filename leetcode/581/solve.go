package main

import "fmt"

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(findUnsortedSubarray([]int{3, 2, 1}))
	fmt.Println(findUnsortedSubarray([]int{}))
}

func findUnsortedSubarray(nums []int) int {
	ln, st, end := len(nums)-1, -1, -1
	if ln < 0 {
		return 0
	}
	mn, mx := nums[ln], nums[0]

	for i := 0; i <= ln; i++ {
		if mx > nums[i] {
			end = i
		}
		if mx < nums[i] {
			mx = nums[i]
		}

		tm := ln - i
		if mn < nums[tm] {
			st = tm
		}
		if mn > nums[tm] {
			mn = nums[tm]
		}
	}

	if st == -1 {
		return 0
	}
	return end - st + 1
}
