package main

import (
	"fmt"
)

func main() {
	fmt.Println(canFinish(4, [][]int{{0, 1}, {1, 2}, {2, 3}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	var visited []bool = make([]bool, numCourses)
	var backage []bool = make([]bool, numCourses)
	var graph [][]int = make([][]int, numCourses)

	for _, item := range prerequisites {
		i, j := item[0], item[1]
		if i == j {
			return false
		}
		if graph[i] == nil {
			graph[i] = make([]int, 0)
		}
		graph[i] = append(graph[i], j)
	}

	for i := 0; i < numCourses; i++ {
		if visited[i] == false {
			if hasCycle(i, graph, visited, backage) {
				return false
			}
		}
	}
	return true
}

func hasCycle(node int, graph [][]int, visited, backage []bool) bool {
	visited[node], backage[node] = true, true
	flag := false
	for _, item := range graph[node] {
		if backage[item] || flag {
			return true
		} else if !visited[item] {
			flag = flag || hasCycle(item, graph, visited, backage)
		}
	}
	backage[node] = false
	return flag
}
