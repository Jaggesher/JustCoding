package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))
	fmt.Println(countSmaller([]int{-1}))
	fmt.Println(countSmaller([]int{-1, -1}))
}

func countSmaller(nums []int) []int {
	var ans []int = make([]int, len(nums))
	var order map[int]int = make(map[int]int)
	copy(ans, nums)
	sort.Ints(nums)
	cnt := 1
	for _, vl := range nums {
		if order[vl] == 0 {
			order[vl] = cnt
			cnt++
		}
	}
	var BIT []int = make([]int, cnt)
	for i := len(nums) - 1; i >= 0; i-- {
		index := order[ans[i]]
		ans[i] = SearchBIT(BIT, index-1)
		UpdateBIT(BIT, index, 1)
	}
	return ans
}

func UpdateBIT(BIT []int, index int, val int) {
	for index < len(BIT) {
		BIT[index] += val
		index += (index & (-index))
	}
}

func SearchBIT(BIT []int, index int) int {
	result := 0
	for index > 0 {
		result += BIT[index]
		index -= (index & -index)
	}
	return result
}
