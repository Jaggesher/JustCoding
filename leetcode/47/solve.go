package main

import "fmt"

func main() {
	fmt.Println("Case 1:", permuteUnique([]int{1, 1, 2}))
	fmt.Println("Case 2:", permuteUnique([]int{1, 2, 3}))
}

/***
 * Time: O(n!) // if all unique
 * Space: O(n! * n) // if all unique
 */

func permuteUnique(nums []int) [][]int {
	var ans [][]int = make([][]int, 0)
	var recursion func(int)

	recursion = func(index int) {
		if index >= len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			ans = append(ans, temp)
			return
		}

		unique := make(map[int]int)
		for i := index; i < len(nums); i++ {
			unique[nums[i]] = i
		}

		for _, vl := range unique {
			nums[index], nums[vl] = nums[vl], nums[index]
			recursion(index + 1)
			nums[index], nums[vl] = nums[vl], nums[index]
		}
	}
	recursion(0)
	return ans
}
