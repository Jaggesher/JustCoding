package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestSubstring("aaabb", 3))
	fmt.Println(longestSubstring("ababbc", 2))
	fmt.Println(longestSubstring("bbaaacbd", 3))
}

func longestSubstring(s string, k int) int {
	if k > len(s) {
		return 0
	}
	var countMap map[rune]int = make(map[rune]int)
	var ans int = 0

	for _, vl := range s {
		countMap[vl]++
	}

	flag := true
	var temp []rune = make([]rune, 0)
	for _, vl := range s {
		if countMap[vl] < k {
			flag = false
			if len(temp) > 0 {
				ans = int(math.Max(float64(ans), float64(longestSubstring(string(temp), k))))
				temp = temp[:0]
			}
			continue
		}
		temp = append(temp, vl)
	}

	if flag {
		return len(s)
	}

	if len(temp) > 0 {
		ans = int(math.Max(float64(ans), float64(longestSubstring(string(temp), k))))
	}

	return ans
}
