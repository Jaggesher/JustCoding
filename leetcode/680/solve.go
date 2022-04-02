package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", validPalindrome("aba"))
	fmt.Println("Case 2: ", validPalindrome("abca"))
	fmt.Println("Case 3: ", validPalindrome("abc"))
}

/***
 * Approach: Greedy, Two-Pointer
 * Time: O(n)
 * Space: O(1)
 */
func validPalindrome(s string) bool {
	var st, end int = 0, len(s) - 1
	for st < end {
		if s[st] != s[end] {
			return checkPalindrome(s, st+1, end) || checkPalindrome(s, st, end-1)
		}
		st, end = st+1, end-1
	}

	return true
}

func checkPalindrome(s string, st, end int) bool {
	for st < end {
		if s[st] != s[end] {
			return false
		}
		st, end = st+1, end-1
	}
	return true
}
