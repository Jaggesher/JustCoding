package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ac"))
}

func longestPalindrome(s string) string {
	st, end, ln := 0, 1, len(s)
	for i := 0; i < ln; i++ {
		tmSt := i
		for (i+1) < ln && s[i] == s[i+1] {
			i++
		}
		tmEnd := i
		for (tmSt-1) >= 0 && (tmEnd+1) < ln && s[tmSt-1] == s[tmEnd+1] {
			tmSt--
			tmEnd++
		}
		tmEnd++
		if (end - st) < (tmEnd - tmSt) {
			st = tmSt
			end = tmEnd
		}
	}
	return s[st:end]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
