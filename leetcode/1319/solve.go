package main

import "fmt"

func main() {
	fmt.Println("Case 1:", makeConnected(4, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	fmt.Println("Case 2:", makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}))
}

func makeConnected(n int, connections [][]int) int {
	if n > len(connections)+1 {
		return -1
	}
	var visited []bool = make([]bool, n)
	var graph [][]int = make([][]int, n)
	for _, vl := range connections {
		graph[vl[0]] = append(graph[vl[0]], vl[1])
		graph[vl[1]] = append(graph[vl[1]], vl[0])
	}

	var count int = 0
	for i, vl := range visited {
		if !vl {
			count++
			bfs([]int{i}, visited, graph)
		}
	}
	return count - 1
}

func bfs(queue []int, visited []bool, graph [][]int) {
	for len(queue) > 0 {
		temp := make([]int, 0)
		for _, vl := range queue {
			if visited[vl] {
				continue
			}
			visited[vl] = true
			temp = append(temp, graph[vl]...)
		}
		queue = temp
	}
}
