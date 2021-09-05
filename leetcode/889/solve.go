package main

import "fmt"

func main() {
	fmt.Println("Case 1:", constructFromPrePost([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}))
	fmt.Println("Case 2:", constructFromPrePost([]int{1}, []int{1}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 || len(postorder) == 0 {
		return nil
	}
	node := TreeNode{Val: preorder[0], Right: nil, Left: nil}
	ln := len(preorder)
	if ln > 1 {
		tm := -1
		for i, vl := range postorder {
			if vl == preorder[1] {
				tm = i + 1
				break
			}
		}

		node.Left = constructFromPrePost(preorder[1:tm+1], postorder[:tm])
		node.Right = constructFromPrePost(preorder[tm+1:], postorder[tm:ln])
	}

	return &node
}
