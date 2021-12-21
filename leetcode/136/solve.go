package main

import "fmt"

func main() {
	fmt.Println("Case 1:", singleNumber([]int{2, 2, 1}))
	fmt.Println("Case 2:", singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println("Case 3:", singleNumber([]int{1}))
}

func singleNumber(nums []int) int {
	var ans int = 0
	for _, vl := range nums {
		ans ^= vl
	}
	return ans
}
