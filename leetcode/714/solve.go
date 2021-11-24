package main

import "fmt"

func main() {
	fmt.Println("Case 1:", maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
	fmt.Println("Case 2:", maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
}

/**
If I am holding a share after today, then either I am just continuing holding the share I had yesterday, or that I held no share yesterday, but bought in one share today: hold = max(hold, cash - prices[i])
If I am not holding a share after today, then either I did not hold a share yesterday, or that I held a share yesterday but I decided to sell it out today: cash = max(cash, hold + prices[i] - fee).
Make sure fee is only incurred once.
**/
func maxProfit(prices []int, fee int) int {
	var cash, hold int = 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		cash = mx(cash, hold+prices[i]-fee)
		hold = mx(hold, cash-prices[i])
	}
	return cash
}

func mx(a, b int) int {
	if a > b {
		return a
	}
	return b
}
