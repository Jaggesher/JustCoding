package main

import "fmt"

func main() {
	fmt.Println("Case1:", checkPalindromeFormation("x", "y"))
	fmt.Println("Case2:", checkPalindromeFormation("ulacfd", "jizalu"))
	fmt.Println("Case3:", checkPalindromeFormation("xbdef", "xecab"))
	fmt.Println("Case3:", checkPalindromeFormation("aaabart", "ghxcfab"))
}

func checkPalindromeFormation(a string, b string) bool {
	st, end := 0, len(a)-1
	if checkPalindrome(a, st, end) || checkPalindrome(b, st, end) {
		return true
	}
	ab, ba := true, true
	for st < end && (ab || ba) {
		if a[st] != b[end] {
			ab = false
		}
		if b[st] != a[end] {
			ba = false
		}
		flag := checkPalindrome(b, st+1, end-1) || checkPalindrome(a, st+1, end-1)
		if ab && flag {
			return true
		}
		if ba && flag {
			return true
		}
		st++
		end--
	}
	return false
}

func checkPalindrome(str string, st int, end int) bool {
	for st < end {
		if str[st] != str[end] {
			return false
		}
		st++
		end--
	}
	return true
}
