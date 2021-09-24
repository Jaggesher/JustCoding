package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1: ", findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println("Case 2: ", findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

/**
	Using HEAP
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := intHeap(nums[0:k])
	heap.Init(&h)
	for k < len(nums) {
		if h[0] < nums[k] {
			h[0] = nums[k]
			heap.Fix(&h, 0)
		}
		k++
	}
	return h[0]
}
*/
