package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(medianSlidingWindow([]int{1, 2, 3, 4}, 3))
	fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(medianSlidingWindow([]int{1, 2, 3, 4, 2, 3, 1, 4, 2}, 3))
	fmt.Println(medianSlidingWindow([]int{1, 4, 2, 3}, 4))
	fmt.Println(medianSlidingWindow([]int{7, 0, 3, 9, 9, 9, 1, 7, 2, 3}, 6))
}

func medianSlidingWindow(nums []int, k int) []float64 {
	var ans []float64 = make([]float64, 0)
	tracker := make([]int, k)
	for i := 0; i < k; i++ {
		tracker[i] = nums[i]
	}
	sort.Ints(tracker)
	ans = append(ans, median(tracker))

	for i := k; i < len(nums); i++ {
		tracker = insert(tracker, nums[i])
		tracker = remove(tracker, nums[i-k])
		ans = append(ans, median(tracker))
	}
	return ans
}

func median(arr []int) float64 {
	ln := len(arr)
	mid := int((ln - 1) / 2)
	if (ln & 1) == 1 {
		return float64(arr[mid])
	}
	return float64((arr[mid] + arr[mid+1])) / 2.0
}

func insert(num []int, vl int) []int {
	st, end := 0, len(num)-1

	if num[st] >= vl {
		tm := []int{vl}
		return append(tm, num...)
	} else if num[end] <= vl {
		return append(num, vl)
	}

	for st < end {
		mid := st + int((end-st)/2)
		if num[mid] <= vl {
			st = mid + 1
		} else {
			end = mid
		}
	}

	temp := make([]int, st)
	copy(temp, num[0:st])
	temp = append(temp, vl)
	temp = append(temp, num[st:]...)
	return temp
}

func remove(num []int, vl int) []int {
	st, end := 0, len(num)-1
	if num[st] == vl {
		return num[1:]
	} else if num[end] == vl {
		return num[0:end]
	}
	for st <= end {
		mid := st + int((end-st)/2)
		if num[mid] == vl {
			temp := num[0:mid]
			if mid+1 < len(num) {
				temp = append(temp, num[mid+1:]...)
			}
			return temp
		} else if num[mid] < vl {
			st = mid + 1
		} else {
			end = mid - 1
		}
	}
	fmt.Println("Can't able to find=", vl, "Inside =", num, "Last Status= ", st)
	return num
}
