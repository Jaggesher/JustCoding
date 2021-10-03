package main

import "fmt"

func main() {
	fmt.Println("Case 1:", summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println("Case 2:", summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	fmt.Println("Case 3:", summaryRanges([]int{}))
	fmt.Println("Case 4:", summaryRanges([]int{1, 2, 3, 4}))
}

func summaryRanges(nums []int) []string {
	var ans []string = make([]string, 0)
	index, ln := 0, len(nums)
	for index < ln {
		end := index
		for (end+1) < ln && (nums[end]+1) == nums[end+1] {
			end++
		}
		if index == end {
			ans = append(ans, fmt.Sprintf("%d", nums[index]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", nums[index], nums[end]))
		}
		index = end + 1
	}
	return ans
}
