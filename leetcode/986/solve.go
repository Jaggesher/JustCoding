package main

import (
	"fmt"
)

func main() {
	fmt.Println("Case 1: ", intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))
	fmt.Println("Case 2: ", intervalIntersection([][]int{{1, 3}, {5, 9}}, [][]int{}))
}

/***
 * Time: O(M+N)
 * Space: O(1)
 */
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var ans [][]int = make([][]int, 0)
	first, sec := 0, 0
	for first < len(firstList) && sec < len(secondList) {
		x, y := max(firstList[first][0], secondList[sec][0]), min(firstList[first][1], secondList[sec][1])
		if x <= y && firstList[first][0] <= x && y <= firstList[first][1] && secondList[sec][0] <= x && y <= secondList[sec][1] {
			ans = append(ans, []int{x, y})
		}
		if firstList[first][1] == secondList[sec][1] {
			first++
			sec++
		} else if firstList[first][1] < secondList[sec][1] {
			first++
		} else {
			sec++
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
