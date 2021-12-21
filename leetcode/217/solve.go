package main

import "fmt"

func main() {
	fmt.Println("Case 1:", containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println("Case 2:", containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println("Case 3:", containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

func containsDuplicate(nums []int) bool {
	var flag map[int]bool = make(map[int]bool)
	for _, vl := range nums {
		if _, ok := flag[vl]; ok {
			return true
		}
		flag[vl] = true
	}
	return false
}
