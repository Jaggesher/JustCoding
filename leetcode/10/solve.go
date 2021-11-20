package main

import "fmt"

func main() {
	fmt.Println("Case 1:", isMatch("aa", "a"))
	fmt.Println("Case 2:", isMatch("aa", "a*"))
	fmt.Println("Case 3:", isMatch("ab", ".*"))
	fmt.Println("Case 4:", isMatch("aab", "c*a*b"))
	fmt.Println("Case 5:", isMatch("mississippi", "mis*is*p*."))
}

func isMatch(s string, p string) bool {
	var memo [22][32]int
	return solve(0, 0, s, p, &memo)
}

func solve(sI, pI int, s, p string, memo *[22][32]int) bool {
	if sI == len(s) && pI == len(p) {
		return true
	}
	if memo[sI][pI] != 0 {
		return memo[sI][pI] == 2
	}
	if (pI+1) < len(p) && p[pI+1] == '*' {
		i := 0
		for true {
			if solve(sI+i, pI+2, s, p, memo) {
				memo[sI][pI] = 2
				return true
			}
			if (sI+i) >= len(s) || p[pI] != '.' && s[sI+i] != p[pI] {
				break
			}
			i++
		}
	} else if sI < len(s) && pI < len(p) && (p[pI] == '.' || p[pI] == s[sI]) {
		if solve(sI+1, pI+1, s, p, memo) {
			memo[sI][pI] = 2
			return true
		}
	}
	memo[sI][pI] = 1
	return false
}
