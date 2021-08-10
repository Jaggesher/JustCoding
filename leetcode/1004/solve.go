package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))

	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))

	fmt.Println(longestOnes([]int{0, 0, 0, 1}, 0))

}

func longestOnes(nums []int, k int) int {
	var ans int = 0
	st, tempSum := 0, 0

	for end := 0; end < len(nums); end++ {
		tempSum += nums[end]

		for ((end - st + 1) - tempSum) > k {
			tempSum -= nums[st]
			st++
		}
		if ans < (end - st + 1) {
			ans = (end - st + 1)
		}
	}
	return ans
}
