package main

import "fmt"

func main() {
	fmt.Println("No test case")
}

type Node struct {
	Val      int
	Children []*Node
}

func diameter(root *Node) int {
	var ans int = 0
	findDistance(root, &ans)
	return ans
}

func findDistance(node *Node, ans *int) int {
	if node == nil {
		return 0
	}
	a, b := 0, 0
	for _, child := range node.Children {
		a, b = topTwo(a, b, findDistance(child, ans))
		if (a + b) > *ans {
			*ans = a + b
		}
	}
	return a + 1
}

func topTwo(a, b, c int) (int, int) {
	if b < c {
		b, c = c, b
	}
	if a < b {
		a, b = b, a
	}
	return a, b
}
