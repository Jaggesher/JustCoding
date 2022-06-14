package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", minDistance("sea", "eat"))       // 2
	fmt.Println("Case 2: ", minDistance("leetcode", "etco")) // 4
}

func minDistance(word1 string, word2 string) int {
	return len(word1) + len(word2) - (2 * lcs(word1, word2))
}

/***
 * Return the Longest Common subsequence.
 * Time: O(nm) => Where n is the length of word1, and m is the length of word2
 * Space: O(nm)
 */
func lcs(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	var dp [][]int = make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[n1][n2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
