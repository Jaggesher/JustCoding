package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", isValid("()"))
	fmt.Println("Case 2: ", isValid("()[]{}"))
	fmt.Println("Case 3: ", isValid("(]"))
}

/***
 * Time: O(n)
 * Space: O(n)
 */

func isValid(s string) bool {
	var stack []rune = make([]rune, 0, len(s))
	mp := map[rune]rune{'{': '}', '(': ')', '[': ']'}

	for _, vl := range s {
		if tm, ok := mp[vl]; ok {
			stack = append(stack, tm)
		} else if len(stack) == 0 || stack[len(stack)-1] != vl {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
