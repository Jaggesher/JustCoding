package main

import "fmt"

func main() {
	fmt.Println("No test case")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/***
 * Time: O(n) => where n is the length of range
 * Space: O(n) =>
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var container []*ListNode = make([]*ListNode, 0)
	tm := head
	for i := 1; tm != nil; i++ {
		if left <= i && i <= right {
			container = append(container, tm)
		}
		tm = tm.Next
	}
	for st, end := 0, len(container)-1; st < end; st, end = st+1, end-1 {
		container[st].Val, container[end].Val = container[end].Val, container[st].Val
	}
	return head
}
