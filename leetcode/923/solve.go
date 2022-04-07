package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8))
	fmt.Println("Case 2: ", threeSumMulti([]int{1, 1, 2, 2, 2, 2}, 5))
}

/***
 * Time: O(n * n)
 * Space: O(n)
 */

func threeSumMulti(arr []int, target int) int {
	var ans int = 0
	var tracker [101]int
	const MOD = 1_000_000_007
	for _, vl := range arr {
		tracker[vl]++
	}

	for i := 0; i < len(arr); i++ {
		tracker[arr[i]]--
		temp := make([]int, 101)
		for j := i + 1; j < len(arr); j++ {
			temp[arr[j]]++
			tm := target - (arr[i] + arr[j])
			if tm >= 0 && tm <= 100 {
				ans += (tracker[tm] - temp[tm])
				ans %= MOD
			}
		}
	}

	return ans
}
