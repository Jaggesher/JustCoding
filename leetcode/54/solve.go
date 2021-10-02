package main

import (
	"fmt"
)

func main() {
	fmt.Println("Case 1:", spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println("Case 2:", spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	fmt.Println("Case 3:", spiralOrder([][]int{{1, 2, 3, 4}}))
}

func spiralOrder(matrix [][]int) []int {
	var m, n, index int = len(matrix), len(matrix[0]), 0
	var ans []int = make([]int, (m * n))
	st, m, n := 0, m-1, n-1
	for st <= m && st <= n {
		//Go Right
		for i := st; i <= n; i++ {
			ans[index] = matrix[st][i]
			index++
		}
		//Go Down
		for i := st + 1; i <= m; i++ {
			ans[index] = matrix[i][n]
			index++
		}

		//Go Left
		for i := n - 1; i >= st && st != m; i-- {
			ans[index] = matrix[m][i]
			index++
		}

		//Go Top
		for i := m - 1; i > st && st != n; i-- {
			ans[index] = matrix[i][st]
			index++
		}
		st, m, n = st+1, m-1, n-1
	}
	return ans
}
