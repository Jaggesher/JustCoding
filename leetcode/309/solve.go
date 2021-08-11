package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1}))
}

func maxProfit(prices []int) int {
	var cache []int = make([]int, len(prices))
	for i := range cache {
		cache[i] = -1
	}
	return dp(0, prices, cache)
}

func dp(index int, prices []int, cache []int) int {
	if index >= len(prices) {
		return 0
	}
	if cache[index] != -1 {
		return cache[index]
	}
	mx := 0
	for i := index + 1; i < len(prices); i++ {
		if prices[index] < prices[i] {
			tm := (prices[i] - prices[index]) + dp(i+2, prices, cache)
			mx = max(mx, tm)
		}
	}
	cache[index] = max(mx, dp(index+1, prices, cache))
	return cache[index]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
