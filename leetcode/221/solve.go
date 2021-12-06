package main

import (
	"fmt"
)

func main() {
	fmt.Println("Case 1:", maximalSquare([][]byte{{'1', '1', '1', '1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1', '1', '1', '0'}, {'1', '1', '1', '1', '1', '1', '1', '0'}, {'1', '1', '1', '1', '1', '0', '0', '0'}, {'0', '1', '1', '1', '1', '0', '0', '0'}}))
}

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	var totalSum [305][305]int
	var ans int = 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				totalSum[i][j] = min(min(totalSum[i-1][j-1], totalSum[i][j-1]), totalSum[i-1][j]) + 1
				if totalSum[i][j] > ans {
					ans = totalSum[i][j]
				}
				continue
			}
			totalSum[i][j] = 0
		}
	}
	return ans * ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
