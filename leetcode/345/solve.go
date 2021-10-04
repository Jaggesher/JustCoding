package main

import "fmt"

func main() {
	fmt.Println("Case 1:", reverseVowels("hello"))
	fmt.Println("Case 2:", reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	st, end := -1, len(s)
	str := []rune(s)
	for st < end {
		st, end = st+1, end-1
		for st < end && !isVowel(str[st]) {
			st++
		}
		for end > st && !isVowel(str[end]) {
			end--
		}
		if st < end {
			str[st], str[end] = str[end], str[st]
		}
	}
	return string(str)
}

func isVowel(ch rune) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return true
	}
	if ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
		return true
	}
	return false
}
