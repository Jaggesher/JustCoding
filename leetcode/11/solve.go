package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println("Case 2: ", maxArea([]int{1, 1}))
}

/***
 * Time: O(n)
 * Space: O(1)
 */
func maxArea(height []int) int {
	var ans, st, end int = 0, 0, len(height) - 1
	for st < end {
		ans = max(ans, min(height[st], height[end])*(end-st))
		if height[st] < height[end] {
			st++
		} else {
			end--
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
