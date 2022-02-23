package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	fmt.Println("Case 2: ", searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
}

/***
 * Approach: 1
 * Time: O( m Log n)
 * Space: O(1)
 */
func searchMatrix1(matrix [][]int, target int) bool {
	for _, arr := range matrix {
		st, end := 0, len(arr)-1
		for st <= end {
			mid := (st + end) / 2
			if target == arr[mid] {
				return true
			} else if target < arr[mid] {
				end = mid - 1
			} else {
				st = mid + 1
			}
		}
	}
	return false
}

/***
 * Approach: 2
 * Time: (m+n)
 * Space: O(1)
 */

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	u, v := m-1, 0
	for u >= 0 && v < n {
		if matrix[u][v] == target {
			return true
		} else if matrix[u][v] < target {
			v++
		} else {
			u--
		}
	}
	return false
}
