package main

import "fmt"

func main() {
	fmt.Println("Case 1:", missingNumber([]int{3, 0, 1}))
	fmt.Println("Case 2:", missingNumber([]int{0, 1}))
	fmt.Println("Case 3:", missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println("Case 4:", missingNumber([]int{0}))
}

func missingNumber(nums []int) int {
	var missing int = len(nums)
	for i, vl := range nums {
		missing ^= (i ^ vl)
	}
	return missing
}
