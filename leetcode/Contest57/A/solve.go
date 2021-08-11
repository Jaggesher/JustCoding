package main

import "fmt"

func main() {
	fmt.Println(areOccurrencesEqual("abacbc"))
	fmt.Println(areOccurrencesEqual("aaabb"))
}

func areOccurrencesEqual(s string) bool {
	var mp map[rune]int = make(map[rune]int, 0)
	var tm int = 0
	for _, vl := range s {
		mp[vl]++
		tm = mp[vl]
	}
	for _, vl := range mp {
		if vl != tm {
			return false
		}
	}
	return true
}
