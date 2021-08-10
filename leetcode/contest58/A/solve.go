package main

import "fmt"

func main() {
	fmt.Println(makeFancyString("leeetcode"))
	fmt.Println(makeFancyString("aaabaaaa"))
	fmt.Println(makeFancyString("aab"))
}

func makeFancyString(s string) string {
	var ans []rune = make([]rune, 0)
	lastChar, count := 'A', 0
	for _, vl := range s {
		if lastChar != vl {
			lastChar = vl
			count = 1
		} else {
			count++
		}

		if count < 3 {
			ans = append(ans, lastChar)
		}
	}
	return string(ans)
}
