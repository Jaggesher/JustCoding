package main

import "fmt"

func main() {
	fmt.Println("No Test case")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/***
 * Time: O(n)
 * Space: O(1)
 */

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	temp := dummy
	for temp != nil {
		next := temp.Next
		if next != nil && next.Next != nil && next.Val == next.Next.Val {
			for next.Next != nil && next.Val == next.Next.Val {
				next = next.Next
			}
			temp.Next = next.Next
		} else {
			temp = temp.Next
		}
	}

	return dummy.Next
}
