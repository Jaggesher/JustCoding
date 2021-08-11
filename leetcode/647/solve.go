package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
}

func countSubstrings(s string) int {
	var ans int = 0
	for i := 0; i < len(s); i++ {
		ans += tryExtend(s, i, i)
		ans += tryExtend(s, i, i+1)
	}
	return ans
}

func tryExtend(s string, left, right int) int {
	var ans int = 0
	for ; left >= 0 && right < len(s); ans, left, right = ans+1, left-1, right+1 {
		if s[left] != s[right] {
			break
		}
	}
	return ans
}
