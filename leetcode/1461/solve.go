package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Case 1: ", hasAllCodes("00110110", 2)) // True
	fmt.Println("Case 2: ", hasAllCodes("0110", 1))     // True
	fmt.Println("Case 3: ", hasAllCodes("0110", 2))     // true
}

/***
 * Time: O(n)
 * Space: O(m) //where m is length of all possible code.
 */
func hasAllCodes(s string, k int) bool {
	var tracker map[string]bool = make(map[string]bool)
	for n := k; n <= len(s); n++ {
		tracker[string(s[n-k:n])] = true
	}
	return len(tracker) == int(math.Pow(2, float64(k)))
}
