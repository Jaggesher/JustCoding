package main

import "fmt"

func main() {
	words1 := []string{"leet", "code"}
	fmt.Println(wordBreak("leetcode", words1))

	words2 := []string{"apple", "pen"}
	fmt.Println(wordBreak("applepenapple", words2))

	words3 := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak("catsandog", words3))

}

func wordBreak(s string, wordDict []string) bool {
	var dictionary [26][]int
	var cache []int = make([]int, len(s))
	for i := 0; i < 26; i++ {
		dictionary[i] = make([]int, 0)
	}
	for index, vl := range wordDict {
		tm := vl[0] - 'a'
		dictionary[tm] = append(dictionary[tm], index)
	}
	if TryBreak(s, cache, wordDict, dictionary, 0) == 1 {
		return true
	}
	return false
}

func TryBreak(s string, cache []int, wordDict []string, dictionary [26][]int, pos int) int {
	ln := len(s)
	if pos == ln {
		return 1
	}
	if cache[pos] != 0 {
		return cache[pos]
	}
	item := s[pos] - 'a'
	for _, vl := range dictionary[item] {
		if (pos + len(wordDict[vl])) <= ln {
			temp := s[pos : pos+len(wordDict[vl])]
			if wordDict[vl] == temp {
				if TryBreak(s, cache, wordDict, dictionary, pos+len(wordDict[vl])) == 1 {
					cache[pos] = 1
					return 1
				}
			}
		}
	}
	cache[pos] = -1
	return -1
}
