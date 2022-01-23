package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}}))
	fmt.Println("Case 2: ", countComponents(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}))
}

/***
 * Time: O(E+V)
 * Space: O(E+V)
 */
func countComponents(n int, edges [][]int) int {
	var ans int = 0
	var grid [][]int = make([][]int, n)
	var visited []bool = make([]bool, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		grid[u] = append(grid[u], v)
		grid[v] = append(grid[v], u)
	}
	for node, flag := range visited {
		if flag {
			continue
		}
		ans++
		visited[node] = true
		queue := []int{node}
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range grid[u] {
				if !visited[v] {
					visited[v] = true
					queue = append(queue, v)
				}
			}
		}
	}
	return ans
}
