package main

import (
	"fmt"
)

func main() {
	fmt.Println("Case 1:", findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println("Case 2:", findDisappearedNumbers([]int{1, 1}))
}

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == -1 {
			continue
		}
		tm := nums[i]
		for nums[tm-1] != -1 {
			tm, nums[tm-1] = nums[tm-1], -1
		}
	}

	var ans []int = make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != -1 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
