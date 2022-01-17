package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println("Case 2: ", lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println("Case 3: ", lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}

/***
 * Time: O(N*N) => N is the length of nums array
 * Space: O(N) => N is the length of nums array.
 */
func lengthOfLIS(nums []int) int {
	var ans int = 0
	var flag []int = make([]int, len(nums))
	for index, vl := range nums {
		tm := 0
		for i := 0; i < index; i++ {
			if vl > nums[i] && tm < flag[i] {
				tm = flag[i]
			}
		}
		flag[index] = tm + 1
		if ans < flag[index] {
			ans = flag[index]
		}
	}
	return ans
}
