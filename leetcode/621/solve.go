package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1: ", leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
	fmt.Println("Case 2: ", leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0))
	fmt.Println("Case 3: ", leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2))
}

/***
 * Time: O(n) => where n is the length of tasks
 * Space: O(1) => since we are only using array of length 27
 */

func leastInterval(tasks []byte, n int) int {
	var count [26]int
	for _, vl := range tasks {
		count[vl-'A']++
	}
	sort.Ints(count[:])
	idealTime := (count[25] - 1) * n
	for i := 24; i >= 0; i-- {
		idealTime -= min(count[25]-1, count[i])
	}
	if idealTime < 0 {
		idealTime = 0
	}
	return len(tasks) + idealTime
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
