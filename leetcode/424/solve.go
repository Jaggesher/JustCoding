package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", characterReplacement("ABAB", 2))
	fmt.Println("Case 2: ", characterReplacement("AABABBA", 1))
}

/***
 * TIME: O(n)+O(n) => O(n)
 * SPACE: O(1)
 */
func characterReplacement(s string, k int) int {
	var flag [26]int
	var ans int
	begin, currentMax := 0, 0
	for index, vl := range s {
		flag[vl-'A']++
		if currentMax < flag[vl-'A'] {
			currentMax = flag[vl-'A']
		}
		for (index - begin - currentMax + 1) > k {
			flag[s[begin]-'A']--
			begin++
		}
		if ans < (index - begin + 1) {
			ans = index - begin + 1
		}
	}

	return ans
}
