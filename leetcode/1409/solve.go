package main

import "fmt"

func main() {
	queries, m := []int{3, 1, 2, 1}, 5
	fmt.Println(processQueries(queries, m))
}

func processQueries(queries []int, m int) []int {
	var temp []int = make([]int, m+1)
	var ans []int = make([]int, len(queries))

	for i := 1; i <= m; i++ {
		temp[i] = i - 1
	}

	for i, val := range queries {
		fmt.Println(temp)
		index := temp[val]
		ans[i] = index

		for l := 1; l <= m; l++ {
			if temp[l] < index {
				temp[l]++
			} else if temp[l] == index {
				temp[l] = 0
			}
		}
	}

	return ans

}
