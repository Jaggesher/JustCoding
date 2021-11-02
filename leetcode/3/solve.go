package main

import "fmt"

func main() {
	fmt.Println("Case 1:", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("Case 2:", lengthOfLongestSubstring("bbbbb"))
	fmt.Println("Case 3:", lengthOfLongestSubstring("pwwkew"))
	fmt.Println("Case 4:", lengthOfLongestSubstring("au"))
}

func lengthOfLongestSubstring(s string) int {
	var ans int = 0
	var flag map[byte]bool = make(map[byte]bool)
	for st, end := 0, 0; end < len(s); end++ {
		ch := s[end]
		if flag[ch] {
			for st < end && flag[ch] {
				delete(flag, s[st])
				st++
			}
		}
		flag[ch] = true
		if ans < len(flag) {
			ans = len(flag)
		}
	}
	return ans
}
