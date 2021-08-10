package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))

	fmt.Println(findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))

	fmt.Println(findItinerary([][]string{{"EZE", "AXA"}, {"TIA", "ANU"}, {"ANU", "JFK"}, {"JFK", "ANU"}, {"ANU", "EZE"}, {"TIA", "ANU"}, {"AXA", "TIA"}, {"TIA", "JFK"}, {"ANU", "TIA"}, {"JFK", "TIA"}}))
}

var ans []string

func findItinerary(tickets [][]string) []string {
	var pathMap map[string][]string = make(map[string][]string)
	var visited map[string]int = make(map[string]int)
	var count int = 1 + len(tickets)
	for _, vl := range tickets {
		if _, ok := pathMap[vl[0]]; !ok {
			pathMap[vl[0]] = make([]string, 0)
		}
		pathMap[vl[0]] = append(pathMap[vl[0]], vl[1])
		visited[vl[0]+vl[1]]++
	}

	ans = make([]string, count)
	for key := range pathMap {
		sort.Strings(pathMap[key])
	}
	dfs(0, count, "JFK", pathMap, visited)
	return ans
}

func dfs(n int, count int, node string, pathMap map[string][]string, visited map[string]int) bool {
	ans[n] = node
	if (n + 1) == count {
		return true
	}

	for _, vl := range pathMap[node] {
		tm := node + vl
		if visited[tm] > 0 {
			visited[tm]--
			if dfs(n+1, count, vl, pathMap, visited) {
				return true
			}
			visited[tm]++
		}
	}
	return false
}
