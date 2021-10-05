package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("No test case")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	val := root.Val
	if float64(val) == target {
		return val
	} else if float64(val) < target && root.Right != nil {
		tm := closestValue(root.Right, target)
		if math.Abs(float64(val)-target) > math.Abs(float64(tm)-target) {
			val = tm
		}
	} else if float64(val) > target && root.Left != nil {
		tm := closestValue(root.Left, target)
		if math.Abs(float64(val)-target) > math.Abs(float64(tm)-target) {
			val = tm
		}
	}

	return val
}
