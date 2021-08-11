package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))
	fmt.Println(countSmaller([]int{-1}))
	fmt.Println(countSmaller([]int{-1, -1}))
}

func countSmaller(nums []int) []int {
	var ans []int = make([]int, len(nums))
	mn, mx := findMinMax(nums)
	var BIT []int = make([]int, mx-mn+1)
	for i := len(nums) - 1; i >= 0; i-- {
		index := (nums[i] - mn + 1)
		ans[i] = SearchBIT(BIT, index-1)
		UpdateBIT(BIT, index, 1)
	}
	return ans
}

func findMinMax(nums []int) (int, int) {
	mn, mx := nums[0], nums[0]
	for _, vl := range nums {
		if mn > vl {
			mn = vl
		} else if mx < vl {
			mx = vl
		}
	}
	return mn, mx
}
func UpdateBIT(BIT []int, index int, val int) {
	for index < len(BIT) {
		BIT[index] += val
		index += (index & (-index))
	}
}

func SearchBIT(BIT []int, index int) int {
	result := 0
	for index > 0 {
		result += BIT[index]
		index -= (index & -index)
	}
	return result
}
