package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1:", combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println("Case 2:", combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}

/***
 * Time: O(2^N)
 * Space: O(N)
 */

func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int = make([][]int, 0)
	sort.Ints(candidates)
	var backtrack func(int, int, []int)
	backtrack = func(index, sum int, soFar []int) {
		if sum == target {
			temp := make([]int, len(soFar))
			copy(temp, soFar)
			ans = append(ans, temp)
			return
		}
		if index >= len(candidates) {
			return
		}
		nextIndex := index + 1
		for nextIndex < len(candidates) && candidates[index] == candidates[nextIndex] {
			nextIndex++
		}
		backtrack(nextIndex, sum, soFar)
		for index < nextIndex && (sum+candidates[index]) <= target {
			sum += candidates[index]
			soFar = append(soFar, candidates[index])
			index++
			backtrack(nextIndex, sum, soFar)
		}
	}
	backtrack(0, 0, []int{})
	return ans
}
