package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1: ", findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	fmt.Println("Case 2: ", findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	fmt.Println("Case 3: ", findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
}

/***
 * Time: O(n Log n)+O(n) => O(n Log n)
 * Space: O(1) if we ignore the space of Sorting.
 */
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	var ans int = 1
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		if end < points[i][0] {
			ans++
			end = points[i][1]
		} else {
			end = min(end, points[i][1])
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
