package main

import (
	"fmt"
)

func main() {
	fmt.Println("Case 1:", maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Println("Case 2:", maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println("Case 3:", maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println("Case 4:", maxProfit([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}))
	fmt.Println("Case 5:", maxProfit([]int{4, 7, 2, 1, 11}))
}

func maxProfit(prices []int) int {
	var ans, ln int = 0, len(prices)
	var mxFuture []int = make([]int, ln)

	temMx := 0
	for i := ln - 1; i >= 0; i-- {
		if prices[i] > temMx {
			temMx = prices[i]
		}
		mxFuture[i] = temMx
	}

	mn, lMax := prices[0], 0
	for i := 0; i < ln; i++ {
		if mn > prices[i] {
			mn = prices[i]
		}
		if lMax < (prices[i] - mn) {
			lMax = (prices[i] - mn)
		}

		if (i + 1) < ln {
			ans = max(ans, lMax+mxFuture[i+1]-prices[i+1])
		} else {
			ans = max(ans, lMax)
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
