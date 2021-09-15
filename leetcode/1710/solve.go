package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1: ", maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
	fmt.Println("Case 2: ", maximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10))
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool { return boxTypes[i][1] > boxTypes[j][1] })
	var ans int = 0
	for _, vl := range boxTypes {
		if vl[0] < truckSize {
			ans += (vl[0] * vl[1])
			truckSize -= vl[0]
		} else {
			ans += (truckSize * vl[1])
			break
		}
	}
	return ans
}
