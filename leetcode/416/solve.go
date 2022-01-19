package main

import "fmt"

func main() {
	fmt.Println("Case 1:", canPartition([]int{1, 5, 11, 5}))
	fmt.Println("Case 2:", canPartition([]int{1, 2, 3, 5}))
	fmt.Println("Case 3:", canPartition([]int{100, 100, 100, 100, 100, 100, 100, 100}))
}

func canPartition(nums []int) bool {
	var sum int = 0
	for _, vl := range nums {
		sum += vl
	}

	if sum&1 == 1 {
		return false
	}
	sum /= 2
	var memo []int = make([]int, sum+1)
	return tryMake(0, sum, memo, nums) == 1
}

func tryMake(index, left int, memo, nums []int) int {
	if memo[left] != 0 {
		return memo[left]
	}
	if left == 0 {
		return 1
	}
	if index >= len(nums) {
		return 2
	}

	for i := index; i < len(nums) && memo[left] != 1; i++ {
		if nums[i] <= left {
			memo[left] = tryMake(i+1, left-nums[i], memo, nums)
		}
	}

	if memo[left] != 1 {
		memo[left] = tryMake(index+1, left, memo, nums)
	}

	return memo[left]
}
