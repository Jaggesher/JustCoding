package main

import (
	"fmt"
)

func main() {
	var stones1 []int = []int{5, 3, 1, 4, 2}
	var stones2 []int = []int{7, 90, 5, 1, 100, 10, 10, 2}
	var stones3 []int = []int{721, 979, 690, 84, 742, 873, 31, 323, 819, 22, 928, 866, 118, 843, 169, 818, 908, 832, 852, 480, 763, 715, 875, 629}

	fmt.Println("Case1: ", stoneGameVII(stones1))
	fmt.Println("Case1: ", stoneGameVII(stones2))
	fmt.Println("Case1: ", stoneGameVII(stones3))
}

func stoneGameVII(stones []int) int {
	var sumOfStones []int = make([]int, len(stones)+1)
	var cache [][]int = make([][]int, len(stones)+1)
	for i := range cache {
		tm := make([]int, len(stones)+1)
		for j := range tm {
			tm[j] = -10000000
		}
		cache[i] = tm
	}
	for i := 1; i <= len(stones); i++ {
		sumOfStones[i] = sumOfStones[i-1] + stones[i-1]
	}
	return PlayOptimally(sumOfStones, 1, len(stones), cache)
}

func PlayOptimally(stones []int, st int, end int, cache [][]int) int {
	if st == end {
		return 0
	}
	if cache[st][end] != -10000000 {
		return cache[st][end]
	}
	takeFirst, takeSecond := stones[end]-stones[st], stones[end-1]-stones[st-1]
	takeFirst -= PlayOptimally(stones, st+1, end, cache)
	takeSecond -= PlayOptimally(stones, st, end-1, cache)
	cache[st][end] = max(takeFirst, takeSecond)
	return cache[st][end]
}

func max(number1 int, number2 int) int {
	if number1 > number2 {
		return number1
	}
	return number2
}
