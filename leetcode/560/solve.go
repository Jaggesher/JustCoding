package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum([]int{3}, 1))
	fmt.Println(subarraySum([]int{3, 1}, 1))
	fmt.Println(subarraySum([]int{-1, -1, 1}, 0))
	fmt.Println(subarraySum([]int{28, 54, 7, -70, 22, 65, -6}, 100))
	fmt.Println(subarraySum([]int{0, 0, 0}, 0))
	fmt.Println(subarraySum([]int{1, 2, 1, 2, 1}, 3))
	fmt.Println(subarraySum([]int{1, -1, 0}, 0))
}

func subarraySum(nums []int, k int) int {
	var ans int = 0
	var discovered map[int]int = make(map[int]int)
	discovered[0] = 1
	sum := 0
	for _, vl := range nums {
		sum += vl
		if _, ok := discovered[sum-k]; ok {
			ans += discovered[sum-k]
		}
		discovered[sum]++
	}
	return ans
}
