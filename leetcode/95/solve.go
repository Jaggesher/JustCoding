package main

import "fmt"

func main() {
	fmt.Println("Case 1:", generateTrees(3))
	fmt.Println("Case 2:", generateTrees(1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	var dp [10][10][]*TreeNode

	for i := 0; i < n; i++ {
		for j := 1; (j + i) <= n; j++ {
			st, end := j, j+i
			dp[st][end] = make([]*TreeNode, 0)
			if st == end {
				dp[st][end] = append(dp[st][end], &TreeNode{Val: st, Left: nil, Right: nil})
				continue
			}

			for root := st; root <= end; root++ {
				if st == root {
					for _, right := range dp[root+1][end] {
						dp[st][end] = append(dp[st][end], &TreeNode{Val: root, Left: nil, Right: right})
					}
				} else if root == end {
					for _, left := range dp[st][root-1] {
						dp[st][end] = append(dp[st][end], &TreeNode{Val: root, Left: left, Right: nil})
					}
				} else {
					for _, left := range dp[st][root-1] {
						for _, right := range dp[root+1][end] {
							dp[st][end] = append(dp[st][end], &TreeNode{Val: root, Left: left, Right: right})
						}
					}
				}
			}
		}
	}
	return dp[1][n]
}
