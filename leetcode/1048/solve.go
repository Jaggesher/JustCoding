package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1: ", longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"}))
	fmt.Println("Case 2: ", longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}))
	fmt.Println("Case 3: ", longestStrChain([]string{"abcd", "dbqca"}))
}

/***
 * Time: O(n) + O(n log n) => O(n log n)
 * Space: O(n)
 */
func longestStrChain(words []string) int {
	var dp map[string]int = make(map[string]int)
	var ans int

	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) }) // O( n log n)
	for _, word := range words {                                                    // O(n)
		var mx = 0
		for index := range word {
			tm := string(word[:index]) + string(word[index+1:]) // O(m) since it's only 16 so => O(1)
			if mx < dp[tm] {
				mx = dp[tm]
			}
		}
		dp[word] = mx + 1
		if ans < dp[word] {
			ans = dp[word]
		}
	}
	return ans
}
