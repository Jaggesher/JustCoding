package main

import (
	"fmt"
)

func main() {
	fmt.Println("Case 1:", isIsomorphic("egg", "add"))
	fmt.Println("Case 2:", isIsomorphic("foo", "bar"))
	fmt.Println("Case 3:", isIsomorphic("bbbaaaba", "aaabbbba"))
}

func isIsomorphic(s string, t string) bool {
	var trackerS, trackerT map[byte]byte = make(map[byte]byte), make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		_, ok1 := trackerS[s[i]]
		_, ok2 := trackerT[t[i]]
		if ok1 != ok2 {
			return false
		} else if !ok1 {
			trackerS[s[i]] = t[i]
			trackerT[t[i]] = s[i]
		} else if (trackerS[s[i]] != t[i]) || (trackerT[t[i]] != s[i]) {
			return false
		}
	}
	return true
}
