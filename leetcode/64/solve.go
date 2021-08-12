package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Case 1:", minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println("Case 2:", minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			down, right := -1, -1
			if j+1 < n {
				right = grid[i][j+1]
			}
			if i+1 < m {
				down = grid[i+1][j]
			}
			if down != -1 && right != -1 {
				grid[i][j] += int(math.Min(float64(down), float64(right)))
			} else if right != -1 {
				grid[i][j] += right
			} else if down != -1 {
				grid[i][j] += down
			}
		}
	}
	return grid[0][0]
}
