package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 4}))
	fmt.Println(largestRectangleArea([]int{4, 4}))
	fmt.Println(largestRectangleArea([]int{2, 1, 2}))
}

func largestRectangleArea(heights []int) int {
	var ans int = heights[0]
	var tracker [][2]int = make([][2]int, 0)
	tracker = append(tracker, [2]int{heights[0], 0})
	for i := 1; i < len(heights); i++ {
		height := heights[i]
		ans = int(math.Max(float64(ans), float64(height)))
		tm, tmIndex := len(tracker), i
		for j := tm - 1; j >= 0; j-- {
			h, index := tracker[j][0], tracker[j][1]
			if h >= height {
				tm = j
				tmIndex = index // cut point
				ans = int(math.Max(float64(ans), float64((i-index+1)*height)))
			} else {
				ans = int(math.Max(float64(ans), float64((i-index+1)*h)))
			}
		}
		tracker = tracker[0:tm]
		tracker = append(tracker, [2]int{height, tmIndex})
	}
	return ans
}
