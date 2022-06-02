package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println("Case 1: ", transpose([][]int{{1, 2, 3}, {4, 5, 6}}))
}

/***
 * Time: O(n * m)
 * Space: O(n * m)
 */
func transpose(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	var transpose [][]int = make([][]int, m)
	for i := 0; i < m; i++ {
		transpose[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			transpose[j][i] = matrix[i][j]
		}
	}
	return transpose
}
