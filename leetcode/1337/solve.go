package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1:", kWeakestRows([][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}, 3))
	fmt.Println("Case 2:", kWeakestRows([][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}}, 2))
}

/***
 * Time: O(nm) + O(n log n)
 * Space: O(n)
 */
func kWeakestRows(mat [][]int, k int) []int {
	var ans []int = make([]int, len(mat))
	var sum []int = make([]int, len(mat))
	for index, arr := range mat {
		ans[index] = index
		for _, vl := range arr {
			sum[index] += vl
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		if sum[ans[i]] == sum[ans[j]] {
			return ans[i] < ans[j]
		}
		return sum[ans[i]] < sum[ans[j]]
	})

	return ans[:k]
}
