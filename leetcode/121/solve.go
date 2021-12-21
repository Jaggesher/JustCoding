package main

import "fmt"

func main() {
	fmt.Println("Case 1:", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println("Case 2:", maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	var ans, mn int = 0, prices[0]
	for _, vl := range prices {
		ans = Max(ans, (vl - mn))
		mn = Min(mn, vl)
	}

	return ans
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
