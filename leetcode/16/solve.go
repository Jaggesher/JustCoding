package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums []int = []int{-1, 0, 1, 1, 55}
	var target int = 3
	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var ans int = nums[0] + nums[1] + nums[2]
	var ln int = len(nums)
	for i := 0; i < ln; i++ {
		for j := i + 1; j < ln-1; j++ {
			tm := nums[i] + nums[j]
			st, end := j+1, ln-1
			for st < end {
				mid := st + int((end-st)/2)
				if mid == st {
					if abs(target-tm-nums[mid]) < abs(target-tm-nums[end]) {
						end = mid
					} else {
						st = end
					}
				} else if (tm + nums[mid]) >= target {
					end = mid
				} else {
					st = mid
				}
			}
			dif1, dif2 := abs(target-ans), abs(target-tm-nums[st])

			//fmt.Println(i, j, st, dif1, dif2)
			if dif1 > dif2 {
				ans = tm + nums[st]
			}

		}
	}
	return ans
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
