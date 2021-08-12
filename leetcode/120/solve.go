package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Case 1: ", minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
	fmt.Println("Case 2: ", minimumTotal([][]int{{-10}}))
}

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += int(math.Min(float64(triangle[i+1][j]), float64(triangle[i+1][j+1])))
		}
	}
	return triangle[0][0]
}
