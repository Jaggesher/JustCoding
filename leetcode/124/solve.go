package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var ans int = root.Val
	calculatePath(root, &ans)
	return ans
}

func calculatePath(node *TreeNode, ans *int) int {
	if node == nil {
		return 0
	}
	possible, left, right := node.Val, calculatePath(node.Left, ans), calculatePath(node.Right, ans)
	if left > 0 {
		possible += left
	}
	if right > 0 {
		possible += right
	}
	if *ans < possible {
		*ans = possible
	}

	tm := left
	if tm < right {
		tm = right
	}
	if tm > 0 {
		return node.Val + tm
	}
	return node.Val
}
