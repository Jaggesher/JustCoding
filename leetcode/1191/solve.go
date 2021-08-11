package main

import "fmt"

func main() {
	arr1 := []int{-5, 6, -7, 10, 12, -2}
	arr2 := []int{1, 2}
	arr3 := []int{1, -2, 1}
	arr4 := []int{-1, -2}
	arr5 := []int{-5, -2, 0, 0, 3, 9, -2, -5, 4}
	arr6 := []int{-1, 1, 2, 3, -8, 4}
	fmt.Println(kConcatenationMaxSum(arr1, 1))
	fmt.Println(kConcatenationMaxSum(arr1, 2))
	fmt.Println(kConcatenationMaxSum(arr1, 3))
	fmt.Println(kConcatenationMaxSum(arr2, 3))
	fmt.Println(kConcatenationMaxSum(arr3, 5))
	fmt.Println(kConcatenationMaxSum(arr4, 7))
	fmt.Println(kConcatenationMaxSum(arr5, 5))
	fmt.Println(kConcatenationMaxSum(arr6, 2))
}

func kConcatenationMaxSum(arr []int, k int) int {
	const mod = 1000000007
	var ans, total int = 0, 0
	ans, total = PossibleMaxAndSum(arr)
	if k == 1 {
		return ans % mod
	}
	if total > 0 {
		ansTm := ((total * k) - PossibleNegativeFromBothSide(arr))
		ans = max(ansTm, ans)
	} else {
		ans = max(ans, PossiblePositiveFromBothSide(arr))
	}
	return ans % mod
}

// Time: O(2n)
// Space: O(1)
func PossibleNegativeFromBothSide(arr []int) int {
	fNeg := 0
	tm := 0
	for i := 0; i < len(arr); i++ {
		tm += arr[i]
		if fNeg > tm {
			fNeg = tm
		}
	}
	lNeg, tm := 0, 0
	for i := len(arr) - 1; i >= 0; i-- {
		tm += arr[i]
		if lNeg > tm {
			lNeg = tm
		}
	}
	return fNeg + lNeg
}

//Time: O(2n)
//Space: O(1)
func PossiblePositiveFromBothSide(arr []int) int {
	fMx, fIndex := 0, -1
	tm := 0
	for i := 0; i < len(arr); i++ {
		tm += arr[i]
		if tm > fMx {
			fMx = tm
			fIndex = i
		}
	}
	lMx, tm := 0, 0
	for i := len(arr) - 1; i > fIndex; i-- {
		tm += arr[i]
		if tm > lMx {
			lMx = tm
		}
	}
	return fMx + lMx
}

//Time: O(n)
//Space: O(1)
func PossibleMaxAndSum(arr []int) (mx int, total int) {
	mx, tm, total := 0, 0, 0
	for i := 0; i < len(arr); i++ {
		if tm < 0 {
			tm = 0
		}
		tm += arr[i]
		mx = max(mx, tm)
		total += arr[i]
	}
	return mx, total
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
