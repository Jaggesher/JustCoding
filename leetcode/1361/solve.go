package main

import "fmt"

func main() {
	fmt.Println("Hello Jaggesher")
	n1, l1, r1 := 4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}
	fmt.Println("Case1:= ", validateBinaryTreeNodes(n1, l1, r1))

	n2, l2, r2 := 4, []int{1, -1, 3, -1}, []int{2, 3, -1, -1}
	fmt.Println("Case1:= ", validateBinaryTreeNodes(n2, l2, r2))

	n3, l3, r3 := 2, []int{1, 0}, []int{-1, -1}
	fmt.Println("Case1:= ", validateBinaryTreeNodes(n3, l3, r3))

	n4, l4, r4 := 6, []int{1, -1, -1, 4, -1, -1}, []int{2, -1, -1, 5, -1, -1}
	fmt.Println("Case1:= ", validateBinaryTreeNodes(n4, l4, r4))

}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	var nodes []int = make([]int, n)
	for i := 0; i < n; i++ {
		nodes[i] = i
	}
	stCount := n
	for i := 0; i < n; i++ {
		parrent, left, right := findParrent(nodes, i), leftChild[i], rightChild[i]
		if left != -1 {
			if nodes[left] != left || parrent == left {
				return false
			}
			nodes[left] = i
			stCount--
		}
		if right != -1 {
			if nodes[right] != right || parrent == right {
				return false
			}
			nodes[right] = i
			stCount--
		}
	}
	if stCount > 1 {
		return false
	}
	return true
}

func findParrent(nodes []int, node int) int {
	for nodes[node] != node {
		node = nodes[node]
	}
	return node
}
