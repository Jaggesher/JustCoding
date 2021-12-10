package main

import "fmt"

func main() {
	fmt.Println("Case 1:", uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println("Case 2:", uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var grid [101][101]int
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	grid[m-1][n-1] = (obstacleGrid[m-1][n-1] + 1) % 2
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if (i + 1) < m {
				grid[i][j] += grid[i+1][j]
			}
			if (j + 1) < n {
				grid[i][j] += grid[i][j+1]
			}
		}
	}
	return grid[0][0]
}
