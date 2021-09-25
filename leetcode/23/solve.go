package main

import "fmt"

func main() {
	fmt.Println("Hello Jaggesher")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	ln := len(lists)
	if ln == 0 {
		return nil
	} else if ln == 1 {
		return lists[0]
	}
	return mergeTwoLists(mergeKLists(lists[:(ln/2)]), mergeKLists(lists[(ln/2):]))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var root *ListNode = &ListNode{Val: 0, Next: nil}
	tail := root
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	}

	if list2 != nil {
		tail.Next = list2
	}
	return root.Next
}
