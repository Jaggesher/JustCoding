package main

import "fmt"

func main() {
	fmt.Println("Case 1:", checkInclusion("ab", "eidbaooo"))
	fmt.Println("Case 2:", checkInclusion("ab", "eidboaoo"))
	fmt.Println("Case 3:", checkInclusion("abc", "ccccbbbbaaaa"))
}

func checkInclusion(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}
	var flag [27]int
	for i := 0; i < l1; i++ {
		flag[(s1[i]-'a')]++
		flag[(s2[i]-'a')]--
	}
	if check(flag) {
		return true
	}
	for i := l1; i < l2; i++ {
		flag[s2[i-l1]-'a']++
		flag[s2[i]-'a']--
		if check(flag) {
			return true
		}
	}
	return false
}

func check(arr [27]int) bool {
	for _, vl := range arr {
		if vl != 0 {
			return false
		}
	}
	return true
}
