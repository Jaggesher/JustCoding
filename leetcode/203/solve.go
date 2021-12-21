package main

import "fmt"

func main() {
	fmt.Println("No test Case")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/***
 * Time: O(n)
 * Space: O(1)
 */
func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{}
	prev := newHead

	for head != nil {
		if head.Val != val {
			prev.Next, prev = head, head
		}
		head, prev.Next = head.Next, nil
	}

	return newHead.Next
}
