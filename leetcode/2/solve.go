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
 * Time: (n)
 * Space: O(1)
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head, extra := &ListNode{Val: 0, Next: nil}, 0
	temp := head
	for l1 != nil || l2 != nil || extra > 0 {
		if l1 != nil {
			extra += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			extra += l2.Val
			l2 = l2.Next
		}

		temp.Next = &ListNode{Val: (extra % 10), Next: nil}
		extra /= 10
		temp = temp.Next
	}

	return head.Next
}
