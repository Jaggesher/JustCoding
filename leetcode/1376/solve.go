package main

import "fmt"

func main() {
	fmt.Println("Case 1:")
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	var mp [][]int = make([][]int, n)
	for i, vl := range manager {
		if vl == -1 {
			continue
		}
		mp[vl] = append(mp[vl], i)
	}
	return dfs(headID, mp, informTime)
}

func dfs(node int, mp [][]int, informTime []int) int {
	if len(mp[node]) == 0 {
		return 0
	}
	mx := 0
	for _, vl := range mp[node] {
		tm := dfs(vl, mp, informTime)
		if mx < tm {
			mx = tm
		}
	}
	return mx + informTime[node]
}
