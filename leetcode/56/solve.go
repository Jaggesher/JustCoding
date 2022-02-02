package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1: ", merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println("Case 2: ", merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println("Case 3: ", merge([][]int{{1, 4}, {2, 3}}))
}

/***
 * Time: O(n) + O(n log n) => O(n log n)
 * Space: O(1) => if exclude the answer array
 */
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { // O(n log n)
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int = make([][]int, 0)
	temp := intervals[0]
	for i := 1; i < len(intervals); i++ { // O(n)
		if temp[1] < intervals[i][0] {
			ans = append(ans, temp)
			temp = intervals[i]
		} else if temp[1] < intervals[i][1] {
			temp[1] = intervals[i][1]
		}
	}
	ans = append(ans, temp)
	return ans
}
