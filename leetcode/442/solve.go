package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println("Case 2: ", findDuplicates([]int{1, 1, 2}))
	fmt.Println("Case 3: ", findDuplicates([]int{1}))
}

/***
 * Time: O(n)
 * Space: O(n)
 */

func findDuplicates(nums []int) []int {
	var ans []int
	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}
	for _, num := range nums {
		num = abs(num) - 1
		if nums[num] < 0 {
			ans = append(ans, num+1)
		} else {
			nums[num] = -nums[num]
		}
	}
	return ans
}
