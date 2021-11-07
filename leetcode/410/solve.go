package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", splitArray([]int{7, 2, 5, 10, 8}, 2))
	fmt.Println("Case 2: ", splitArray([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println("Case 2: ", splitArray([]int{1, 4, 4}, 3))
}

func splitArray(nums []int, m int) int {
	st, end := 0, 0
	for _, vl := range nums {
		end += vl
		if st < vl {
			st = vl
		}
	}
	for st <= end {
		mid := (st + end + 1) / 2
		cut := possibleCuts(nums, mid)
		if cut <= m {
			end = mid - 1
		} else {
			st = mid + 1
		}
	}
	return st
}

func possibleCuts(nums []int, mx int) int {
	sum, cut := 0, 1
	for _, vl := range nums {
		if (sum + vl) > mx {
			cut++
			sum = vl
		} else {
			sum += vl
		}
	}
	return cut
}
