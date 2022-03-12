package main

import "fmt"

func main() {
	fmt.Println("No test case")
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/***
 * Approach: 1
 * Time: O(n)
 * Space: O(n)
 */
func copyRandomList(head *Node) *Node {
	var tracker []*Node = make([]*Node, 0)
	ans := &Node{Val: 0, Next: nil, Random: nil}

	for index, temp, tempHead := 0, ans, head; tempHead != nil; index++ {
		tm := &Node{Val: tempHead.Val, Next: nil, Random: nil}
		tracker = append(tracker, tm)
		temp.Next, temp = tm, tm
		tempHead.Val, tempHead = index, tempHead.Next
	}

	for cp, original := ans.Next, head; cp != nil && original != nil; cp, original = cp.Next, original.Next {
		if original.Random != nil {
			cp.Random = tracker[original.Random.Val]
		}
	}

	return ans.Next
}
