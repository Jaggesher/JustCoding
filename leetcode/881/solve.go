package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1:", numRescueBoats([]int{1, 2}, 3))
	fmt.Println("Case 2:", numRescueBoats([]int{3, 2, 2, 1}, 3))
	fmt.Println("Case 3:", numRescueBoats([]int{3, 5, 3, 4}, 5))
}

/***
 * Approach: Greedy
 * Time: O(n Log n) + O(n) => O(n Log n)
 * Space: O(n)
 */

func numRescueBoats(people []int, limit int) int {
	var ans int = 0
	sort.Ints(people)
	st, end := 0, len(people)-1
	for st <= end {
		if people[st]+people[end] <= limit {
			st++
		}
		ans, end = ans+1, end-1
	}
	return ans
}
