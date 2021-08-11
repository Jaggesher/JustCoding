package main

import (
	"fmt"
)

func main() {
	var s string = "dsahjpjauf"
	var words []string = []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}
	fmt.Println(numMatchingSubseq(s, words))

}

func numMatchingSubseq(s string, words []string) int {
	var ans int = 0
	var dictionary [][]int = make([][]int, 26)
	for i, vl := range s {
		dictionary[vl-'a'] = append(dictionary[vl-'a'], i)
	}
	for _, word := range words {
		flag := true
		pos := -1
		for _, vl := range word {
			tm := binarySearch(dictionary[vl-'a'], pos+1)
			if tm == -1 {
				flag = false
				break
			}
			pos = tm
		}
		if flag {
			ans++
		}
	}
	return ans
}

func binarySearch(arr []int, num int) int {
	if len(arr) == 0 {
		return -1
	}
	st, end := 0, len(arr)-1
	for st < end {
		mid := st + int(((end - st) / 2))
		if arr[mid] == num {
			return num
		} else if arr[mid] > num {
			end = mid
		} else {
			st = mid + 1
		}
	}
	if arr[st] < num {
		return -1
	}
	return arr[st]
}
