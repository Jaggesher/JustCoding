package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println("Case 2: ", containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))
	fmt.Println("Case 3: ", containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
}

/***
 * Time: O(n *k)
 * Space: O(1)
 */
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for index, vl := range nums {
		st := (index - k)
		if st < 0 {
			st = 0
		}

		for st < index {
			if Abs(nums[st]-vl) <= t {
				return true
			}
			st++
		}
	}
	return false
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
