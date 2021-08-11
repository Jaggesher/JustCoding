package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
}

func findMin(nums []int) int {
	var ans int = 5001
	st, end := 0, len(nums)-1
	for st < end {
		mid := st + ((end - st) / 2)
		if nums[mid] <= nums[end] {
			end = mid - 1
			ans = min(ans, nums[mid])
		} else {
			ans = min(ans, nums[st])
			st = mid + 1
		}
	}
	ans = min(ans, nums[st])
	return ans
}

func min(a int, b int) (mn int) {
	if a < b {
		return a
	}
	return b
}
