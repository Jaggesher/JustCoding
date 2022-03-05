package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Case 1: ", champagneTower(1, 1, 1))
	fmt.Println("Case 2: ", champagneTower(2, 1, 1))
	fmt.Println("Case 3: ", champagneTower(100000009, 33, 17))
}

func champagneTower(poured int, query_row int, query_glass int) float64 {
	var stack [101][101]float64
	stack[0][0] = float64(poured)
	for i := 0; i <= query_row; i++ {
		for j := 0; j <= i; j++ {
			temp := float64(stack[i][j]-1) / 2
			if temp > 0 {
				stack[i+1][j] += temp
				stack[i+1][j+1] += temp
			}
		}
	}
	return math.Min(stack[query_row][query_glass], 1)
}
