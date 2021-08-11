package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{0}))
	fmt.Println(rob([]int{5, 1, 1, 10}))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var cache []int = make([]int, len(nums))
	resetArr(cache, -1)
	from0 := robBest(nums, 0, cache, len(nums)-1)
	resetArr(cache, -1)
	from1 := robBest(nums, 1, cache, len(nums))
	if from0 > from1 {
		return from0
	}
	return from1
}

func resetArr(arr []int, vl int) {
	for i := range arr {
		arr[i] = vl
	}
}

func robBest(nums []int, index int, cache []int, ln int) int {
	if index >= ln {
		return 0
	}
	if cache[index] != -1 {
		return cache[index]
	}
	take := nums[index] + robBest(nums, index+2, cache, ln)
	skip := robBest(nums, index+1, cache, ln)
	cache[index] = skip
	if take > skip {
		cache[index] = take
	}
	return cache[index]
}
