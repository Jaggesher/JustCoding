package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("Aabb"))
}

func frequencySort(s string) string {
	var dictionary map[rune]int = make(map[rune]int)
	for _, vl := range s { // O(n)
		dictionary[vl]++
	}
	type Counter struct {
		char  rune
		count int
	}
	temp := make([]Counter, 0)
	for key, vl := range dictionary { // O(n)
		temp = append(temp, Counter{char: key, count: vl})
	}
	sort.Slice(temp, func(i, j int) bool { //O(n(log(n)))
		return temp[i].count > temp[j].count
	})
	var ans string = ""
	for _, vl := range temp { // O(n) as total
		for vl.count > 0 {
			ans += string(vl.char)
			vl.count--
		}
	}
	return ans
}
