package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1:", minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println("Case 2:", minMeetingRooms([][]int{{7, 10}, {2, 4}}))
}

type PQ []int

func (h PQ) Len() int           { return len(h) }
func (h PQ) Less(i, j int) bool { return h[i] < h[j] }
func (h PQ) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PQ) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *PQ) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minMeetingRooms(intervals [][]int) int {
	var ans int = 0
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	pq := &PQ{}
	heap.Init(pq)

	for _, vl := range intervals {
		for pq.Len() > 0 && (*pq)[0] <= vl[0] {
			heap.Pop(pq)
		}
		heap.Push(pq, vl[1])
		if ans < pq.Len() {
			ans = pq.Len()
		}
	}
	return ans
}
