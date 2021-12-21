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
 * Time: O(n)+O(n/2) => O(n)
 * Space: O(1)
 */
func isPalindrome(head *ListNode) bool {
	var prev, fast, slow *ListNode = nil, head, head

	for fast != nil && fast.Next != nil {
		temp := slow

		fast = fast.Next.Next
		slow = slow.Next

		temp.Next, prev = prev, temp
	}

	if fast != nil {
		slow = slow.Next
	}

	for prev != nil && slow != nil {
		if prev.Val != slow.Val {
			return false
		}
		prev = prev.Next
		slow = slow.Next
	}

	return true
}
