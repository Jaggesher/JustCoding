package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Case 1: ", numDecodings("12"))
	fmt.Println("Case 2: ", numDecodings("226"))
	fmt.Println("Case 3: ", numDecodings("06"))
	fmt.Println("Case 4: ", numDecodings("1212"))
}

/***
 * Time: O(n) => where n is the length of string s
 * Space: O(1) => As we are using constant space
 */
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	oneBack, twoBack := 1, 1

	for i := 1; i < len(s); i++ {
		current := 0
		if s[i] != '0' {
			current += oneBack
		}
		tm, _ := strconv.Atoi(s[i-1 : i+1])
		if tm > 9 && tm < 27 {
			current += twoBack
		}
		oneBack, twoBack = current, oneBack
	}
	return oneBack
}
