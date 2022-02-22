package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println("Case 2: ", searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
}

/***
 * Time: O(mn)
 * Space: O(1)
 */
func searchMatrix(matrix [][]int, target int) bool {
	st, end := 0, len(matrix)-1
	for st < end {
		mid := (st + end) / 2
		if matrix[mid][len(matrix[mid])-1] < target {
			st = mid + 1
		} else {
			end = mid
		}
	}
	m := st
	st, end = 0, len(matrix[m])-1
	for st < end {
		mid := (st + end) / 2
		if matrix[m][mid] < target {
			st = mid + 1
		} else {
			end = mid
		}
	}
	return matrix[m][st] == target
}
