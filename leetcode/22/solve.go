package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

var ans []string

func generateParenthesis(n int) []string {
	ans = make([]string, 0)
	TryGenerare("", n, 0, 0)
	return ans
}

func TryGenerare(sofar string, n int, open int, close int) {
	if len(sofar) == (n * 2) {
		ans = append(ans, sofar)
		return
	}
	if open < n {
		TryGenerare(sofar+"(", n, open+1, close)
	}
	if open > close {
		TryGenerare(sofar+")", n, open, close+1)
	}
}
