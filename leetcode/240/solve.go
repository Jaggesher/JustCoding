package main

import "fmt"

func main() {
	fmt.Println("Case 1: ")
}

/***
 * Approach: 1
 * Time: O( m Log n)
 * Space: O(1)
 */
func searchMatrix(matrix [][]int, target int) bool {
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
