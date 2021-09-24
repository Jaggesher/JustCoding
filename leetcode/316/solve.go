package main

import (
	"fmt"
)

func main() {
	fmt.Println("Case 1: ", removeDuplicateLetters("bcabc"))
	fmt.Println("Case 2: ", removeDuplicateLetters("cbacdcbc"))
}

func removeDuplicateLetters(s string) string {
	var lastSeen [27]int
	var flag [27]bool
	var stack []rune = make([]rune, 0)

	for index, vl := range s {
		lastSeen[vl-'a'] = index
	}

	for index, vl := range s {
		if flag[vl-'a'] {
			continue
		}

		ln := len(stack)
		for ln > 0 {
			char := stack[ln-1]
			if char > vl && lastSeen[char-'a'] > index {
				flag[char-'a'] = false
				ln--
				continue
			}
			break
		}

		stack = stack[:ln]
		stack = append(stack, vl)
		flag[vl-'a'] = true
	}
	return string(stack)
}
