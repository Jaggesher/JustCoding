package main

import "fmt"

func main() {
	fmt.Println("Case 1:", getFactors(1))
	fmt.Println("Case 2:", getFactors(12))
	fmt.Println("Case 3:", getFactors(37))
	fmt.Println("Case 4:", getFactors(32))
	fmt.Println("Case 5:", getFactors(8192))
}

func getFactors(n int) [][]int {
	return backTrack(2, n)
}

func backTrack(st, n int) [][]int {
	var result [][]int = make([][]int, 0)
	for i := st; (i * i) <= n; i++ {
		if n%i == 0 {
			result = append(result, []int{i, n / i})
			resp := backTrack(i, n/i)
			for _, vl := range resp {
				result = append(result, append(vl, i))
			}
		}
	}
	return result
}
