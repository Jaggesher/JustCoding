package main

import "fmt"

func main() {
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	fmt.Println(checkSubarraySum([]int{1, 2, 1, 6}, 6))
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 6))
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 13))
}

func checkSubarraySum(nums []int, k int) bool {
	var cuts map[int]int = make(map[int]int)
	var sum int = nums[0]
	cuts[sum%k] = 0
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		tm := sum % k
		if tm == 0 {
			return true
		}
		pos, ok := cuts[tm]
		if ok && pos < (i-1) {
			return true
		}
		if !ok {
			cuts[tm] = i
		}
	}
	return false
}
