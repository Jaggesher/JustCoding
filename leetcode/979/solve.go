package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello Jaggesher")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	_, ans := dfs(root)
	return ans
}

func dfs(node *TreeNode) (coin, move int) {
	if node == nil {
		return 0, 0
	}
	coinsL, moveL := dfs(node.Left)
	coinsR, moveR := dfs(node.Right)

	coins, move := coinsL+coinsR, moveL+moveR
	coins += (node.Val - 1)
	move += int(math.Abs(float64(coins)))

	return coins, move
}
