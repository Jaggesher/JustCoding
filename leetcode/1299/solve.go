package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("Case 1:", minAvailableDuration([][]int{{10, 50}, {60, 120}, {140, 210}}, [][]int{{0, 15}, {60, 70}}, 8))
	fmt.Println("Case 2:", minAvailableDuration([][]int{{10, 50}, {60, 120}, {140, 210}}, [][]int{{0, 15}, {60, 70}}, 12))
}

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	sort.Slice(slots1, func(i, j int) bool { return slots1[i][0] < slots1[j][0] })
	sort.Slice(slots2, func(i, j int) bool { return slots2[i][0] < slots2[j][0] })
	i1, i2 := 0, 0
	for i1 < len(slots1) && i2 < len(slots2) {
		st, end := int(math.Max(float64(slots1[i1][0]), float64(slots2[i2][0]))), int(math.Min(float64(slots1[i1][1]), float64(slots2[i2][1])))
		if (st + duration) <= end {
			return []int{st, st + duration}
		}
		if slots1[i1][1] < slots2[i2][1] {
			i1++
		} else {
			i2++
		}
	}
	return []int{}
}
