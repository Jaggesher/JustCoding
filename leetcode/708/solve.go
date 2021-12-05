package main

import "fmt"

func main() {
	fmt.Println("No Test case")
}

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		aNode := &Node{Val: x, Next: nil}
		aNode.Next = aNode
		return aNode
	}
	parent, node := aNode, aNode.Next
	for ; node != aNode; parent, node = node, node.Next {
		if (parent.Val <= x && x <= node.Val) || (parent.Val > node.Val && (parent.Val <= x || node.Val >= x)) {
			parent.Next = &Node{Val: x, Next: node}
			return aNode
		}
	}
	parent.Next = &Node{Val: x, Next: node}
	return aNode
}
