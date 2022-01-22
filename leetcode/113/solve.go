package main

import "fmt"

func main() {
	fmt.Println("No test case")
}

type Node struct {
	Val       int
	Neighbors []*Node
}

/***
 * Time: O(N + M) => N is the number of Nodes & M is the number of edges
 * Space: O(N)
 */
func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	var nodes [101]Node
	var visited [101]bool
	for i := 1; i <= 100; i++ {
		nodes[i] = Node{Val: i, Neighbors: make([]*Node, 0)}
	}
	var queue []*Node = make([]*Node, 0)
	queue = append(queue, node)
	for len(queue) > 0 {
		tm := queue[0]
		queue = queue[1:]
		if visited[tm.Val] {
			continue
		}
		visited[tm.Val] = true
		for _, vl := range tm.Neighbors {
			queue = append(queue, vl)
			nodes[tm.Val].Neighbors = append(nodes[tm.Val].Neighbors, &nodes[vl.Val])
		}
	}
	return &nodes[node.Val]
}
