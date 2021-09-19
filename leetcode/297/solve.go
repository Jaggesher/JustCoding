package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	data  []string
	index int
}

func Constructor() Codec {
	return Codec{data: []string{}, index: 0}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "n"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.data = strings.Split(data, ",")
	this.index = -1
	return this.construct()
}

func (this *Codec) construct() *TreeNode {
	this.index++
	if this.data[this.index] == "n" {
		return nil
	}
	vl, _ := strconv.Atoi(this.data[this.index])
	tm := TreeNode{Val: vl, Left: this.construct(), Right: this.construct()}
	return &tm
}

func main() {
	fmt.Println("No test case")
	/**
	 * Your Codec object will be instantiated and called as such:
	 * ser := Constructor();
	 * deser := Constructor();
	 * data := ser.serialize(root);
	 * ans := deser.deserialize(data);
	 */
}
