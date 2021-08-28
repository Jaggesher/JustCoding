package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", findRedundantDirectedConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
	fmt.Println("Case 2: ", findRedundantDirectedConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}))
	fmt.Println("Case 3: ", findRedundantDirectedConnection([][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}}))
	fmt.Println("Case 3: ", findRedundantDirectedConnection([][]int{{4, 2}, {1, 5}, {5, 2}, {5, 3}, {2, 4}}))
}

func findRedundantDirectedConnection(edges [][]int) []int {
	var n int = len(edges) + 1
	var inDegree, outDegree [1001]int
	var seen [1001][1001]bool
	loop, sU, sV := false, 0, 0

	for _, vl := range edges {
		u, v := vl[0], vl[1]
		outDegree[u]++
		inDegree[v]++
		if seen[v][u] == true {
			loop, sU, sV = true, u, v
		}
		seen[u][v] = true
	}

	if loop {
		outDegree[sU]--
		inDegree[sV]--
		if testTree(inDegree, outDegree, n) {
			return []int{sU, sV}
		} else {
			return []int{sV, sU}
		}
	}

	for i := len(edges) - 1; i >= 0; i-- {
		u, v := edges[i][0], edges[i][1]
		outDegree[u]--
		inDegree[v]--
		if testTree(inDegree, outDegree, n) {
			return edges[i]
		}
		outDegree[u]++
		inDegree[v]++
	}

	return []int{}
}

func testTree(inDegree, outDegree [1001]int, n int) bool {
	flag := 0
	for i := 1; i < n; i++ {
		if inDegree[i] > 1 {
			return false
		} else if inDegree[i] == 0 {
			flag++
			if outDegree[i] == 0 {
				return false
			}
		}
	}
	return flag == 1
}
