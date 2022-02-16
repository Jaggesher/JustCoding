package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println("Case 2: ", insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	fmt.Println("Case 3: ", insert([][]int{{1, 2}}, []int{0, 0}))
}

/***
 * Time: O(n) =>
 * Space: O(1) => if we ignore the answer array
 */
func insert(intervals [][]int, newInterval []int) [][]int {
	var ans [][]int = make([][]int, 0)
	position := 0
	for position < len(intervals) && intervals[position][1] < newInterval[0] {
		position++
	}
	ans = make([][]int, position)
	copy(ans, intervals[:position])

	for position < len(intervals) && !(newInterval[1] < intervals[position][0] || intervals[position][1] < newInterval[0]) {
		if newInterval[0] > intervals[position][0] {
			newInterval[0] = intervals[position][0]
		}
		if newInterval[1] < intervals[position][1] {
			newInterval[1] = intervals[position][1]
		}
		position++
	}
	ans = append(ans, newInterval)

	if position < len(intervals) {
		ans = append(ans, intervals[position:]...)
	}
	return ans
}
