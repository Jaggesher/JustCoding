package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//var nums []int = []int{1, 0, -1, 0, -2, 2}
	//var nums []int = []int{2, 2, 2, 2, 2}
	var nums []int = []int{-3, -2, -1, 0, 0, 1, 2, 3}
	var target int = 0
	fmt.Println(fourSum(nums, target))
}
func fourSum(nums []int, target int) [][]int {
	var ans [][]int = make([][]int, 0)
	var flag map[string]bool = make(map[string]bool)
	sort.Ints(nums)
	ln := len(nums)
	for i := 0; i < ln; i++ {
		for j := i + 1; j < ln; j++ {
			st, end, rest := j+1, ln-1, target-(nums[i]+nums[j])

			for st < end {
				tm := nums[st] + nums[end]
				if tm == rest {
					str := strconv.Itoa(nums[i]) + "=" + strconv.Itoa(nums[j]) + "=" + strconv.Itoa(nums[st])
					if !flag[str] {
						ans = append(ans, []int{nums[i], nums[j], nums[st], nums[end]})
						flag[str] = true
					}
					st++
					end--
				} else if tm < rest {
					st++
				} else {
					end--
				}
			}
		}
	}
	return ans
}
