package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("No test case")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Pair struct {
	level int
	val   int
}

func verticalTraversal(root *TreeNode) [][]int {
	var keys []int = make([]int, 0)
	var tracker map[int][]Pair = make(map[int][]Pair)
	dfs(root, 0, 0, &keys, tracker)

	sort.Ints(keys)
	var ans [][]int = make([][]int, len(keys))
	for i, vl := range keys {
		sort.Slice(tracker[vl], func(i, j int) bool {
			if tracker[vl][i].level == tracker[vl][j].level {
				return tracker[vl][i].val < tracker[vl][j].val
			}
			return tracker[vl][i].level < tracker[vl][j].level
		})
		ans[i] = make([]int, len(tracker[vl]))
		for index, tm := range tracker[vl] {
			ans[i][index] = tm.val
		}
	}
	return ans
}

func dfs(node *TreeNode, i, j int, keys *[]int, tracker map[int][]Pair) {
	if node == nil {
		return
	}
	if tracker[j] == nil {
		tracker[j] = make([]Pair, 0)
		*keys = append(*keys, j)
	}
	tracker[j] = append(tracker[j], Pair{level: i, val: node.Val})
	dfs(node.Left, i+1, j-1, keys, tracker)
	dfs(node.Right, i+1, j+1, keys, tracker)
}
