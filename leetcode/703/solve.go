package main

import (
	"container/heap"
	"fmt"
)

func main() {
	obj := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println("Case 1: ", obj.Add(3))
	fmt.Println("Case 2: ", obj.Add(5))
	fmt.Println("Case 3: ", obj.Add(10))
	fmt.Println("Case 4: ", obj.Add(9))
	fmt.Println("Case 5: ", obj.Add(4))
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	var e interface{}
	*h, e = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return e
}

type KthLargest struct {
	k      int
	minHep MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	minHeap := MinHeap(nums)
	heap.Init(&minHeap)
	for minHeap.Len() > k {
		heap.Pop(&minHeap)
	}
	return KthLargest{k, minHeap}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.minHep, val)
	if this.minHep.Len() > this.k {
		heap.Pop(&this.minHep)
	}
	return this.minHep[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
