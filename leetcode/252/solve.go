package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1: ", canAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println("Case 2: ", canAttendMeetings([][]int{{7, 10}, {2, 4}}))
}

/***
 * Time: O(n Log n) + O(n) => O(n Log n)
 * Space: O(1)
 */

func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] > intervals[i][0] {
			return false
		}
	}
	return true
}
