package main

import (
	"fmt"
)

func main() {
	fmt.Println("No test case")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/***
 * Time: O(n)
 * Space: O(1)
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	tail, count := head, 1
	for tail.Next != nil {
		tail = tail.Next
		count++
	}
	k = (count - (k % count)) // Cut point
	tail.Next = head          // Making ring

	for k > 1 {
		head = head.Next
		k--
	}

	tail = head.Next
	head.Next = nil
	return tail
}
