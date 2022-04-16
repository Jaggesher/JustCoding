package main

import "fmt"

func main() {
	fmt.Println("Case 1:", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println("Case 2:", maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println("Case 3:", maxProfit([]int{7, 6, 4, 3, 1}))
}

/***
 * Time: O(n)
 * Space: O(1)
 */
func maxProfit(prices []int) int {
	var ans int = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}
