package main

import "fmt"

func main() {
	fmt.Println(longestArithSeqLength([]int{3, 6, 9, 12}))
	fmt.Println(longestArithSeqLength([]int{9, 4, 7, 2, 10}))
	fmt.Println(longestArithSeqLength([]int{20, 1, 15, 3, 10, 5, 8}))
	fmt.Println(longestArithSeqLength([]int{0, 0, 0, 0, 0}))
}

func longestArithSeqLength(nums []int) int {
	var ans int = 0
	var dictionary [][]int = make([][]int, 501)
	for i := 0; i <= 500; i++ {
		dictionary[i] = make([]int, 0)
	}
	for index, vl := range nums {
		dictionary[vl] = append(dictionary[vl], index)
	}
	n := len(nums)
	for i := 0; i < n && ans < (n-i); i++ {
		for j := i + 1; j < n; j++ {
			diff := nums[j] - nums[i]
			tm, pos := 2, j
			for next := nums[j] + diff; next >= 0 && next <= 500; next += diff {
				pos = binarySearch(dictionary[next], pos+1)
				if pos == -1 {
					break
				}
				tm++
			}
			ans = max(ans, tm)
		}
	}
	return ans
}

func binarySearch(arr []int, val int) int {
	if len(arr) == 0 {
		return -1
	}
	st, end := 0, len(arr)-1
	for st < end {
		mid := st + int((end-st)/2)
		if arr[mid] < val {
			st = mid + 1
		} else if arr[mid] >= val {
			end = mid
		}
	}
	if arr[st] < val {
		return -1
	}
	return arr[st]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
