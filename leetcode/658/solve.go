package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("Case 1:", findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	fmt.Println("Case 2:", findClosestElements([]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5))
}

/***
 * Approach: 1
 * Time: O(n log n) + O(k log k)
 * Space: O(k)
 */
func findClosestElements(arr []int, k int, x int) []int {
	sort.Slice(arr, func(a, b int) bool {
		tm1, tm2 := math.Abs(float64(arr[a]-x)), math.Abs(float64(arr[b]-x))
		if tm1 == tm2 {
			return arr[a] < arr[b]
		}
		return tm1 < tm2
	})
	ans := arr[:k]
	sort.Ints(ans)
	return ans
}
