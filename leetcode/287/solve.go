package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Case 1:", findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println("Case 2:", findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Println("Case 3:", findDuplicate([]int{1, 1}))
}

func findDuplicate1(nums []int) int {
	for _, vl := range nums {
		tm := int(math.Abs(float64(vl)))
		if nums[tm] < 0 {
			return tm
		}
		nums[tm] = -nums[tm]
	}
	return -1
}

func findDuplicate(nums []int) int {
	hare, turtle := nums[0], nums[0]
	for flag := true; flag; flag = hare != turtle {
		hare = nums[nums[hare]]
		turtle = nums[turtle]
	}
	hare = nums[0]
	for hare != turtle {
		hare = nums[hare]
		turtle = nums[turtle]
	}
	return hare
}
