package main

import "fmt"

func main() {
	fmt.Println("Case 1:", findPeakGrid([][]int{{1, 4}, {3, 2}}))
	fmt.Println("Case 2:", findPeakGrid([][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}))
	fmt.Println("Case 3:", findPeakGrid([][]int{{7, 44, 15, 31, 2, 40, 36}, {21, 40, 42, 5, 41, 30, 45}, {20, 42, 27, 8, 9, 3, 20}, {32, 8, 7, 16, 35, 9, 25}, {30, 24, 43, 48, 45, 35, 27}, {38, 48, 47, 10, 27, 42, 7}, {32, 40, 27, 18, 3, 45, 24}, {14, 29, 16, 24, 7, 44, 35}}))
}

func findPeakGrid(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i++ {
		mx := 0
		for l := 0; l < n; l++ {
			if mat[i][mx] < mat[i][l] {
				mx = l
			}
		}
		tm := 0
		for l := 0; l < m; l++ {
			if mat[tm][mx] < mat[l][mx] {
				tm = l
			}
		}
		if i == tm {
			return []int{i, mx}
		}
	}
	return []int{-1, -1}
}
