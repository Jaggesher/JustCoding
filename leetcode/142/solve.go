package main

import "fmt"

func main() {
	fmt.Println("No Test case ")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/***
 * Time: O(N)
 * Space: O(1)
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for flag := fast != nil && fast.Next != nil; flag; flag = fast != nil && fast.Next != nil && fast != slow {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil && fast.Next != nil {
		fast = head
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}
		return fast
	}

	return nil
}
