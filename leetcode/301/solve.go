package main

import "fmt"

func main() {
	fmt.Println("Case 1:", removeInvalidParentheses("()())()"))
	fmt.Println("Case 2:", removeInvalidParentheses("(a)())()"))
	fmt.Println("Case 3:", removeInvalidParentheses(")("))
	fmt.Println("Case 4:", removeInvalidParentheses(")o(v("))
}

func removeInvalidParentheses(s string) []string {
	var set map[string]bool = make(map[string]bool)
	backtrack(0, 0, len(s)-findRemoveable(s), []byte{}, s, set)
	var ans []string = make([]string, 0, len(set))
	for key := range set {
		ans = append(ans, key)
	}
	return ans
}

func backtrack(index, open, ln int, running []byte, s string, set map[string]bool) {
	if index >= len(s) {
		if ln == len(running) && open == 0 {
			set[string(running)] = true
		}
		return
	}
	possible := len(running) + len(s) - index
	if possible < ln {
		return
	}

	if s[index] != '(' && s[index] != ')' {
		backtrack(index+1, open, ln, append(running, s[index]), s, set)
		return
	}
	backtrack(index+1, open, ln, running, s, set)
	if s[index] == '(' {
		backtrack(index+1, open+1, ln, append(running, s[index]), s, set)
	} else if open > 0 {
		backtrack(index+1, open-1, ln, append(running, s[index]), s, set)
	}
}

func findRemoveable(s string) int {
	minRemove, temp := 0, 0
	for _, vl := range s {
		if vl == '(' {
			temp++
		} else if vl == ')' {
			temp--
		}
		if temp < 0 {
			minRemove -= temp
			temp = 0
		}
	}
	return minRemove + temp
}
