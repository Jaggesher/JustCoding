package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", reorganizeString("aab"))
	fmt.Println("Case 2: ", reorganizeString("aaab"))
}

/***
 * Time: O(n)
 * Space: O(26) => O(1)
 */

func reorganizeString(s string) string {
	var count [27]int
	for _, vl := range s { // O(n)
		count[vl-'a']++
	}
	var ans []byte = make([]byte, len(s))

	lastUsed := -1
	for index := range ans { // O(n) => O(n * 26) => O(n)
		mx := -1
		for i := 0; i < 26; i++ { // O(26)
			if count[i] > 0 && (mx < 0 || count[i] > count[mx]) && i != lastUsed {
				mx = i
			}
		}
		if mx == -1 {
			return ""
		}
		lastUsed = mx
		ans[index] = byte('a' + mx)
		count[mx]--
	}

	return string(ans)
}
