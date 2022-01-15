package main

import (
	"fmt"
)

func main() {
	fmt.Println("Case 1: ", partition("aab"))
	fmt.Println("Case 2: ", partition("a"))
}

/***
 * Time: O(N * 2^N) Where N is the size of s
 * Memory: O(N*N)+O(N) => O(N*N) where N is the size of s
 */
func partition(s string) [][]string {
	var ans [][]string = make([][]string, 0)
	var memo [17][17]bool

	var recursion func(int, []string)
	recursion = func(index int, soFar []string) {
		if index >= len(s) {
			temp := make([]string, len(soFar))
			copy(temp, soFar)
			ans = append(ans, temp)
		}
		for cut := index; cut < len(s); cut++ {
			if s[index] == s[cut] && ((cut-index) <= 2 || memo[index+1][cut-1]) {
				memo[index][cut] = true
				recursion(cut+1, append(soFar, string(s[index:(cut+1)])))
			}
		}
	}
	recursion(0, make([]string, 0))
	return ans
}
