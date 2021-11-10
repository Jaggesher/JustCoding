package main

import (
	"fmt"
)

func main() {
	fmt.Println("Case 1:", minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println("Case 2:", minWindow("a", "a"))
	fmt.Println("Case 3:", minWindow("a", "aa"))
}

func minWindow(s string, t string) string {
	var ans string = ""
	var trackerT, trackerS [256]int
	for _, vl := range t {
		trackerT[vl]++
	}
	count := len(t)
	for left, right := 0, 0; right < len(s); right++ {
		ch := s[right]
		trackerS[ch]++
		if trackerS[ch] <= trackerT[ch] {
			count--
		}
		for left < right && trackerT[s[left]] < trackerS[s[left]] {
			trackerS[s[left]]--
			left++
		}
		if count == 0 && ((len(ans) == 0) || len(ans) > (right-left+1)) {
			ans = string([]rune(s)[left : right+1])
		}
	}
	return ans
}
