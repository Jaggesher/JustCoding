package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", maxScoreIndices([]int{0, 0, 1, 0}))
	fmt.Println("Case 2: ", maxScoreIndices([]int{0, 0, 0}))
	fmt.Println("Case 3: ", maxScoreIndices([]int{1, 1}))
}

/***
 * TIme: O(n) + O(n) => O(n)
 * Space: O(1)
 */
func maxScoreIndices(nums []int) []int {
	var zeros, ones int = 0, 0
	for _, vl := range nums {
		ones += vl
	}
	tempMax := ones
	ans := []int{0}
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] == 0 {
			zeros++
		} else {
			ones--
		}
		if tempMax < (zeros + ones) {
			tempMax = (zeros + ones)
			ans = []int{i}
		} else if tempMax == (zeros + ones) {
			ans = append(ans, i)
		}
	}
	return ans
}
