package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", combinationSum4([]int{1, 2, 3}, 4))
	fmt.Println("Case 2: ", combinationSum4([]int{9}, 3))
}

/***
 * Approach-1
 * Time: O(N*S) => where N is the length of nums and S is the target value
 * Space: O(S) => Where S is the length of target
 */

func combinationSum4(nums []int, target int) int {
	var memo []int = make([]int, target+1)
	memo[0] = 1
	for index := 1; index <= target; index++ {
		for _, vl := range nums {
			if vl <= index {
				memo[index] += memo[index-vl]
			}
		}
	}
	return memo[target]
}

/***
 * Approach-2
 * Time: O(N*S) => where N is the length of nums and S is the target value
 * Space: O(S) => Where S is the length of target
 */
func combinationSum4_2(nums []int, target int) int {
	var memo []int = make([]int, target+1)
	for i := range memo {
		memo[i] = -1
	}
	memo[0] = 1
	return trySum(target, nums, memo)
}

func trySum(left int, nums, memo []int) int {
	if memo[left] >= 0 {
		return memo[left]
	}
	temp := 0
	for _, vl := range nums {
		if left >= vl {
			temp += trySum(left-vl, nums, memo)
		}
	}
	memo[left] = temp
	return memo[left]
}
