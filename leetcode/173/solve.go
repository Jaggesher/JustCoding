package main

import "fmt"

func main() {
	fmt.Println("No test Case")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	nodes []*TreeNode
}

func insertIfAny(nodes []*TreeNode) []*TreeNode {
	for nodes[len(nodes)-1].Left != nil {
		nodes = append(nodes, nodes[len(nodes)-1].Left)
	}
	return nodes
}

func Constructor(root *TreeNode) BSTIterator {
	nodes := insertIfAny([]*TreeNode{root})
	return BSTIterator{nodes}
}

func (this *BSTIterator) Next() int {
	tm := this.nodes[len(this.nodes)-1]
	this.nodes = this.nodes[:len(this.nodes)-1]
	if tm.Right != nil {
		this.nodes = append(this.nodes, tm.Right)
		this.nodes = insertIfAny(this.nodes)
	}
	return tm.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.nodes) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
