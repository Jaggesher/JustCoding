package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("Case 1: ", kClosest([][]int{{1, 3}, {-2, 2}}, 1))
	fmt.Println("Case 1: ", kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}

type Node struct {
	point    []int
	distance int
}

type NodeHeap []Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].distance > h[j].distance }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/***
 * Approach: PQ(MAX Heap)
 * TIME: O(n * log K)
 * SPACE: O(k)
 */
func kClosest(points [][]int, k int) [][]int {
	var pq NodeHeap = make(NodeHeap, k)
	var ans [][]int = make([][]int, k)
	for i := 0; i < k; i++ {
		pq[i] = Node{
			point:    points[i],
			distance: (points[i][0] * points[i][0]) + (points[i][1] * points[i][1]),
		}
	}

	heap.Init(&pq)

	for i := k; i < len(points); i++ {
		tmDistance := (points[i][0] * points[i][0]) + (points[i][1] * points[i][1])
		if pq[0].distance > tmDistance {
			pq[0] = Node{point: points[i], distance: tmDistance}
			heap.Fix(&pq, 0)
		}
	}

	for index := range ans {
		ans[len(ans)-index-1] = heap.Pop(&pq).(Node).point
	}
	return ans
}
