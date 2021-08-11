package main

import "fmt"

func main() {
	rains1 := []int{1, 2, 3, 4}
	fmt.Println("Case1: ", avoidFlood(rains1))

	rains2 := []int{1, 2, 0, 0, 2, 1}
	fmt.Println("Case2: ", avoidFlood(rains2))

	rains3 := []int{1, 2, 0, 1, 2}
	fmt.Println("Case3: ", avoidFlood(rains3))

	rains4 := []int{69, 0, 0, 0, 69}
	fmt.Println("Case4: ", avoidFlood(rains4))

	rains5 := []int{10, 20, 20}
	fmt.Println("Case5: ", avoidFlood(rains5))

	rains6 := []int{0, 1, 1}
	fmt.Println("Case6: ", avoidFlood(rains6))

}

func avoidFlood(rains []int) []int {
	var ans []int = make([]int, len(rains))
	var lakeStatus map[int]int = make(map[int]int)
	var freeDays []int = make([]int, 0)
	for day, vl := range rains {
		if vl == 0 {
			freeDays = append(freeDays, day+1)
			ans[day] = 1
		} else {
			if lakeStatus[vl] == 0 {
				ans[day] = -1
				lakeStatus[vl] = day + 1
			} else {
				possibleDryDay := binarySearch(freeDays, lakeStatus[vl])
				if possibleDryDay == -1 {
					return make([]int, 0)
				} else {
					ans[freeDays[possibleDryDay]-1] = vl
					ans[day] = -1
					lakeStatus[vl] = day + 1

					copy(freeDays[possibleDryDay])
					tm := make([]int, 0)
					tm = append(tm, freeDays[:possibleDryDay]...)
					tm = append(tm, freeDays[possibleDryDay+1:]...)
					freeDays = tm
				}
			}
		}
	}
	return ans
}

func binarySearch(arr []int, val int) int {
	st, end := 0, len(arr)-1
	for st < end {
		mid := st + int((end-st)/2)
		if arr[mid] <= val {
			st = mid + 1
		} else {
			end = mid
		}
	}
	if len(arr) == 0 || arr[st] <= val {
		return -1
	}
	return st
}
