package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", permute([]int{1, 2, 3}))
	fmt.Println("Case 2: ", permute([]int{0, 1}))
}

/***
 * Time: O(n!)
 * Space O(n! * m) => O(n!)
 */

func permute(nums []int) [][]int {
	var ans [][]int = make([][]int, 0)

	var recursion func(int)
	recursion = func(index int) {
		if index >= len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			ans = append(ans, temp)
			return
		}
		for i := index; i < len(nums); i++ {
			nums[index], nums[i] = nums[i], nums[index]
			recursion(index + 1)
			nums[index], nums[i] = nums[i], nums[index]
		}
	}
	recursion(0)
	return ans
}
