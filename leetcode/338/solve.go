package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", countBits(2))
	fmt.Println("Case 2: ", countBits(5))
}

/***
 * Time: O(n)
 * Space: O(1)
 * Logic: Reset most significant set bit by (n &(n-1))
 */

func countBits(n int) []int {
	var ans []int = make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i&(i-1)] + 1
	}
	return ans
}
