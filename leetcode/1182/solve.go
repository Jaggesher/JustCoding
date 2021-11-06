package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Case 1:", shortestDistanceColor([]int{1, 1, 2, 1, 3, 2, 2, 3, 3}, [][]int{{1, 3}, {2, 2}, {6, 1}}))
	fmt.Println("Case 2:", shortestDistanceColor([]int{1, 2}, [][]int{{0, 3}}))
	fmt.Println("Case 3:", shortestDistanceColor([]int{3, 1, 1, 2, 3, 3, 2, 1, 2, 3, 1, 1, 3, 2, 3, 1, 1, 1, 1, 2, 2, 1, 2, 2, 2, 1, 1, 1, 1, 2, 3, 3, 3, 1, 3, 2, 1, 1, 2, 2, 1, 3, 1, 2, 1, 1, 2, 2, 1, 2},
		[][]int{{10, 2}, {0, 1}, {32, 3}, {1, 1}, {41, 1}, {48, 3}, {0, 3}, {46, 2}, {48, 2}, {28, 1}, {47, 2}, {11, 2}, {49, 3}, {3, 3}, {40, 1}, {1, 2}, {42, 2}, {47, 2}, {36, 3}, {23, 1}, {7, 3}, {47, 2}, {13, 3}, {36, 1}, {17, 2}, {46, 2}, {38, 2}, {0, 1}, {38, 3}, {36, 3}, {33, 1}, {11, 3}, {39, 2}, {10, 2}, {44, 3}, {5, 1}, {36, 3}, {44, 3}, {38, 1}, {9, 1}, {9, 1}, {35, 3}, {10, 1}, {9, 1}, {0, 3}, {1, 1}, {0, 3}, {28, 1}, {22, 3}, {15, 1}}))
}

func shortestDistanceColor(colors []int, queries [][]int) []int {
	var ans []int = make([]int, len(queries))
	var dictionary [][]int = [][]int{{}, {}, {}, {}}
	for index, val := range colors {
		dictionary[val] = append(dictionary[val], index)
	}
	for index, val := range queries {
		i, color := val[0], val[1]
		ans[index] = findBest(dictionary[color], i)
	}
	return ans
}

func findBest(arr []int, val int) int {
	if len(arr) == 0 {
		return -1
	}
	st, end := 0, len(arr)-1
	for st < end {
		mid := st + (end-st)/2
		if arr[mid] == val {
			return 0
		} else if arr[mid] < val {
			st = mid + 1
		} else {
			end = mid - 1
		}
	}
	if arr[st] > val && (st-1) >= 0 {
		end = st - 1
	} else if (st + 1) < len(arr) {
		end = st + 1
	}
	s1, s2 := math.Abs(float64(arr[st]-val)), math.Abs(float64(arr[end]-val))
	if s1 < s2 {
		return int(s1)
	}
	return int(s2)
}
