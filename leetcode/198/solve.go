package main

import "fmt"

func main() {
	fmt.Println("Case 1:", rob([]int{1, 2, 3, 1}))
	fmt.Println("Case 2:", rob([]int{2, 7, 9, 3, 1}))
}

func rob(nums []int) int {
	with, without := 0, 0
	for _, vl := range nums {
		temp := max(with, without)
		with = vl + without
		without = temp
	}
	return max(with, without)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
