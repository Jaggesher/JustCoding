package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", isSubsequence("abc", "ahbgdc"))
	fmt.Println("Case 2: ", isSubsequence("axc", "ahbgdc"))
}

/***
 * Time: O(s+t)
 * Space: O(1)
 */
func isSubsequence(s string, t string) bool {
	sIndex, tIndex := 0, 0
	for sIndex < len(s) && tIndex < len(t) {
		if s[sIndex] == t[tIndex] {
			sIndex++
		}
		tIndex++
	}
	return sIndex == len(s)
}
