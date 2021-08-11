package main

import "fmt"

func main() {
	var edges1 [][]int = [][]int{{1, 2}, {1, 3}, {2, 3}}
	var edges2 [][]int = [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	var edges3 [][]int = [][]int{{1, 4}, {3, 4}, {1, 3}, {1, 2}, {4, 5}}

	fmt.Println("Input: ", edges1, "Out: ", findRedundantConnection(edges1))
	fmt.Println("Input: ", edges2, "Out: ", findRedundantConnection(edges2))
	fmt.Println("Input: ", edges3, "Out: ", findRedundantConnection(edges3))
}

func findRedundantConnection(edges [][]int) []int {
	var ans []int = make([]int, 2)
	var n int = len(edges)
	var unionSet []int = make([]int, n+1)
	for i := 1; i <= n; i++ {
		unionSet[i] = i
	}
	for _, vl := range edges {
		//fmt.Println(unionSet)
		u, v := vl[0], vl[1]
		prU, prV := findParrent(unionSet, u), findParrent(unionSet, v)

		if prU != prV {
			unionSet[prV] = prU
		} else {
			ans[0], ans[1] = u, v
		}
	}
	return ans
}

func findParrent(unionSet []int, node int) int {
	for node != unionSet[node] {
		node = unionSet[node]
	}
	return node
}
