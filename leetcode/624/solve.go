package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Case 1:", maxDistance([][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}))
	fmt.Println("Case 2:", maxDistance([][]int{{1}, {1}}))
}

func maxDistance(arrays [][]int) int {
	var ans = 0
	mn, mx := arrays[0][0], arrays[0][len(arrays[0])-1]
	for i := 1; i < len(arrays); i++ {
		n := len(arrays[i]) - 1

		ans = Max(ans, Max(int(math.Abs(float64(mn-arrays[i][n]))), int(math.Abs(float64(mx-arrays[i][0])))))

		mn = Min(mn, arrays[i][0])
		mx = Max(mx, arrays[i][n])
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
