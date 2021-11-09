package main

import "fmt"

func main() {
	fmt.Println("Case 1:", lengthOfLongestSubstringKDistinct("eceba", 2))
	fmt.Println("Case 2:", lengthOfLongestSubstringKDistinct("aa", 1))
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	var ans int = 0
	var tracker [256]int
	diff := 0
	for st, end := 0, 0; end < len(s); end++ {
		tracker[s[end]]++
		if tracker[s[end]] == 1 {
			diff++
		}
		for diff > k {
			tracker[s[st]]--
			if tracker[s[st]] == 0 {
				diff--
			}
			st++
		}
		if ans < (end - st + 1) {
			ans = end - st + 1
		}
	}
	return ans
}
