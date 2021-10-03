package main

import "fmt"

func main() {
	fmt.Println("Case 1:", uniquePaths(3, 7))
	fmt.Println("Case 2:", uniquePaths(3, 2))
}

func uniquePaths(m int, n int) int {
	var data []int = make([]int, n)
	for i := 0; i < m; i++ {
		data[0] = 1
		for j := 1; j < n; j++ {
			data[j] += data[j-1]
		}
	}
	return data[n-1]
}
