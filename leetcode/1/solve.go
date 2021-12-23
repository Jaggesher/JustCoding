package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println("Case 2: ", twoSum([]int{3, 2, 4}, 6))
	fmt.Println("Case 3: ", twoSum([]int{3, 3}, 6))
}

/***
 * Time: O(n)
 * Space: O(n)
 */

func twoSum(nums []int, target int) []int {
	var seen map[int]int = make(map[int]int)
	for i, vl := range nums {
		required := target - vl
		if index, ok := seen[required]; ok {
			return []int{index, i}
		}
		seen[vl] = i
	}
	return []int{-1, -1}
}
