package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", peakIndexInMountainArray([]int{0, 1, 0}))
	fmt.Println("Case 2: ", peakIndexInMountainArray([]int{0, 2, 1, 0}))
	fmt.Println("Case 3: ", peakIndexInMountainArray([]int{0, 10, 5, 2}))
}

func peakIndexInMountainArray(arr []int) int {
	st, end := 0, len(arr)
	for st < end {
		mid := st + (end-st)/2
		if arr[mid] < arr[mid+1] {
			st = mid + 1
		} else {
			end = mid
		}
	}
	return st
}
