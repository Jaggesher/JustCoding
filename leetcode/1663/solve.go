package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", getSmallestString(3, 27))
	fmt.Println("Case 2: ", getSmallestString(5, 73))
}

/***
 * Time: O(n)
 * Space: O(1) => As we aren't using any extra space.
 */
func getSmallestString(n int, k int) string {
	var ans []byte = make([]byte, n)

	for index := 0; index < n; index++ {
		mn := k - (26 * (n - 1 - index))
		if mn <= 0 {
			ans[index] = 'a'
			k--
		} else {
			ans[index] = byte('a' + mn - 1)
			k -= mn
		}
	}

	return string(ans)
}
