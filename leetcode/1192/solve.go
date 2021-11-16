package main

import "fmt"

func main() {
	fmt.Println("Case 1:", criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}))
	fmt.Println("Case 2:", criticalConnections(2, [][]int{{0, 1}}))
	fmt.Println("Case 3:", criticalConnections(6, [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 0}, {2, 4}, {3, 4}, {3, 5}}))
}

func criticalConnections(n int, connections [][]int) [][]int {
	var ans [][]int = make([][]int, 0)

	var graph [][]int = make([][]int, n)
	var ranks []int = make([]int, n)

	for _, el := range connections {
		graph[el[0]] = append(graph[el[0]], el[1])
		graph[el[1]] = append(graph[el[1]], el[0])
	}
	dfs(0, graph, ranks, 1, &ans)
	return ans
}

func dfs(node int, graph [][]int, ranks []int, discoveredRank int, ans *[][]int) int {
	ranks[node] = discoveredRank
	minNeighboredRank := discoveredRank + 1
	for _, next := range graph[node] {
		if next == node || (ranks[next] > 0 && ranks[next] == discoveredRank-1) {
			continue
		}
		if ranks[next] > 0 {
			minNeighboredRank = min(minNeighboredRank, ranks[next])
			continue
		}
		temp := dfs(next, graph, ranks, discoveredRank+1, ans)

		if temp > ranks[node] {
			*ans = append(*ans, []int{node, next})
		}
		minNeighboredRank = min(minNeighboredRank, temp)
	}
	return minNeighboredRank
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
