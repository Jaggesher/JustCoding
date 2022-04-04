package main

import "fmt"

func main() {
	fmt.Println("No Test Case")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/***
 * Time: O(n)
 * Space: O(1)
 */
func swapNodes(head *ListNode, k int) *ListNode {
	first, last := head, head
	for index, temp := 1, head; temp != nil; temp, index = temp.Next, index+1 {
		if index == k {
			first = temp
			last = head
		} else if index > k {
			last = last.Next
		}
	}

	first.Val, last.Val = last.Val, first.Val
	return head
}
