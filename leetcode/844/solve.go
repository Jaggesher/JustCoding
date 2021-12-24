package main

import "fmt"

func main() {
	fmt.Println("Case 1:", backspaceCompare("ab#c", "ad#c"))
}

/***
 * Time: O(n)
 * Space: O(1)
 */

func backspaceCompare(s string, t string) bool {
	sEnd, tEnd := len(s)-1, len(t)-1
	for sEnd >= 0 || tEnd >= 0 {
		for skip := 0; (sEnd >= 0) && (s[sEnd] == '#' || skip > 0); sEnd-- {
			if s[sEnd] == '#' {
				skip++
			} else {
				skip--
			}
		}

		for skip := 0; tEnd >= 0 && (t[tEnd] == '#' || skip > 0); tEnd-- {
			if t[tEnd] == '#' {
				skip++
			} else {
				skip--
			}
		}

		if sEnd < 0 && tEnd < 0 {
			return true
		}

		if sEnd < 0 || tEnd < 0 || s[sEnd] != t[tEnd] {
			return false
		}
		sEnd, tEnd = sEnd-1, tEnd-1
	}
	return true
}
