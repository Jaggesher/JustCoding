package main

import "fmt"

func main() {
	fmt.Println("Case 1:", sumOfDistancesInTree(6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}))
	fmt.Println("Case 2:", sumOfDistancesInTree(2, [][]int{{1, 0}}))
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	var ans []int = make([]int, n)
	var nodes []int = make([]int, n)
	var graph [][]int = make([][]int, n)

	for i := 0; i < n; i++ {
		ans[i] = -1
		graph[i] = make([]int, 0)
	}

	for _, vl := range edges {
		u, v := vl[0], vl[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	ans[0], _ = dfs(0, graph, nodes, 0)
	bfs(nodes, ans, graph, n)
	return ans
}

func bfs(nodes []int, ans []int, graph [][]int, n int) {
	queue := []int{0}
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		for _, vl := range graph[node] {
			if ans[vl] == -1 {
				queue = append(queue, vl)
				ans[vl] = ans[node] - nodes[vl] + (n - nodes[vl])
			}
		}

	}
}

func dfs(node int, graph [][]int, nodes []int, level int) (sum, child int) {
	sum += level
	nodes[node] = 1
	for _, vl := range graph[node] {
		if nodes[vl] == 0 {
			tmSum, tmChild := dfs(vl, graph, nodes, level+1)
			sum += tmSum
			nodes[node] += tmChild
		}
	}
	return sum, nodes[node]
}
