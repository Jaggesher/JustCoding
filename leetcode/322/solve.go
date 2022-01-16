package main

import (
	"fmt"
)

func main() {
	fmt.Println("Case 1: ", coinChange2([]int{1, 2, 5}, 11))
	fmt.Println("Case 2: ", coinChange2([]int{2}, 3))
}

/***
 * Time: O(S*n) => S is the Amount, & n is the number of coins
 * Space: O(S) => S is the Amount
 */

func coinChange(coins []int, amount int) int {
	var flag map[int]int = make(map[int]int)
	var recursion func(int) int
	recursion = func(left int) int {
		if left <= 0 {
			return left
		}
		if vl, ok := flag[left]; ok {
			return vl
		}
		min := -1
		for _, coin := range coins {
			tm := recursion(left-coin) + 1
			if tm > 0 && (min == -1 || min > tm) {
				min = tm
			}
		}
		flag[left] = min
		return flag[left]
	}
	recursion(amount)
	return flag[amount]
}

func coinChange2(coins []int, amount int) int {
	var dp []int = make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if (i - coin) >= 0 {
				dp[i] = Min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == (amount + 1) {
		return -1
	}
	return dp[amount]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
