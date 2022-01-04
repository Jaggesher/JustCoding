package main

import "fmt"

func main() {
	fmt.Println("Case 1:", combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println("Case 2:", combinationSum([]int{2, 3, 5}, 8))
	fmt.Println("Case 3:", combinationSum([]int{2}, 1))
}

/***
 * Let N be the Number of Candidates & T be the target Value & M is the minimal
 * Time: O(N^((T/M)+1))
 * Space: O(T/M)
 */

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int = make([][]int, 0)
	var backtrack func(int, int, []int)
	backtrack = func(index int, sum int, soFar []int) {
		if sum == target {
			temp := make([]int, len(soFar))
			copy(temp, soFar)
			ans = append(ans, temp)
			return
		}
		if index >= len(candidates) {
			return
		}
		backtrack(index+1, sum, soFar)
		for tm := sum + candidates[index]; tm <= target; tm += candidates[index] {
			soFar = append(soFar, candidates[index])
			backtrack(index+1, tm, soFar)
		}
	}
	backtrack(0, 0, []int{})
	return ans
}
