package main

import (
	"fmt"
	"sort"
)

func main() {
	people1 := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(people1))

	people2 := [][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}
	fmt.Println(reconstructQueue(people2))

	//[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
}

func reconstructQueue(people [][]int) [][]int {
	var n int = len(people)
	var ans [][]int = make([][]int, n)
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] < people[j][0]
	})

	for i := 0; i < n; i++ {
		cnt, seen, seenCount := 0, -1, 0
		for index, vl := range people {
			if seen != vl[0] {
				seen = vl[0]
				seenCount = 0
			}
			if vl[1] == -1 {
				cnt++
				seenCount++
				continue
			} else if vl[1] == (i - cnt + seenCount) {
				ans[i] = []int{vl[0], vl[1]}
				people[index][1] = -1
				break
			}
		}
	}
	return ans
}
