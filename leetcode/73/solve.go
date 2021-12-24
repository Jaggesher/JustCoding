package main

import "fmt"

func main() {
	matrix1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix1)
	fmt.Println("Case 1:", matrix1)

	matrix2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix2)
	fmt.Println("Case 1:", matrix2)
}

/***
 * Approach: Try to mark which row & column need to be set 0, take special care for first row & column
 * Time: O(m*n)
 * Space: O(1)
 */
func setZeroes(matrix [][]int) {
	m, n, firstCol := len(matrix), len(matrix[0]), false
	for i := 0; i < m; i++ {
		firstCol = firstCol || (matrix[i][0] == 0)
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}

	if firstCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
