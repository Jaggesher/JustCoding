package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1:", eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println("Case 2:", eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
	fmt.Println("Case 3:", eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}))
}

/***
 * Approach: [Greedy]
 * Time: O(n log n)
 * Space: O(1)
 */

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { // n log n
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	var ans, ptr int = 0, 0
	for i := 1; i < len(intervals); i++ {
		if intervals[ptr][1] <= intervals[i][0] {
			ptr = i
		} else if intervals[ptr][0] <= intervals[i][0] && intervals[ptr][1] >= intervals[i][1] {
			ptr = i
			ans++
		} else {
			ans++
		}
	}
	return ans
}
