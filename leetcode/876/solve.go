package main

import "fmt"

func main() {
	fmt.Println("No Test case")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil {
		if fast.Next == nil {
			return slow.Next
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
