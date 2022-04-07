package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1: ", lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println("Case 2: ", lastStoneWeight([]int{1}))
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	var e interface{}
	*h, e = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return e
}

/***
 * Approach: 2
 * Time: O(n log n)
 * Space: O(n)
 */

func lastStoneWeight_1(stones []int) int {
	h := intHeap(stones)
	heap.Init(&h)
	for h.Len() > 1 {
		x, y := heap.Pop(&h).(int), heap.Pop(&h).(int)
		heap.Push(&h, (x - y))
	}
	return heap.Pop(&h).(int)
}

/***
 * Approach: 2
 * Time: O(n*n)
 * Space: O(n)
 */
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Slice(stones, func(a, b int) bool { return stones[a] > stones[b] })
		stones[1] = stones[0] - stones[1]
		stones = stones[1:]
	}
	return stones[0]
}
