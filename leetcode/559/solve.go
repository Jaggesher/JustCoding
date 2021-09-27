package main

import "fmt"

func main() {
	fmt.Println("No test case")
}

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	mx := 1
	for _, vl := range root.Children {
		tm := 1 + maxDepth(vl)
		if tm > mx {
			mx = tm
		}
	}
	return mx
}
