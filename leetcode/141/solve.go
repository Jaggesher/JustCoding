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
 * Time: O(n+K) => O(n)
 * Space: O(1)
 */

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
