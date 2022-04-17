package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", firstUniqChar("leetcode"))
	fmt.Println("Case 2: ", firstUniqChar("loveleetcode"))
	fmt.Println("Case 3: ", firstUniqChar("aabb"))
}

/***
 * Time: O(n)
 * Space: O(1)
 */
func firstUniqChar(s string) int {
	var ans int
	var status [27]int

	for index, vl := range s {
		if status[vl-'a'] == 0 {
			status[vl-'a'] = index + 1
		} else {
			status[vl-'a'] = -1
		}
	}
	for _, vl := range status {
		if vl > 0 && (ans == 0 || vl < ans) {
			ans = vl
		}
	}

	return ans - 1
}
