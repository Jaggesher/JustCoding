package main

import "fmt"

func main() {
	fmt.Println("Case 1:", lengthOfLongestSubstringTwoDistinct("eceba"))
	fmt.Println("Case 2:", lengthOfLongestSubstringTwoDistinct("ccaabbb"))
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
	var ans int = 0
	var tracker [256]int
	diff := 0
	for st, end := 0, 0; end < len(s); end++ {
		tracker[s[end]]++
		if tracker[s[end]] == 1 {
			diff++
		}
		for diff > 2 {
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
