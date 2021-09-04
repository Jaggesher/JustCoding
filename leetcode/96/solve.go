package main

import "fmt"

func main() {
	fmt.Println("Case 1:", numTrees(3))
	fmt.Println("Case 2:", numTrees(1))
	fmt.Println("Case 3:", numTrees(19))
}

func numTrees(n int) int {
	var dp [20][20]int

	for i := 0; i < n; i++ {
		for j := 1; (j + i) <= n; j++ {
			st, end, val := j, j+i, 0
			for root := st; root <= end; root++ {
				left, right := 1, 1
				if st <= (root - 1) {
					left = dp[st][root-1]
				}
				if (root + 1) <= end {
					right = dp[root+1][end]
				}
				val += (left * right)
			}
			dp[st][end] = val
		}
	}
	return dp[1][n]
}
