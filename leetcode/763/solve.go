package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println("Case 2: ", partitionLabels("eccbbbbdec"))
}

/***
 * Approach: Greedy
 * Time: O(n)
 * Space: O(1)
 */
func partitionLabels(s string) []int {
	var lastSeen [27]int
	var ans []int = make([]int, 0)
	for index, vl := range s {
		lastSeen[vl-'a'] = index
	}

	st, seen := 0, lastSeen[s[0]-'a']
	for index, vl := range s {
		if index == seen {
			ans = append(ans, (index-st)+1)
			st = index + 1
			if st < len(s) {
				seen = lastSeen[s[st]-'a']
			}
		} else if seen < lastSeen[vl-'a'] {
			seen = lastSeen[vl-'a']
		}
	}

	return ans
}
