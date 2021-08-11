package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(9))
	fmt.Println(numSquares(10000))
}

func numSquares(n int) int {
	var visited []int = make([]int, n+1)
	var queue [][2]int = make([][2]int, 0)
	queue = append(queue, [2]int{n, 0})
	for len(queue) > 0 {
		num, level := queue[0][0], queue[0][1]
		queue = queue[1:]
		tm := int(math.Sqrt(float64(num)))
		for i := tm; i > 0; i-- {
			numTm := num - (i * i)
			if numTm == 0 {
				return level + 1
			}
			if visited[numTm] == 0 {
				visited[numTm] = level + 1
				queue = append(queue, [2]int{numTm, level + 1})
			}
		}

	}
	return n
}
