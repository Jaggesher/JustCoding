package main

import "fmt"

func main() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	fmt.Println(findMinHeightTrees(1, [][]int{}))
	fmt.Println(findMinHeightTrees(2, [][]int{{0, 1}}))
	fmt.Println(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}))
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n <= 2 {
		ans := []int{}
		for i := 0; i < n; i++ {
			ans = append(ans, i)
		}
		return ans
	}

	var graph [][]int = make([][]int, n)
	var tracker []int = make([]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	for _, vl := range edges {
		u, v := vl[0], vl[1]
		graph[v] = append(graph[v], u)
		graph[u] = append(graph[u], v)
		tracker[u]++
		tracker[v]++
	}

	nodes := []int{}
	for i := 0; i < n; i++ {
		if tracker[i] <= 1 {
			nodes = append(nodes, i)
		}
	}

	for n > 2 {
		n -= len(nodes)
		tm := []int{}
		for _, u := range nodes {
			for _, v := range graph[u] {
				tracker[v]--
				if tracker[v] == 1 {
					tm = append(tm, v)
				}
			}
		}
		nodes = tm
	}
	return nodes
}
