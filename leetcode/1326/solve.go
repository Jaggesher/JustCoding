package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("Case 1: ", minTaps(5, []int{3, 4, 1, 1, 0, 0}))
	fmt.Println("Case 2: ", minTaps(3, []int{0, 0, 0, 0}))
	fmt.Println("Case 3: ", minTaps(7, []int{1, 2, 1, 0, 2, 1, 0, 1}))
	fmt.Println("Case 4: ", minTaps(8, []int{4, 0, 0, 0, 0, 0, 0, 0, 4}))
	fmt.Println("Case 5: ", minTaps(8, []int{4, 0, 0, 0, 4, 0, 0, 0, 4}))

}

func minTaps(n int, ranges []int) int {
	var keys []int = make([]int, 0)
	var tapMap map[int]int = make(map[int]int)
	for i := 0; i <= n; i++ {
		st, end := i-ranges[i], i+ranges[i]
		if _, ok := tapMap[st]; !ok {
			keys = append(keys, st)
			tapMap[st] = end
		} else {
			tapMap[st] = int(math.Max(float64(tapMap[st]), float64(end)))
		}
	}
	sort.Ints(keys)
	end, index := -101, 0

	for index < len(keys) && keys[index] <= 0 {
		key := keys[index]
		if tapMap[key] > end {
			end = tapMap[key]
		}
		index++
	}
	ans := 1
	for index < len(keys) && end < n {
		tm := end
		for index < len(keys) && keys[index] <= end {
			key := keys[index]
			if tapMap[key] > tm {
				tm = tapMap[key]
			}
			index++
		}
		if tm == end {
			break
		}
		ans++
		end = tm
	}

	if end >= n {
		return ans
	}
	return -1
}
