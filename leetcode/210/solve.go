package main

import "fmt"

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {3, 0}, {2, 3}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}

var ans []int

func findOrder(numCourses int, prerequisites [][]int) []int {
	var visited []bool = make([]bool, numCourses)
	var backage []bool = make([]bool, numCourses)
	var graph [][]int = make([][]int, numCourses)
	ans = make([]int, 0)
	for _, item := range prerequisites {
		v, u := item[0], item[1]
		if u == v {
			return ans
		}
		if graph[u] == nil {
			graph[u] = make([]int, 0)
		}
		graph[u] = append(graph[u], v)
	}

	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			if tryMakeTopSort(i, graph, visited, backage) {
				return []int{}
			}
		}
	}

	for st, end := 0, len(ans)-1; st < end; st, end = st+1, end-1 {
		ans[st], ans[end] = ans[end], ans[st]
	}

	return ans
}

func tryMakeTopSort(node int, graph [][]int, visited, backage []bool) bool {
	visited[node], backage[node] = true, true
	flag := false
	for _, item := range graph[node] {
		if backage[item] || flag {
			return true
		} else if !visited[item] {
			flag = flag || tryMakeTopSort(item, graph, visited, backage)
		}
	}
	ans = append(ans, node)
	backage[node] = false
	return flag
}
