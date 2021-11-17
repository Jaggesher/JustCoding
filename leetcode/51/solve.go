package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Case 1:", solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	var ans [][]string = make([][]string, 0)
	backtrack(0, n, make([]int, 0, n), &ans)
	return ans
}

func backtrack(row, n int, queens []int, ans *[][]string) {
	if row == n {
		*ans = append(*ans, formResult(n, queens))
		return
	}

	for next := 0; next < n; next++ {
		flag := true
		for index, vl := range queens {
			if vl == next || math.Abs(float64(row-index)) == math.Abs(float64(next-vl)) {
				flag = false
				break
			}
		}
		if flag {
			backtrack(row+1, n, append(queens, next), ans)
		}
	}
}

func formResult(n int, queens []int) []string {
	temp := make([]string, 0, n)
	for _, vl := range queens {
		tm := make([]rune, n)
		for i := 0; i < n; i++ {
			if i == vl {
				tm[i] = 'Q'
			} else {
				tm[i] = '.'
			}
		}
		temp = append(temp, string(tm))
	}
	return temp
}
