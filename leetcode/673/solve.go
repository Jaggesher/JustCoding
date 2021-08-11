package main

import "fmt"

func main() {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	fmt.Println(findNumberOfLIS([]int{1, 2, 2, 5, 5, 6, 6, 7}))
	fmt.Println(findNumberOfLIS([]int{10, 2}))
}

func findNumberOfLIS(nums []int) int {
	var ans int = 0
	var lengths []int = make([]int, len(nums))
	var weights []int = make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		mx, wg := 0, 1
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				if mx == lengths[j] {
					wg += weights[j]
				} else if mx < lengths[j] {
					mx = lengths[j]
					wg = weights[j]
				}
			}
		}
		lengths[i] = mx + 1
		weights[i] = wg
	}
	tmMx := 0
	for i := 0; i < len(nums); i++ {
		if lengths[i] == tmMx {
			ans += weights[i]
		} else if lengths[i] > tmMx {
			tmMx = lengths[i]
			ans = weights[i]
		}
	}

	return ans
}
