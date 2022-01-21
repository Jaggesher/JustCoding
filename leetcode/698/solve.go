package main

import (
	"fmt"
)

func main() {
	fmt.Println("Case 1: ", canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	fmt.Println("Case 2: ", canPartitionKSubsets([]int{1, 2, 3, 4}, 3))
	fmt.Println("Case 3: ", canPartitionKSubsets([]int{3, 1, 1, 1, 3, 3}, 3))
	fmt.Println("Case 3: ", canPartitionKSubsets([]int{129, 17, 74, 57, 1421, 99, 92, 285, 1276, 218, 1588, 215, 369, 117, 153, 22}, 3))
}

/***
 * Time: O( N * 2^N)
 * Space: O(N)+O(2^N) = O(2^N)
 */
func canPartitionKSubsets(nums []int, k int) bool {
	var sum int
	for _, vl := range nums {
		sum += vl
	}
	if (sum % k) > 0 {
		return false
	}
	var taken []byte = make([]byte, len(nums))
	var memo map[string]bool = make(map[string]bool)
	return backtrack(0, sum, 0, sum/k, taken, nums, memo)
}

func backtrack(index, left, soFar, cut int, taken []byte, nums []int, memo map[string]bool) bool {
	if vl, ok := memo[string(taken)]; ok {
		return vl
	}

	if left == 0 {
		return true
	}
	if soFar == cut {
		return backtrack(0, left, 0, cut, taken, nums, memo)
	}
	result := false
	for i := index; i < len(taken) && !result; i++ {
		if taken[i] != '1' && (soFar+nums[i]) <= cut {
			taken[i] = '1'
			result = backtrack(i+1, left-nums[i], soFar+nums[i], cut, taken, nums, memo)
			taken[i] = 0
		}
	}
	memo[string(taken)] = result
	return result
}
