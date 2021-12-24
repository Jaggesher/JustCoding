package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1:", indexPairs("thestoryofleetcodeandme", []string{"story", "fleet", "leetcode"}))
	fmt.Println("Case 2:", indexPairs("ababa", []string{"aba", "ab"}))
}

/***
 * lets w lengths of words
 * lets n length of Text
 * Time: O(w log w) + O(w *t)
 * Space: O( w * n)
 */
func indexPairs(text string, words []string) [][]int {
	var ans [][]int
	// O( words log words)
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	// O( words * text)
	for i := range text {
		for _, word := range words {
			j := i + len(word)
			if j > len(text) {
				break
			}
			if word == text[i:i+len(word)] {
				ans = append(ans, []int{i, j - 1})
			}
		}
	}
	return ans
}
