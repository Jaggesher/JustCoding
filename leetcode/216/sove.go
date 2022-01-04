package main

import "fmt"

func main() {
	fmt.Println("Case 1:", combinationSum3(3, 7))
	fmt.Println("Case 2:", combinationSum3(3, 9))
	fmt.Println("Case 3:", combinationSum3(4, 1))
}

/***
 * Time: less then O(9Ck) => in worse case O(9*9)
 * Space: O(k)
 */

func combinationSum3(k int, n int) [][]int {
	var ans [][]int = make([][]int, 0)

	var backtrack func(int, int, []int)
	backtrack = func(num, sum int, soFar []int) {
		if len(soFar) == k && sum == n {
			temp := make([]int, k)
			copy(temp, soFar)
			ans = append(ans, temp)
			return
		}

		if len(soFar) >= k {
			return
		}

		for i := num + 1; i < 10 && (sum+i) <= n; i++ {
			backtrack(i, sum+i, append(soFar, i))
		}
	}
	backtrack(0, 0, []int{})

	return ans
}
