package main

import "fmt"

func main() {
	fmt.Println("No test Case")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	var ans []*ListNode = make([]*ListNode, k)
	if head == nil {
		return ans
	}
	count := 1
	for temp := head; temp.Next != nil; temp = temp.Next {
		count++
	}

	n, extra := int(count/k), count%k

	for i := 0; i < k; i++ {
		ans[i] = head
		if (i + 1) >= k {
			break
		}

		for j := 1; j < n && head != nil; j++ {
			head = head.Next
		}

		if n > 0 && extra > 0 {
			head = head.Next
		}

		if head != nil {
			head.Next, head = nil, head.Next
		}
		extra--
	}

	return ans
}
