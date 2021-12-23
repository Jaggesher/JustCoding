package main

import "fmt"

func main() {
	fmt.Println("Case 1:", construct2DArray([]int{1, 2, 3, 4}, 2, 2))
	fmt.Println("Case 2:", construct2DArray([]int{1, 2, 3}, 1, 3))
	fmt.Println("Case 3:", construct2DArray([]int{1, 2}, 1, 1))
}

/***
 * Time: O(m*n)
 * Space: O(m*n)
 */

func construct2DArray(original []int, m int, n int) [][]int {
	if (m * n) != len(original) {
		return [][]int{}
	}
	var ans [][]int = make([][]int, 0)
	for i := 0; i < m; i++ {
		ans = append(ans, original[(i*n):n*(i+1)])
	}
	return ans
}
