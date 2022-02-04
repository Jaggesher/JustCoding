package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", findFinalValue([]int{5, 3, 6, 1, 12}, 3))
	fmt.Println("Case 2: ", findFinalValue([]int{2, 7, 9}, 4))
}

/***
 * Time: O(n)
 * Space: O(n)
 */

func findFinalValue(nums []int, original int) int {
	var flag [1001]bool
	for _, vl := range nums {
		flag[vl] = true
	}
	for original <= 1000 && flag[original] {
		original *= 2
	}
	return original
}
