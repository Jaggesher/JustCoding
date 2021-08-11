package main

import (
	"fmt"
	"sort"
)

func main() {
	times1 := [][]int{{3, 10}, {1, 5}, {2, 6}, {4, 10}}
	fmt.Println(smallestChair(times1, 0))

	times2 := [][]int{{1, 4}, {2, 3}, {4, 6}}
	fmt.Println(smallestChair(times2, 1))

}

func smallestChair(times [][]int, targetFriend int) int {
	var data [][]int = make([][]int, 0)
	for _, vl := range times {
		if vl[0] <= times[targetFriend][0] {
			data = append(data, vl)
		}
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})
	var seat []int = make([]int, len(data))
	for i := 0; i < len(data); i++ {
		seat[i] = -1
	}
	ans := -1
	for _, item := range data {
		st, end := item[0], item[1]
		for i := 0; i < len(seat); i++ {
			if seat[i] == -1 || seat[i] <= st {
				seat[i] = end
				ans = i
				break
			}
		}
	}
	return ans
}
