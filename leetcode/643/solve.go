package main

import "fmt"

func main() {
	var nums []int = []int{1, 12, -5, -6, 50, 3}
	var k int = 4
	fmt.Println(findMaxAverage(nums, k))
}

func findMaxAverage(nums []int, k int) float64 {
	var ans float64 = 0
	for i := 0; i < k && i < len(nums); i++ {
		ans += (float64(nums[i]) / float64(k))
	}
	tm := ans
	for i := k; i < len(nums); i++ {
		tm -= float64(nums[i-k]) / float64(k)
		tm += float64(nums[i]) / float64(k)
		ans = max(ans, tm)
	}
	return ans
}

func max(a float64, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
