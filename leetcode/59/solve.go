package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", generateMatrix(3))
	fmt.Println("Case 1: ", generateMatrix(1))
}

func generateMatrix(n int) [][]int {
	var ans [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	for i, j, count := 0, n-1, 1; i <= j; i, j = i+1, j-1 {
		// go Right
		for index := i; index <= j; index++ {
			ans[i][index] = count
			count++
		}

		// Go Down
		for index := i + 1; index <= j; index++ {
			ans[index][j] = count
			count++
		}

		//Go Left
		for index := j - 1; index >= i; index-- {
			ans[j][index] = count
			count++
		}

		//Go Top
		for index := j - 1; index > i; index-- {
			ans[index][i] = count
			count++
		}
	}

	return ans
}
