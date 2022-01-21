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
 * Time: O(N log N)
 * Space: O(log N)
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := FindMid(head)
	rightHead := mid.Next
	mid.Next = nil
	left, right := sortList(head), sortList(rightHead)
	return merge(left, right)
}

func FindMid(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for flag := fast != nil && fast.Next != nil; flag; flag = fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	temp := dummy
	for left != nil && right != nil {
		if left.Val <= right.Val {
			temp.Next = left
			left = left.Next
			temp = temp.Next
		} else {
			temp.Next = right
			right = right.Next
			temp = temp.Next
		}
	}
	if left != nil {
		temp.Next = left
	}

	if right != nil {
		temp.Next = right
	}
	return dummy.Next
}
