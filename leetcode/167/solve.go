package main

import "fmt"

func main() {
	fmt.Println("Case 1:", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println("Case 2:", twoSum([]int{2, 3, 4}, 6))
	fmt.Println("Case 3:", twoSum([]int{-1, 0}, -1))
}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		tm := numbers[left] + numbers[right]
		if tm == target {
			return []int{left + 1, right + 1}
		} else if tm < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}
