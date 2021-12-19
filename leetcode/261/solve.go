package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}))
	fmt.Println("Case 2: ", validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}))
}

func validTree(n int, edges [][]int) bool {
	if n != len(edges)+1 {
		return false
	}

	var visited [2005]bool
	var matrix [2005][]int
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		matrix[u] = append(matrix[u], v)
		matrix[v] = append(matrix[v], u)
	}

	visited[0] = true
	queue := []int{0}
	n--

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range matrix[u] {
			if !visited[v] {
				n--
				queue = append(queue, v)
				visited[v] = true
			}
		}
	}
	return n == 0
}
