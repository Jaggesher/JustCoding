package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
	fmt.Println("Case 2: ", shiftGrid([][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}, 4))
	fmt.Println("Case 3: ", shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 9))
}

/***
 * Time: O(n * m * k)
 * Space: O(1) / O(m *n )[if we leave input data as it is]
 */

func shiftGrid(grid [][]int, k int) [][]int {
	for n := 0; n < k; n++ {
		var previous int
		for i := 0; i < len(grid); i++ {
			previous = grid[i][len(grid[i])-1]
			for j := 0; j < len(grid[i]); j++ {
				grid[i][j], previous = previous, grid[i][j]
			}
		}

		previous = grid[len(grid)-1][0]
		for i := 0; i < len(grid); i++ {
			grid[i][0], previous = previous, grid[i][0]
		}
	}
	return grid
}
