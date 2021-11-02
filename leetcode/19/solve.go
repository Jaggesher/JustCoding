package main

import "fmt"

func main() {
	fmt.Println("No test case")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var prev *ListNode = nil
	for head.Next != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	head.Next, prev = prev, nil
	for head != nil {
		n--
		if n == 0 {
			head = head.Next
			continue
		}
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
