package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println("Case 2: ", findTargetSumWays([]int{1}, 1))
}

/***
 * Time: O(t*n) => t is sum range
 * Memory: O(t) =>
 */
func findTargetSumWays(nums []int, target int) int {
	total := 0
	for _, vl := range nums {
		total += vl
	}
	if target > total || target < -total {
		return 0
	}

	var dp []int = make([]int, (total*2)+1)

	dp[total+nums[0]] = 1
	dp[total-nums[0]] += 1

	for i := 1; i < len(nums); i++ {
		temp := make([]int, (total*2)+1)
		for st := -total; st <= total; st++ {
			if dp[total+st] > 0 {
				temp[st+total+nums[i]] += dp[st+total]
				temp[st+total-nums[i]] += dp[st+total]
			}
		}
		dp = temp
	}

	return dp[total+target]
}
