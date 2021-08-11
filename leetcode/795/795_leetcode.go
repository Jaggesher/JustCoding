package main

import "fmt"

func main() {
	var nums []int = []int{1, 2, 3, 2, 1}
	var left, right int = 3, 3

	fmt.Println(numSubarrayBoundedMax(nums, left, right))
}
func numSubarrayBoundedMax(nums []int, left int, right int) int {
	var ans int = 0
	var allInRange, lowValued int = 0, 0
	for _, val := range nums {
		if val < left {
			allInRange++
			lowValued++
		} else if val >= left && val <= right {
			allInRange++
			ans -= allPossible(lowValued)
			lowValued = 0
		} else {
			ans -= allPossible(lowValued)
			ans += allPossible(allInRange)
			lowValued, allInRange = 0, 0
		}
	}
	ans -= allPossible(lowValued)
	ans += allPossible(allInRange)
	return ans
}

func allPossible(n int) int {
	if n&1 == 1 {
		return n * int((n+1)/2)
	}
	return (n + 1) * int(n/2)
}
